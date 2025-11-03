package main

import "fmt"

// 構造体の定義
type Rectangle struct {
	Width  int
	Height int
}

type Circle struct {
	Radius float64
}

// ========================================
// 関数：独立して存在する
// ========================================

func calculateRectangleArea(width, height int) int {
	return width * height
}

func calculateCircleArea(radius float64) float64 {
	return 3.14159 * radius * radius
}

// ========================================
// メソッド：特定の型に紐付く
// ========================================

// 値レシーバー：元の値は変更されない
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

// 値レシーバー：周囲の長さを計算
func (r Rectangle) Perimeter() int {
	return 2 * (r.Width + r.Height)
}

// ポインタレシーバー：元の値を変更できる
func (r *Rectangle) Scale(factor int) {
	r.Width *= factor
	r.Height *= factor
}

// ポインタレシーバー：幅と高さを設定
func (r *Rectangle) SetDimensions(width, height int) {
	r.Width = width
	r.Height = height
}

// Circleのメソッド
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c *Circle) SetRadius(radius float64) {
	c.Radius = radius
}

// ========================================
// 値レシーバー vs ポインタレシーバーの比較
// ========================================

func (r Rectangle) TryModifyValue() {
	r.Width = 999 // コピーを変更するだけ
	fmt.Println("  値レシーバー内:", r.Width)
}

func (r *Rectangle) TryModifyPointer() {
	r.Width = 999 // 元の値を変更する
	fmt.Println("  ポインタレシーバー内:", r.Width)
}

func main() {
	fmt.Println("=== 関数 vs メソッド ===")

	rect := Rectangle{Width: 10, Height: 5}

	// 関数の呼び出し：独立している
	area1 := calculateRectangleArea(rect.Width, rect.Height)
	fmt.Println("関数で計算した面積:", area1)

	// メソッドの呼び出し：型に紐付いている
	area2 := rect.Area()
	fmt.Println("メソッドで計算した面積:", area2)

	fmt.Println("\n=== メソッドの利点 ===")

	// メソッドは読みやすい
	fmt.Println("面積:", rect.Area())
	fmt.Println("周囲の長さ:", rect.Perimeter())

	// 関数だと引数が増える
	fmt.Println("面積:", calculateRectangleArea(rect.Width, rect.Height))

	fmt.Println("\n=== ポインタレシーバー ===")

	// ポインタレシーバーで値を変更
	fmt.Println("拡大前:", rect)
	rect.Scale(2)
	fmt.Println("拡大後:", rect)

	rect.SetDimensions(20, 10)
	fmt.Println("変更後:", rect)

	fmt.Println("\n=== 値レシーバー vs ポインタレシーバー ===")

	rect2 := Rectangle{Width: 100, Height: 50}
	fmt.Println("元の値:", rect2.Width)

	// 値レシーバー：元の値は変わらない
	rect2.TryModifyValue()
	fmt.Println("値レシーバー後:", rect2.Width) // 100のまま

	// ポインタレシーバー：元の値が変わる
	rect2.TryModifyPointer()
	fmt.Println("ポインタレシーバー後:", rect2.Width) // 999に変わる

	fmt.Println("\n=== 異なる型のメソッド ===")

	circle := Circle{Radius: 5.0}
	fmt.Println("円の面積:", circle.Area())

	circle.SetRadius(10.0)
	fmt.Println("半径変更後の面積:", circle.Area())

	fmt.Println("\n=== メソッドの自動的なポインタ変換 ===")

	// Goは自動的に&や*を補完してくれる
	rect3 := Rectangle{Width: 5, Height: 3}
	rect3.Scale(3) // (*rect3).Scale(3) と同じ

	pRect := &Rectangle{Width: 7, Height: 4}
	fmt.Println("ポインタでも値レシーバーを呼べる:", pRect.Area()) // (*pRect).Area() と同じ
}
