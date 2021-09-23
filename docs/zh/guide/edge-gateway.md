# 边缘网关

EdgeMesh 的边缘网关提供了通过网关的方式访问集群内部服务的能力，本章节会指导您从头部署一个边缘网关。

![edgemesh-ingress-gateway](/images/guide/em-ig.png)

## 部署

```shell
$ kubectl apply -f build/agent/kubernetes/edgemesh-gateway/03-configmap.yaml
$ kubectl apply -f build/agent/kubernetes/edgemesh-gateway/04-deployment.yaml
```

::: tip
边缘网关与 Agent 使用相同的 [镜像](https://hub.docker.com/r/kubeedge/edgemesh-agent)，只在配置上有细微区别。
:::

## HTTP 网关

**创建 Gateway 资源对象和路由规则 VirtualService**

```shell
$ kubectl apply -f examples/hostname-lb-random-gateway.yaml
pod/hostname-lb-edge2 created
pod/hostname-lb-edge3 created
service/hostname-lb-svc created
gateway.networking.istio.io/edgemesh-gateway configured
destinationrule.networking.istio.io/hostname-lb-edge created
virtualservice.networking.istio.io/edgemesh-gateway-svc created
```

**查看 edgemesh-gateway 是否创建成功**

```shell
$ kubectl get gw -n edgemesh-test
NAME               AGE
edgemesh-gateway   3m30s
```

**最后，使用 IP 和 Gateway 暴露的端口来进行访问**

```shell
$ curl 192.168.0.211:12345
```

## HTTPS 网关

**创建测试密钥文件**

```bash
$ openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=kubeedge.io"
Generating a RSA private key
............+++++
.......................................................................................+++++
writing new private key to 'tls.key'
-----
```

**根据密钥文件创建 Secret 资源对象**

```bash
$ kubectl create secret tls gw-secret --key tls.key --cert tls.crt -n edgemesh-test
secret/gw-secret created
```

**创建绑定了 Secret 的 Gateway 资源对象和路由规则 VirtualService**

```bash
$ kubectl apply -f examples/hostname-lb-random-gateway-tls.yaml
pod/hostname-lb-edge2 created
pod/hostname-lb-edge3 created
service/hostname-lb-svc created
gateway.networking.istio.io/edgemesh-gateway configured
destinationrule.networking.istio.io/hostname-lb-edge created
virtualservice.networking.istio.io/edgemesh-gateway-svc created
```

**最后，使用证书进行 HTTPS 访问**

```bash
$ curl -k --cert ./tls.crt --key ./tls.key https://192.168.0.129:12345
```
