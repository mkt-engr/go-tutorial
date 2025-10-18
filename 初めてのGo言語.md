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

# 3 章 合成型

## 配列

```go
var x =[3]int
var x = [3]int{10,20,30} //初期化
var x = [...]int{10,20,30} //...で任意の数
```

これら 2 つは型が違う。
配列の大きさを指定するときに変数は使えない。
配列の型は実行時ではなくコンパイル時に決定できる必要がある。

```go
var x =[3]int
var y =[4]int
```

配列はこう言った制限があるためにあまり使われないらしい
代わりにスライスを使う

## スライス

スライスは可変長の配列のイメージ

スライスは長さを指定しない

```go
//配列
var x = [...]int{10,20,30} //...で任意の数
//スライス
var x = []int{10,20,30} //...で任意の数

```

大きさを指定する場合は make を使う

```go
x := make([]int ,5)
```

## copy

スライスを分割した時に同じメモリを使わないようにする

```go
x := []int{1,2,3,4}
y := make([]int,4)
num :=  copy(y,x) xからyにコピーする
fmt.Println(y,num) // [1 2 3 4] 4
```

## マップ

```go
// これは書き込みができない
var nilMap map[string]int // string→intのマップ。初期値はnil
nilMap["abc"]=1 //これはエラーになる
```

こちらは書き込みができる

```go
totalWins := map[string]int
totalWins["ライターズ"] = 1
```

### マップの比較

`maps` を利用する

```go
maps.Equal(m,n);
```

## 構造体 struct

```go
type person struct {
  name string //名前
  age int 　// 年齢
  pet string //ペット
}
```

この構造体を持つ変数を宣言できる

```go
var fred person
bob := person{} //全フィールドがゼロ値で初期化される
```

初期化するには２種類の方法がある

順番が大事なやつ。使いにくそう。

```go
julia := person{
  "ジュリア",//name
  40,//age
  cat,//pet
}
```

またはこんな書き方。こっちの方が良さそう。

```go
beth := person{
  age:30 ,
  name:"ベス"
}
```

# 参考

初めての Go 言語 第 2 版のリポジトリ
https://github.com/mushahiroyuki/lgo2
