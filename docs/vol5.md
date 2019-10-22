# vol5

## protobuf

```
$ protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/user/user.proto
```

### I

- protoファイルが存在するディレクトリ

### go_out

- grcpでbufferファイルを作成するので `plugin=grpc`
- `source_relative` で実行場所からの相対パス配下に 引数で渡したファイルを作成する
  - ./proto/user/user.pb.go


## k8s

### 基本のキ

https://ucwork.hatenablog.com/entry/2019/02/27/205709

```
# よくやる
$ kubectl create secret generic mysql-pass --from-literal=password=password
```

### serviceの理解

http://chidakiyo.hatenablog.com/entry/2018/09/10/Kubernetes_NodePort_vs_LoadBalancer_vs_Ingress%3F_When_should_I_use_what%3F_%28Kubernetes_NodePort_%E3%81%A8_LoadBalancer_%E3%81%A8_Ingress_%E3%81%AE%E3%81%A9%E3%82%8C%E3%82%92%E4%BD%BF%E3%81%86

### ingress理解

https://kubernetes.io/docs/concepts/services-networking/ingress/

### minikubeはローカルイメージを参照しない

https://www.sambaiz.net/article/151/

### minikubeのメモリはデフォルトで2GB、CPUも...

https://qiita.com/loftkun/items/7400d5ae0e7b1fb7d26e
https://qiita.com/loftkun/items/41ecde082778254de79a

### kubeconfigを環境変数に

https://qiita.com/shoichiimamura/items/91208a9b30e701d1e7f2

```
$ export KUBECONFIG=$KUBECONFIG:/Users/originbenntou/.kube/config
```

### pv pvc

こいつが基本
https://qiita.com/toshihirock/items/0c91bbedf0e144acf6fc

オプションとかを解説
https://cstoku.dev/posts/2018/k8sdojo-12/
