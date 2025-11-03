package main

import "fmt"

// 同じ型の引数が連続する場合、型指定を省略できる例
func greet(firstName, lastName string, age int) string {
	return fmt.Sprintf("%s %sさん、%d歳", firstName, lastName, age)
}

// 複数の同じ型グループがある例
func calculate(x, y int, message string, a, b float64) string {
	sum := x + y
	product := a * b
	return fmt.Sprintf("%s: 整数の合計=%d, 小数の積=%.2f", message, sum, product)
}

// string, int, string の例
func createProfile(firstName, lastName string, age int, city string) string {
	return fmt.Sprintf("%s %sさん（%d歳）は%s在住", firstName, lastName, age, city)
}

// 可変長引数の例 - 整数の合計を計算
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// 可変長引数と通常の引数を組み合わせた例
func greetAll(greeting string, names ...string) string {
	result := greeting + ": "
	for i, name := range names {
		if i > 0 {
			result += ", "
		}
		result += name
	}
	return result
}

// 複数の戻り値 - 除算の結果とエラーを返す
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("ゼロで除算できません")
	}
	return a / b, nil
}

// 名前付き戻り値
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

// deferの例 - 関数終了時に実行される
func deferExample() {
	fmt.Println("1. 関数開始")
	defer fmt.Println("4. defer: 関数終了時に実行される")
	fmt.Println("2. 処理中...")
	fmt.Println("3. 処理完了")
	// 関数終了時にdeferが実行される
}

// 複数のdeferはLIFO（後入れ先出し）で実行される
func multipleDeferExample() {
	fmt.Println("開始")
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	fmt.Println("終了")
	// 実行順序: defer 3 → defer 2 → defer 1
}

// deferの実用例：ファイルクローズ
func readFileExample() {
	fmt.Println("\nファイル処理の例（疑似コード）")
	fmt.Println("ファイルをオープン")
	defer fmt.Println("defer: ファイルをクローズ")
	fmt.Println("ファイルを読み込み")
	fmt.Println("データを処理")
	// 関数終了時に必ずクローズされる
}

func main2() {
	// firstName と lastName はどちらも string なので省略可能
	result1 := greet("太郎", "山田", 25)
	fmt.Println(result1)

	// x, y は int、a, b は float64 でそれぞれ省略
	result2 := calculate(10, 20, "計算結果", 3.5, 2.0)
	fmt.Println(result2)

	// string, int, string のパターン
	result3 := createProfile("花子", "佐藤", 30, "東京")
	fmt.Println(result3)

	// 可変長引数 - 任意の数の引数を渡せる
	fmt.Println("sum(1, 2, 3):", sum(1, 2, 3))
	fmt.Println("sum(10, 20, 30, 40, 50):", sum(10, 20, 30, 40, 50))
	fmt.Println("sum():", sum()) // 引数なしでもOK

	// スライスを展開して渡す
	numbers := []int{5, 10, 15, 20}
	fmt.Println("sum(numbers...):", sum(numbers...))

	// 可変長引数と通常の引数の組み合わせ
	fmt.Println(greetAll("こんにちは", "太郎", "花子", "次郎"))
	fmt.Println(greetAll("Hello", "Alice", "Bob"))

	// 複数の戻り値
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// ゼロ除算のエラー
	_, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("エラー:", err2)
	}

	// 名前付き戻り値
	min, max := minMax(3, 7, 2, 9, 1, 5)
	fmt.Printf("最小値: %d, 最大値: %d\n", min, max)

	// deferの例
	fmt.Println("\n--- deferの基本 ---")
	deferExample()

	fmt.Println("\n--- 複数のdefer ---")
	multipleDeferExample()

	// deferの実用例
	readFileExample()
}
