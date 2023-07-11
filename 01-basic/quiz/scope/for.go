package main

import "fmt"

func main() {
	var c int
	for a, b := 1, 10; a < b; a++ {
		c = a + b
	}
	fmt.Println(c)
	// output: 19
	fmt.Println()
	// 上述代码相当于：
	forUnfold()
}

func forUnfold() {
	var c int
	a, b := 1, 10 // 变量提到前边
	for ; a < b; a++ {
		c = a + b
	}
	fmt.Println(c)
}
