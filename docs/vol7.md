# 久しぶりにつけてみて

``` bash
$ minikube delete
$ unset KUBECONFIG
$ export KUBECONFIG=$KUBECONFIG:/Users/originbenntou/.kube/config
$ minikube start

$ kubectl create secret generic mysql-pass --from-literal=password=password
$ minikube addons list
$ # ingressの有効化
$ minikube addons enable ingress
$ skaffold dev
```

# TODO

- k8sにnamespaceをつける
    - skaffoldで単一runさせるために必要
    - https://skaffold.dev/docs/references/cli/#skaffold-run
    

# interface

https://medium.com/since-i-want-to-start-blog-that-looks-like-men-do/%E5%88%9D%E5%BF%83%E8%80%85%E3%81%AB%E9%80%81%E3%82%8A%E3%81%9F%E3%81%84interface%E3%81%AE%E4%BD%BF%E3%81%84%E6%96%B9-golang-48eba361c3b4

何でも入る型としての使い方で、関数を集めた使い方と2つある（まず）

```go
type Person interface {
    getEmail() string
}
```

PersonインターフェースはgetEmail関数を持っていればどんな構造体も受け入れる

```go
	var s, t Person
	s = Student{
		Name:   "Yamada",
		Number: 999,
		Grade:  5,
	}
	t = Teacher{
		Name: "Tsubomi",
	}
```

こうすると、似たような構造体をひとつのPersonインターフェースにまとめることができる

```go
func sendEmail(p Person) (context string) {
	from := p.getEmail()
	context = `
  送信元 : ` + from + `
  これはテスト用のメールです。
  よろしくお願いします。
  `
	return context
}
```

引数にPersonインターフェースを渡せば、Student,Teacher構造体どちらも受け入れる

```go
type EmptyInter interface{
}
```

このように定義することで、何でも入るインターフェースを用意できる
=>何でも入る型としての使い方
