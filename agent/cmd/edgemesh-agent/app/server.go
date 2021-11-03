package app

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/cli/globalflag"
	"k8s.io/component-base/term"
	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"

	"github.com/kubeedge/beehive/pkg/core"
	"github.com/kubeedge/edgemesh/agent/cmd/edgemesh-agent/app/config"
	"github.com/kubeedge/edgemesh/agent/cmd/edgemesh-agent/app/config/validation"
	"github.com/kubeedge/edgemesh/agent/cmd/edgemesh-agent/app/options"
	"github.com/kubeedge/edgemesh/agent/pkg/chassis"
	"github.com/kubeedge/edgemesh/agent/pkg/dns"
	"github.com/kubeedge/edgemesh/agent/pkg/gateway"
	"github.com/kubeedge/edgemesh/agent/pkg/proxy"
	"github.com/kubeedge/edgemesh/agent/pkg/tunnel"
	"github.com/kubeedge/edgemesh/common/constants"
	"github.com/kubeedge/edgemesh/common/informers"
	commonutil "github.com/kubeedge/edgemesh/common/util"
	"github.com/kubeedge/kubeedge/pkg/util"
	"github.com/kubeedge/kubeedge/pkg/util/flag"
	"github.com/kubeedge/kubeedge/pkg/version"
	"github.com/kubeedge/kubeedge/pkg/version/verflag"
)

func NewEdgeMeshAgentCommand() *cobra.Command {
	opts := options.NewEdgeMeshAgentOptions()
	cmd := &cobra.Command{
		Use: "edgemesh-agent",
		Long: `edgemesh-agent is a part of EdgeMesh which provides a simple network solution
for the inter-communications between services at edge scenarios.`,
		Run: func(cmd *cobra.Command, args []string) {
			verflag.PrintAndExitIfRequested()
			flag.PrintMinConfigAndExitIfRequested(config.NewEdgeMeshAgentConfig())
			flag.PrintDefaultConfigAndExitIfRequested(config.NewEdgeMeshAgentConfig())
			flag.PrintFlags(cmd.Flags())

			if errs := opts.Validate(); len(errs) > 0 {
				klog.Fatal(util.SpliceErrors(errs))
			}

			agentCfg, err := opts.Config()
			if err != nil {
				klog.Fatal(err)
			}
			resetConfigByMode(detectRunningMode(), agentCfg)

			if errs := validation.ValidateEdgeMeshAgentConfiguration(agentCfg); len(errs) > 0 {
				klog.Fatal(util.SpliceErrors(errs.ToAggregate().Errors()))
			}

			klog.Infof("Version: %+v", version.Get())
			if err = Run(agentCfg); err != nil {
				klog.Fatalf("run edgemesh-agent failed: %v", err)
			}
		},
	}
	fs := cmd.Flags()
	namedFs := opts.Flags()
	verflag.AddFlags(namedFs.FlagSet("global"))
	flag.AddFlags(namedFs.FlagSet("global"))
	globalflag.AddGlobalFlags(namedFs.FlagSet("global"), cmd.Name())
	for _, f := range namedFs.FlagSets {
		fs.AddFlagSet(f)
	}

	usageFmt := "Usage:\n  %s\n"
	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		cliflag.PrintSections(cmd.OutOrStderr(), namedFs, cols)
		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		cliflag.PrintSections(cmd.OutOrStdout(), namedFs, cols)
	})

	return cmd
}

