# works-hai-backend

## 実行手順

1. go をインストール
2. goimports をインストール：

```
go install golang.org/x/tools/cmd/goimports@latest
```

3. トップディレクトリ下に image ディレクトリを手動で作成
4. （lint 関係消して）make
5. config ディレクトリ下に、Haruma から渡される private ディレクトリを配置
6. （eng.ut-1x に接続した状態で）ipconfig して、[自分の ip アドレス]:8080 を config/private/config.yaml の dev に書く
7. config/private/cors.json の ip も修正する
8. ./main.exe でサーバーが立つ

## 注意

CORS は Firebase 側にも反映させないといけない

現在の設定を確認する方法（Firebase にログイン必要）：

```
gsutil cors get gs://*****.appspot.com
```

cors.json があるディレクトリで ↓ を実行すると設定が反映される：

```
gsutil cors set cors.json gs://*****.appspot.com
```
