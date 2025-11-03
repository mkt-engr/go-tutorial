package main

import "fmt"

// 1. ポインタの基本
func pointerBasics() {
	fmt.Println("=== ポインタの基本 ===")

	// 通常の変数
	x := 10
	fmt.Println("x の値:", x)
	fmt.Println("x のアドレス:", &x)

	// ポインタの宣言（&でアドレスを取得）
	p := &x
	fmt.Println("p の値（アドレス）:", p)
	fmt.Println("p が指す値（*で逆参照）:", *p)

	// ポインタ経由で値を変更
	*p = 20
	fmt.Println("*p = 20 の後:")
	fmt.Println("x の値:", x)   // 20（元の変数も変わる）
	fmt.Println("*p の値:", *p) // 20
}

// 2. 値渡し vs ポインタ渡し
func modifyValue(x int) {
	x = 100
	fmt.Println("  関数内の x:", x)
}

func modifyPointer(x *int) {
	*x = 100
	fmt.Println("  関数内の *x:", *x)
}

func valueVsPointer() {
	fmt.Println("\n=== 値渡し vs ポインタ渡し ===")

	// 値渡し
	a := 10
	fmt.Println("値渡し前の a:", a)
	modifyValue(a)
	fmt.Println("値渡し後の a:", a) // 10（変わらない）

	// ポインタ渡し
	b := 10
	fmt.Println("\nポインタ渡し前の b:", b)
	modifyPointer(&b)
	fmt.Println("ポインタ渡し後の b:", b) // 100（変わる！）
}

// 3. 構造体とポインタ
type Person struct {
	Name string
	Age  int
}

func birthday(p *Person) {
	p.Age++ // (*p).Age と書いてもいいが、省略可能
}

func structPointer() {
	fmt.Println("\n=== 構造体とポインタ ===")

	person := Person{Name: "太郎", Age: 20}
	fmt.Println("誕生日前:", person)

	birthday(&person)
	fmt.Println("誕生日後:", person)
}

// 4. nilポインタ
func nilPointer() {
	fmt.Println("\n=== nilポインタ ===")

	var p *int
	fmt.Println("p の値:", p) // <nil>

	if p == nil {
		fmt.Println("p は nil です")
	}

	// nilポインタを逆参照するとpanicになる
	// fmt.Println(*p) // これは実行時エラー！

	// 正しい使い方
	x := 10
	p = &x
	if p != nil {
		fmt.Println("p が指す値:", *p)
	}
}

// 5. newでポインタを作成
func newPointer() {
	fmt.Println("\n=== new でポインタを作成 ===")

	// new は指定した型のゼロ値を持つポインタを返す
	p := new(int)
	fmt.Println("new(int) の値:", *p) // 0（ゼロ値）
	fmt.Println("new(int) のアドレス:", p)

	*p = 42
	fmt.Println("代入後の値:", *p)

	// 構造体でも使える
	person := new(Person)
	person.Name = "花子"
	person.Age = 25
	fmt.Println("Person:", *person)
}

// 6. ポインタを使うべき時
func whenToUsePointers() {
	fmt.Println("\n=== ポインタを使うべき時 ===")

	type LargeData struct {
		data [1000]int
	}

	// 値渡し：構造体全体がコピーされる（遅い・メモリを使う）
	processByValue := func(d LargeData) {
		fmt.Println("値渡し: データのコピーを処理")
	}

	// ポインタ渡し：アドレスだけコピー（速い）
	processByPointer := func(d *LargeData) {
		fmt.Println("ポインタ渡し: アドレスだけコピー")
	}

	large := LargeData{}
	processByValue(large)
	processByPointer(&large)
}

func main() {
	pointerBasics()
	valueVsPointer()
	structPointer()
	nilPointer()
	newPointer()
	whenToUsePointers()
}