// Run runs EdgeMesh Agent
func Run(cfg *config.EdgeMeshAgentConfig) error {
	trace := 1

	klog.Infof("[%d] New informers manager", trace)
	ifm, err := informers.NewManager(cfg.KubeAPIConfig)
	if err != nil {
		return err
	}
	trace++

	klog.Infof("[%d] Prepare agent to run", trace)
	if err = prepareRun(cfg, ifm); err != nil {
		return err
	}
	trace++

	klog.Infof("[%d] Register beehive modules", trace)
	if errs := registerModules(cfg, ifm); len(errs) > 0 {
		return fmt.Errorf(util.SpliceErrors(errs))
	}
	trace++

	// As long as either the proxy module or the gateway module is enabled,
	// the go-chassis plugins must also be install.
	if cfg.Modules.EdgeProxyConfig.Enable || cfg.Modules.EdgeGatewayConfig.Enable {
		klog.Infof("[%d] Install go-chassis plugins", trace)
		chassis.Install(cfg.GoChassisConfig, ifm)
		trace++
	}

	klog.Infof("[%d] Start informers manager", trace)
	ifm.Start(wait.NeverStop)
	trace++

	klog.Infof("[%d] Start all modules", trace)
	core.Run()

	klog.Infof("edgemesh-agent exited")
	return nil
}

// registerModules register all the modules started in edgemesh-agent
func registerModules(c *config.EdgeMeshAgentConfig, ifm *informers.Manager) []error {
	var errs []error
	if err := dns.Register(c.Modules.EdgeDNSConfig, ifm); err != nil {
		errs = append(errs, err)
	}
	if err := proxy.Register(c.Modules.EdgeProxyConfig, ifm); err != nil {
		errs = append(errs, err)
	}
	if err := gateway.Register(c.Modules.EdgeGatewayConfig, ifm); err != nil {
		errs = append(errs, err)
	}

	mode := tunnel.UnknownMode
	if c.Modules.EdgeGatewayConfig.Enable {
		mode = tunnel.ClientMode
	}
	if c.Modules.EdgeProxyConfig.Enable {
		mode = tunnel.ServerClientMode
	}

	if err := tunnel.Register(c.Modules.TunnelAgentConfig, ifm, mode); err != nil {
		errs = append(errs, err)
	}
	return errs
}

// prepareRun prepares edgemesh-agent to run
func prepareRun(c *config.EdgeMeshAgentConfig, ifm *informers.Manager) error {
	// set dns and proxy modules listenInterface
	if c.Modules.EdgeDNSConfig.Enable || c.Modules.EdgeProxyConfig.Enable {
		if err := commonutil.CreateDummyDevice(c.CommonConfig.DummyDeviceName, c.CommonConfig.DummyDeviceIP); err != nil {
			return fmt.Errorf("create dummy device %s err: %v", c.CommonConfig.DummyDeviceName, err)
		}
		c.Modules.EdgeDNSConfig.ListenInterface = c.CommonConfig.DummyDeviceName
		c.Modules.EdgeProxyConfig.ListenInterface = c.CommonConfig.DummyDeviceName
	}

	// set proxy module subNet, subNet equals to k8s service-cluster-ip-range
	if c.Modules.EdgeProxyConfig.Enable && c.Modules.EdgeProxyConfig.SubNet == "" {
		subNet, err := getClusterServiceCIDR(ifm.GetKubeClient())
		if err != nil {
			return fmt.Errorf("get service-cluster-ip-range err: %v", err)
		}
		c.Modules.EdgeProxyConfig.SubNet = subNet

		if err := resetConfigMapSubNet(c.CommonConfig.ConfigMapName, subNet, ifm.GetKubeClient()); err != nil {
			return fmt.Errorf("reset edgemesh-agent configmap subNet err: %v", err)
		}
	}

	return nil
}

// getClusterServiceCIDR creates an impossible service to cause an error,
// and obtains service-cluster-ip-range from the error message
func getClusterServiceCIDR(kubeClient kubernetes.Interface) (string, error) {
	if kubeClient == nil {
		return "", fmt.Errorf("kubeClient is nil")
	}

	badService := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "bad-service",
		},
		Spec: corev1.ServiceSpec{
			Type:      "ClusterIP",
			ClusterIP: "0.0.0.0", // this is an impossible cluster ip
			Ports:     []corev1.ServicePort{{Port: 443}},
		},
	}

	svc, err := kubeClient.CoreV1().Services(constants.EdgeMeshNamespace).Create(context.Background(), &badService, metav1.CreateOptions{})
	if err == nil {
		return "", fmt.Errorf("impossible happened, %s was created successfully", svc.Name)
	}

	errMsg := fmt.Sprintf("%v", err)
	errKey := "The range of valid IPs is "
	if ok := strings.Contains(errMsg, errKey); !ok {
		return "", fmt.Errorf("unexpected error: %v", err)
	}

	info := strings.Split(errMsg, errKey)
	if len(info) != 2 {
		return "", fmt.Errorf("invalid error: %v", err)
	}

	return info[1], nil
}

