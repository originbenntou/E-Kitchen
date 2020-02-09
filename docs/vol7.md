# 久しぶりにつけてみて

``` bash
$ minikube delete
$ export KUBECONFIG=$KUBECONFIG:/Users/originbenntou/.kube/config
$ minikube start

$ kubectl create secret generic mysql-pass --from-literal=password=password
$ minikube addons list
$ # ingressの有効化
$ minikube addons enable ingress
```

# TODO

- k8sにnamespaceをつける
    - skaffoldで単一runさせるために必要
    - https://skaffold.dev/docs/references/cli/#skaffold-run
    

