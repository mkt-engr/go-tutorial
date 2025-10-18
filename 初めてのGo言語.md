# １章

下記のような`%s`に割り当てる変数が明示されていない場合に`go vet`を実行すると構文エラーが出る。

```go
fmt.Println("Hello, Go! %s!\n")
```

下記のようにするとよい

```sh
go vet ./...
```

# ２章

# 参考

初めての Go 言語 第 2 版のリポジトリ
https://github.com/mushahiroyuki/lgo2