// resetConfigMapSubNet reset edgemesh-agent configmap subNet value
func resetConfigMapSubNet(name, subNet string, kubeClient kubernetes.Interface) error {
	if name == "" {
		return fmt.Errorf("configmap name is empty")
	}

	if subNet == "" {
		return fmt.Errorf("subNet is empty")
	}

	if kubeClient == nil {
		return fmt.Errorf("kubeClient is nil")
	}

	cm, err := kubeClient.CoreV1().ConfigMaps(constants.EdgeMeshNamespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	data, ok := cm.Data[constants.EdgeMeshAgentConfigFileName]
	if !ok {
		return fmt.Errorf("configmap data %s not found", constants.EdgeMeshAgentConfigFileName)
	}

	var config config.EdgeMeshAgentConfig
	if err := yaml.Unmarshal([]byte(data), &config); err != nil {
		return err
	}

	if config.Modules.EdgeProxyConfig.SubNet != "" {
		klog.V(4).Infof("subNet has already been set up")
		return nil
	}

	// set configmap subNet value
	config.Modules.EdgeProxyConfig.SubNet = subNet
	newData, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	klog.V(4).Infof("new configmap:\n%v", string(newData))

	// overwrite old configmap data
	cm.Data[constants.EdgeMeshAgentConfigFileName] = string(newData)
	if _, err := kubeClient.CoreV1().ConfigMaps(constants.EdgeMeshNamespace).Update(context.Background(), cm, metav1.UpdateOptions{}); err != nil {
		return err
	}

	return nil
}

const (
	EdgeMode  string = "EdgeMode"
	CloudMode string = "CloudMode"
)

// detectRunningMode detects whether the edgemesh-agent is running on cloud node or edge node.
// It will recognize whether there is KUBERNETES_PORT in the container environment variable, because
// edged will not inject KUBERNETES_PORT environment variable into the container, but kubelet will.
// what is edged: https://kubeedge.io/en/docs/architecture/edge/edged/
func detectRunningMode() string {
	_, exist := os.LookupEnv("KUBERNETES_PORT")
	if exist {
		return CloudMode
	}
	return EdgeMode
}

// resetConfigByMode will reset the edgemesh-agent configuration according to the mode.
func resetConfigByMode(mode string, c *config.EdgeMeshAgentConfig) {
	// if the user sets KubeConfig, nothing will be processed
	if c.KubeAPIConfig.KubeConfig != "" {
		return
	}

	klog.Infof("edgemesh-agent running on %s", mode)

	if mode == EdgeMode {
		// edgemesh-agent relies on the list-watch function of KubeEdge when it runs at the edge node.
		// KubeEdge v1.6+ starts to support this function until KubeEdge v1.7+ tends to be stable.
		// what is KubeEdge list-watch: https://github.com/kubeedge/kubeedge/blob/master/CHANGELOG/CHANGELOG-1.6.md
		c.KubeAPIConfig.Master = "http://127.0.0.1:10550"
		// ContentType only supports application/json
		// see issue: https://github.com/kubeedge/kubeedge/issues/3041
		c.KubeAPIConfig.ContentType = "application/json"
	}

	if mode == CloudMode {
		// when edgemesh-agent is running on the cloud, we do not need to enable edgedns,
		// because all domain name resolution can be done by CoreDNS or kube-dns.
		c.Modules.EdgeDNSConfig.Enable = false
	}
}
