# Getting Started

## Dependencies

[KubeEdge Dependencies](https://kubeedge.io/en/docs/#dependencies)

[KubeEdge v1.7+](https://github.com/kubeedge/kubeedge/releases)

::: tip
EdgeMesh relies on the [Local APIServer](https://github.com/kubeedge/kubeedge/blob/master/CHANGELOG/CHANGELOG-1.6.md) function of KubeEdge. KubeEdge v1.6+ starts to support this function until KubeEdge v1.7+ tends to be stable
:::

## Helm Installation

- **Step 1**: Enable Local APIServer

Refer to [Manual Installation-Step 3](#step3), enable Local APIServer.

- **Step 2**: Install Charts

Make sure you have installed Helm 3

```
helm install edgemesh \
  --set server.nodeName=<your node name> \
  --set server.publicIP=<your node eip> \
  https://raw.githubusercontent.com/kubeedge/edgemesh/main/build/helm/edgemesh.tgz
```

server.nodeName specifies the node deployed by edgemesh-server, and server.publicIP specifies the public IP of the node. The server.publicIP can be omitted, because edgemesh-server will automatically detect and configure the public IP of the node, but it is not guaranteed to be correct.

**Example：**

```shell
helm install edgemesh \
  --set server.nodeName=k8s-node1 \
  --set server.publicIP=119.8.211.54 \
  https://raw.githubusercontent.com/kubeedge/edgemesh/main/build/helm/edgemesh.tgz
```

::: warning
Please set server.nodeName and server.publicIP according to your K8s cluster, otherwise edgemesh-server may not run
:::

- **Step 3**: Check it out

```shell
$ helm ls
NAME            NAMESPACE       REVISION        UPDATED                                 STATUS          CHART           APP VERSION
edgemesh        default         1               2021-11-01 23:30:02.927346553 +0800 CST deployed        edgemesh-0.1.0  1.8.0
```

```shell
$ kubectl get all -n kubeedge
NAME                                   READY   STATUS    RESTARTS   AGE
pod/edgemesh-agent-4rhz4               1/1     Running   0          76s
pod/edgemesh-agent-7wqgb               1/1     Running   0          76s
pod/edgemesh-agent-9c697               1/1     Running   0          76s
pod/edgemesh-server-5f6d5869ff-4568p   1/1     Running   0          5m8s

NAME                            DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
daemonset.apps/edgemesh-agent   3         3         3       3            3           <none>          76s

NAME                              READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/edgemesh-server   1/1     1            1           5m8s

NAME                                         DESIRED   CURRENT   READY   AGE
replicaset.apps/edgemesh-server-5f6d5869ff   1         1         1       5m8s
```

## Manual Installation

- **Step 1**: Download EdgeMesh

```shell
$ git clone https://github.com/kubeedge/edgemesh.git
$ cd edgemesh
```

<a name="step3"></a>
- **Step 2**: Create CRDs

```shell
$ kubectl apply -f build/crds/istio/
```

- **Step 3**: Enable Local APIServer

At the edge node, open metaServer module (if your KubeEdge < 1.8.0, you also need to close edgeMesh module), and restart edgecore

```shell
$ vim /etc/kubeedge/config/edgecore.yaml
modules:
  ..
  edgeMesh:
    enable: false
  ..
  metaManager:
    metaServer:
      enable: true
..
```

```shell
$ systemctl restart edgecore
```

On the cloud, open the dynamicController module, and restart cloudcore

```shell
$ vim /etc/kubeedge/config/cloudcore.yaml
modules:
  ..
  dynamicController:
    enable: true
..
```

```shell
$ systemctl restart cloudcore
```

At the edge node, check if Local APIServer works

```shell
$ curl 127.0.0.1:10550/api/v1/services
{"apiVersion":"v1","items":[{"apiVersion":"v1","kind":"Service","metadata":{"creationTimestamp":"2021-04-14T06:30:05Z","labels":{"component":"apiserver","provider":"kubernetes"},"name":"kubernetes","namespace":"default","resourceVersion":"147","selfLink":"default/services/kubernetes","uid":"55eeebea-08cf-4d1a-8b04-e85f8ae112a9"},"spec":{"clusterIP":"10.96.0.1","ports":[{"name":"https","port":443,"protocol":"TCP","targetPort":6443}],"sessionAffinity":"None","type":"ClusterIP"},"status":{"loadBalancer":{}}},{"apiVersion":"v1","kind":"Service","metadata":{"annotations":{"prometheus.io/port":"9153","prometheus.io/scrape":"true"},"creationTimestamp":"2021-04-14T06:30:07Z","labels":{"k8s-app":"kube-dns","kubernetes.io/cluster-service":"true","kubernetes.io/name":"KubeDNS"},"name":"kube-dns","namespace":"kube-system","resourceVersion":"203","selfLink":"kube-system/services/kube-dns","uid":"c221ac20-cbfa-406b-812a-c44b9d82d6dc"},"spec":{"clusterIP":"10.96.0.10","ports":[{"name":"dns","port":53,"protocol":"UDP","targetPort":53},{"name":"dns-tcp","port":53,"protocol":"TCP","targetPort":53},{"name":"metrics","port":9153,"protocol":"TCP","targetPort":9153}],"selector":{"k8s-app":"kube-dns"},"sessionAffinity":"None","type":"ClusterIP"},"status":{"loadBalancer":{}}}],"kind":"ServiceList","metadata":{"resourceVersion":"377360","selfLink":"/api/v1/services"}}
```

- **Step 4**: Deploy edgemesh-server

```shell
$ kubectl apply -f build/server/edgemesh/
namespace/kubeedge configured
serviceaccount/edgemesh-server created
clusterrole.rbac.authorization.k8s.io/edgemesh-server created
clusterrolebinding.rbac.authorization.k8s.io/edgemesh-server created
configmap/edgemesh-server-cfg created
deployment.apps/edgemesh-server created
```

::: warning
Please set the value of 05-configmap.yaml's publicIP and 06-deployment.yaml's nodeName according to your K8s cluster, otherwise edgemesh-server may not run
:::

- **Step 5**: Deploy edgemesh-agent

```shell
$ kubectl apply -f build/agent/kubernetes/edgemesh-agent/
namespace/kubeedge configured
serviceaccount/edgemesh-agent created
clusterrole.rbac.authorization.k8s.io/edgemesh-agent created
clusterrolebinding.rbac.authorization.k8s.io/edgemesh-agent created
configmap/edgemesh-agent-cfg created
daemonset.apps/edgemesh-agent created
```

- **Step 6**: Check it out

```shell
$ kubectl get all -n kubeedge
NAME                                   READY   STATUS    RESTARTS   AGE
pod/edgemesh-agent-4rhz4               1/1     Running   0          76s
pod/edgemesh-agent-7wqgb               1/1     Running   0          76s
pod/edgemesh-agent-9c697               1/1     Running   0          76s
pod/edgemesh-server-5f6d5869ff-4568p   1/1     Running   0          5m8s

NAME                            DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
daemonset.apps/edgemesh-agent   3         3         3       3            3           <none>          76s

NAME                              READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/edgemesh-server   1/1     1            1           5m8s

NAME                                         DESIRED   CURRENT   READY   AGE
replicaset.apps/edgemesh-server-5f6d5869ff   1         1         1       5m8s
```
