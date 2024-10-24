## ルートを生やす
openapi.oas3.yml に適当に書く
```code 
ogen -package ogen -target ogen -clean openAPI/openapi.oas3.yml
```

```code 
curl -H "X-Api-Key: hogehoge" 192.168.11.2:8080/view
curl -H "X-Api-Key: hogehoge" 10.9.5.19:8080/view
```
こんな感じのコマンドでアクセスできるようになる。


## フォルダの見方
### ```presentation```: 
controller/ のhandler たちが rooting の登録を行う。基本的にimpl_repository, service, usecase の初期化を行う。
```code
// サーバーの初期設定
hdl, err := ogen.NewServer(
    &controller.HaiHandler{},
    &auth.HaiSecurityHandler{},
)
```
```main.go```のこれがおそらくHaiHandlerに登録されているメソッドを一括で実行してそう。

### ```domain/repository``` <---> ```infrastructure/repository```
impl_repository からrepository に暗黙に型変換を行っている。
実体はHaiHandler を初期化する時に作られたuseRepositoryで、これをservice に渡して適切なサービス、ハンドラーに作り替えるときに型変換が行われている。