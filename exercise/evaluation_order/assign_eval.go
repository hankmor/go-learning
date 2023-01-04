package main

import "fmt"

// 赋值语句的求值顺序

func main() {
	n0, n1 := 1, 2
	n0, n1 = n0+n1, n0 // n0, n1 值都已经确定，直接求值
	fmt.Println(n0, n1)
	// output: 3 1
}
