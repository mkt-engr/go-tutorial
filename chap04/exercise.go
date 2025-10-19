package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//問1
	randomSlice := make([]int, 100)
	for i := range 100 {
		randomSlice[i] = rand.Intn(100)
	}

	//問2
	for _, v := range randomSlice {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind!")
		}
	}

	//問3、4
	// var total int
	total := 0
	for i := 0; i < 100; i++ {
		total = total + i
		fmt.Printf("i=%v total=%v\n", i, total)
	}
	fmt.Println("Total:", total)

}
