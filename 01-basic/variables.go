package main

import "fmt"

// 变量
func main() {
	// ==============
	// 变量申明
	// ==============

	// 变量申明：var 变量名 变量类型
	var v0 rune   // 字符
	var v1 int    // 整型
	var v2 string // 字符串
	// 变量赋值
	v0 = 'A'
	print("v0=", v0, " v1=", v1, " v2=", v2) // v0=65 v1=0 v2=
	println()

	// 可以将多个变量合并申明
	var (
		v3 [2]int                 // 数组
		v4 []int                  // 数组切片
		v5 struct{ i int }        // 结构体
		v6 *int                   // 指针
		v7 map[string]int         // map
		v8 func(a int, b int) int // 匿名函数
	)
	fmt.Print("v3=", v3, " v4=", v4, " v5=", v5, " v6=", v6, " v7=", v7, " v8=", v8)
	// v3=[0 0] v4=[] v5={0} v6=<nil> v7=map[] v8=<nil>

	// ==============
	// 变量初始化
	// ==============

	// ==============
	// 变量赋值
	// ==============

	// ==============
	// 匿名变量
	// ==============
}
