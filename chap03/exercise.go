package main

import "fmt"

func main() {
	// 構造体の定義
	type person struct {
		name string // 名前
		age  int    // 年齢
		pet  string // ペット
	}

	// ゼロ値で初期化
	var fred person
	bob := person{} // 全フィールドがゼロ値で初期化される
	fmt.Println("fred:", fred)
	fmt.Println("bob:", bob)

	// 順番による初期化
	julia := person{
		"ジュリア", // name
		40,       // age
		"cat",    // pet
	}
	fmt.Println("julia:", julia)

	// フィールド名を指定した初期化（推奨）
	beth := person{
		age:  30,
		name: "ベス",
	}
	fmt.Println("beth:", beth)
}
