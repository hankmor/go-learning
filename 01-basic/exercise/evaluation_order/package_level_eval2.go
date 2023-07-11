package main

import "fmt"

// 带空变量的包级别的变量初始化

var (
	e = g + f // 4 + 2 = 6
	f = fn1() // 先初始化，2
	_ = fn1() // 包级别的空变量同样也会被初始化，执行 fn1，得到 3
	g = fn1() // 后初始化，4
	h = 1     // h 最终为 4
)

func fn1() int {
	h++
	return h
}

func main() {
	fmt.Println(e, f, g, h)
	// output: 6 2 4 4
}
