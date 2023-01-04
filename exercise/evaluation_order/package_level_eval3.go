package main

import "fmt"

// 处于同一行申明的变量，无论那个变量先初始化，其他变量都会同时被初始化

var (
	a1     = c1    // 先初始化 c1
	b1, c1 = fn2() // b1, c1 处于同一行申明，则会一起初始化, 故 b1 先与 c1 初始化
	d1     = 1
)

func fn2() (int, int) {
	d1++
	return d1, d1 + 1
}

func main() {
	fmt.Println(a1, b1, c1, d1)
	// output: 3 2 3 2
}
