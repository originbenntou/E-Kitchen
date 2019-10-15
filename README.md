E-Kitchen
===

フロントから参照するデータを管理するCMS

## Description

- E-Lunch（現在廃止中）から参照
    - https://github.com/originbenntou/E-Lunch
- まだユーザー認証が完成した段階...
    - https://github.com/originbenntou/E-Kitchen/issues 
- 技術習得を目的にして開発

### 主な技術

- Go
    - go mod
- gRPC
    - protocol buffer
- Kubernetes
    - skaffold
- docker
    - multi stage build
- gorilla/mux
- GORM

## Demo

- クローンして `skaffold run` すればローカルで動きます！！

## Requirement & Install

### skaffold

```shell script
# macos
$ brew install skaffold
```
https://skaffold.dev/docs/getting-started/

### Kubernetes

以下のいずれかの方法でKubernetesをローカルにセットアップ

#### 1.Docker in Mac

- Preference > Kubernetes > enable Kubernetes
    - check and apply

#### 2.Minikube

以下を参考にMinikubeをインストール
https://kubernetes.io/ja/docs/tasks/tools/install-minikube/

#### 起動確認

```shell script
$ kubectl config get-contexts
## 現在使用しているローカルkubernetesを確認

$ kubectl get nodes
## masterノードが起動していることを確認
```

## Usage

## VS. 

## Licence

[MIT]() の予定

## Author

[originbenntou](https://github.com/originbenntou)
