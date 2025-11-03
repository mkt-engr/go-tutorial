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

# ６章 ポインタ

## ポインタとは

ポインタは**変数のメモリアドレス**を格納する変数。値そのものではなく、値が保存されている場所を指す。

## 基本的な使い方

```go
// 通常の変数
x := 10

// ポインタの宣言（&でアドレスを取得）
p := &x  // xのアドレスをpに格納

// ポインタから値を取得（*で逆参照）
fmt.Println(*p)  // 10

// ポインタ経由で値を変更
*p = 20
fmt.Println(x)   // 20（元の変数も変わる）
```

**記号の意味：**

- `&` = アドレス演算子（変数のアドレスを取得）
- `*` = 逆参照演算子（ポインタが指す値を取得/設定）

## 値渡し vs ポインタ渡し

```go
// 値渡し（コピーされる）
func modifyValue(x int) {
 x = 100  // コピーを変更するだけ
}

// ポインタ渡し（参照渡し）
func modifyPointer(x *int) {
 *x = 100  // 元の値を変更できる
}

func main() {
 a := 10
 modifyValue(a)
 fmt.Println(a)  // 10（変わらない）

 b := 10
 modifyPointer(&b)
 fmt.Println(b)  // 100（変わる！）
}
```

## 構造体とポインタ

```go
type Person struct {
 Name string
 Age  int
}

func birthday(p *Person) {
 p.Age++  // (*p).Age と書いてもいいが、省略可能
}

func main() {
 person := Person{Name: "太郎", Age: 20}
 birthday(&person)
 fmt.Println(person)  // {太郎 21}
}
```

## nil ポインタ

```go
var p *int  // nil（何も指していない）

if p == nil {
 fmt.Println("p は nil です")
}

// nilポインタを逆参照するとpanicになる
// fmt.Println(*p)  // 実行時エラー！

// 使う前にnilチェック
x := 10
p = &x
if p != nil {
 fmt.Println(*p)  // 安全
}
```

## new でポインタを作成

```go
// new は指定した型のゼロ値を持つポインタを返す
p := new(int)
fmt.Println(*p)  // 0（ゼロ値）

*p = 42
fmt.Println(*p)  // 42
```

## ポインタを使うべき時

**1. 関数で値を変更したい時**

```go
func increment(x *int) {
 *x++
}
```

**2. 大きなデータをコピーしたくない時**

```go
type LargeStruct struct {
 data [1000000]int
}

// ポインタ渡し：速い
func process(s *LargeStruct) {
 // ...
}
```

**3. nil（存在しない）を表現したい時**

```go
var person *Person  // 「まだ存在しない」を表現
if person == nil {
 person = &Person{Name: "太郎"}
}
```

# 7 章 型、メソッド、インターフェース

## 関数とメソッドの違い

### 関数

- 独立して存在する
- 特定の型に紐付かない

```go
// 関数：独立している
func calculateArea(width, height int) int {
 return width * height
}

// 呼び出し
area := calculateArea(10, 5)
```

### メソッド

- 特定の型（レシーバー）に紐付く
- その型の「動作」を定義する

```go
type Rectangle struct {
 Width  int
 Height int
}

// メソッド：Rectangle型に紐付く
func (r Rectangle) Area() int {
 return r.Width * r.Height
}

// 呼び出し
rect := Rectangle{Width: 10, Height: 5}
area := rect.Area()  // より読みやすい
```

## レシーバーの種類

### 値レシーバー

コピーが渡される。元の値は変更されない。

```go
func (r Rectangle) Area() int {
 // rはコピー（元の値は変更されない）
 return r.Width * r.Height
}
```

### ポインタレシーバー

ポインタが渡される。元の値を変更できる。

```go
func (r *Rectangle) Scale(factor int) {
 // rはポインタ（元の値を変更できる）
 r.Width *= factor
 r.Height *= factor
}

// 使用例
rect := Rectangle{Width: 10, Height: 5}
rect.Scale(2)  // {20 10}
```

## どちらを使うべきか

**ポインタレシーバーを使う場合：**

- メソッドでレシーバーを変更する必要がある
- レシーバーが大きな構造体（コピーコストが高い）
- 一貫性のため（同じ型のメソッドは統一する）

**値レシーバーを使う場合：**

- レシーバーを変更しない
- レシーバーが小さい型（int、bool、小さな構造体）
- 不変性を保ちたい

## メソッドの自動変換

Go は自動的に`&`や`*`を補完してくれる。

```go
rect := Rectangle{Width: 5, Height: 3}
rect.Scale(2)  // (&rect).Scale(2) と同じ

pRect := &Rectangle{Width: 7, Height: 4}
area := pRect.Area()  // (*pRect).Area() と同じ
```
