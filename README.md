# works-hai-backend

実行手順

1. go をインストール
2. goimports をインストール：

```
go install golang.org/x/tools/cmd/goimports@latest
```

3. （lint 関係消して）make
4. config ディレクトリ下に、Haruma から渡される private ディレクトリを配置
5. （eng.ut-1x に接続した状態で）ipconfig して、[自分の ip アドレス]:8080 を config/private/config.yaml の dev に書く
6. ./main.exe でサーバーが立つ
