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
var x = [3]int
var x = [3]int{10, 20, 30}    // 初期化
var x = [...]int{10, 20, 30}  // ...で任意の数
```

これら 2 つは型が違う。
配列の大きさを指定するときに変数は使えない。
配列の型は実行時ではなくコンパイル時に決定できる必要がある。

```go
var x = [3]int
var y = [4]int
```

配列はこう言った制限があるためにあまり使われないらしい
代わりにスライスを使う

## スライス

スライスは可変長の配列のイメージ

スライスは長さを指定しない

```go
// 配列
var x = [...]int{10, 20, 30}  // ...で任意の数
// スライス
var x = []int{10, 20, 30}  // 長さを指定しない
```

大きさを指定する場合は make を使う

```go
x := make([]int, 5)
```

## copy

スライスを分割した時に同じメモリを使わないようにする

```go
x := []int{1, 2, 3, 4}
y := make([]int, 4)
num := copy(y, x)  // xからyにコピーする
fmt.Println(y, num)  // [1 2 3 4] 4
```

## マップ

```go
// これは書き込みができない
var nilMap map[string]int  // string→intのマップ。初期値はnil
nilMap["abc"] = 1          // これはエラーになる
```

こちらは書き込みができる

```go
totalWins := map[string]int
totalWins["ライターズ"] = 1
```

### マップの比較

`maps` を利用する

```go
maps.Equal(m, n)
```

## 構造体 struct

```go
type person struct {
 name string  // 名前
 age  int     // 年齢
 pet  string  // ペット
}
```

この構造体を持つ変数を宣言できる

```go
var fred person
bob := person{}  // 全フィールドがゼロ値で初期化される
```

初期化するには２種類の方法がある

順番が大事なやつ。使いにくそう。

```go
julia := person{
 "ジュリア",  // name
 40,        // age
 "cat",     // pet
}
```

またはこんな書き方。こっちの方が良さそう。

```go
beth := person{
 age:  30,
 name: "ベス",
}
```

# ４章 ブロック、シャドーイング、制御構造

## if 文

スコープが TS とは違う。
if 文で値を宣言できる。

```go
if n := rand.Intn(10); n == 0 {
  fmt.Println("小さすぎます:", n)
} else if n > 5 {
  fmt.Println("大きすぎます:", n)
} else {
  fmt.Println("いい感じの数です:", n)
}
```

## for 文

3 つの要素はいずれも省略できる

```go
for i := 0; i < 10; i++ {
 // ループ処理
}
```

### for-range ループ

文字列、配列、スライス、マップで使える。

```go
evenVals := []int{2, 4, 6, 8, 10, 12}
for i, v := range evenVals {
 fmt.Println(v, i)
}
```

## switch

go の switch は break を書かない
その代わり次のケース文を実行する`fallthrough`というのがある

# ５章 関数

関数の基本形
TS と違って返り値の型にコロンはつけない

```go
func  div(num int , denom int) int{

}
```

同じ型の引数が連続する場合は次のように最後以外の変数の型指定を省略できる

```go
func calculate(x, y int, message string, a, b float64)
```

## 可変長の変数

TS とは違い、変数ではなく型に`...`をつける

```go
// 可変長引数の例 - 整数の合計を計算
func sum(nums ...int) int {
  total := 0
   for _, num := range nums {
    total += num
   }
  return total
}
```

使い方はこんな感じ。
使う時は変数の後に`...`を入れる

```go
numbers := []int{5, 10, 15, 20}
sum(numbers...)  // ...でスライスを展開して渡す
```

## 複数の戻り値

Go の関数は複数の値を返すことができる。エラーハンドリングでよく使われる。
返り値が 2 つの場合、それを受け止める変数も 2 つ必要になる。TS みたいにオブジェクトとしてまとめて 1 つに受け取ることはできない
戻り値を無視したい場合は、`_`と変数宣言する

```go
// 除算の結果とエラーを返す
func divide(a, b float64) (float64, error) {
 if b == 0 {
  return 0, fmt.Errorf("ゼロで除算できません")
 }
 return a / b, nil
}

// 使い方
result, err := divide(10, 2)
if err != nil {
 fmt.Println("エラー:", err)
} else {
 fmt.Println("結果:", result)
}
```

### 名前付き戻り値

戻り値に名前をつけることができる。名前付き戻り値を使うと、`return`だけで値を返せる。

```go
func minMax(numbers ...int) (min, max int) {
 if len(numbers) == 0 {
  return 0, 0
 }
 min = numbers[0]
 max = numbers[0]
 for _, num := range numbers {
  if num < min {
   min = num
  }
  if num > max {
   max = num
  }
 }
 return // 名前付き戻り値は return だけでOK
}

// 使い方
min, max := minMax(3, 7, 2, 9, 1, 5)
fmt.Printf("最小値: %d, 最大値: %d\n", min, max)
```

## 関数型の宣言

```go
type opFuncType func(int,int) int
var opMap = map[string]onFuncType{
 "+":add
}
```

## 無名関数

```go
func main(){
 f := func(j int){
  fmt.Println("無名関数の中で",j,"を出力")
 }
}

for i := 0; i<5; i++{
 f(i)
}
```

## クロージャ

```go
func main(){
 a := 20
 f := func(){
  fmt.Println(a)
 a=30
 }
 f() // 20
 fmt.Println(a) //30
}
```

## defer

ファイルやネットワーク接続といった一時的なりシースを作成することがある。このリソースはクリーンアップが必要。
関数が正常終了・以上終了するに関わらず必ず解放する必要がある。

```go
// deferの実用例：ファイルクローズ
func readFileExample() {
 fmt.Println("\nファイル処理の例（疑似コード）")
 fmt.Println("ファイルをオープン")
 defer fmt.Println("defer: ファイルをクローズ")
 fmt.Println("ファイルを読み込み")
 fmt.Println("データを処理")
 // 関数終了時に必ずクローズされる
}
```

# 参考

初めての Go 言語 第 2 版のリポジトリ
https://github.com/mushahiroyuki/lgo2
