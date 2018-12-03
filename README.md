# quiz02

## how to run

```bash
$ go run main.go -port 8000
```

default port is `8000`

or using realize (task runner)

```bash
$ go get github.com/oxequa/realize
$ realize start
```

access browser

[http://localhost:8000](http://localhost:8000)

## goal

`main.go` の `saveFile` func を完成させて下さい。

`data/` 配下にアップロードされたファイルを保存します。
> `ディレクトリトラバーサル` など対応できればして下さい

アップロードでは複数のファイルが選択されることがあります。(multiple)

アップロードの結果をHTMLに描画する必要は特にありません。

## 参考

[reader など](https://go-tour-jp.appspot.com/methods/21)や
`writer` `File` interfaceなどが参考になると思います。
