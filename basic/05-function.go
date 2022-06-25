package main

import (
	"fmt"
)

func main() {
	// ==============
	// 函数定义
	// ==============

	println("==== 函数定义 ====")

	// ==============
	// 函数调用
	// ==============

	println("==== 函数调用 ====")
	tp := f()
	println(tp)
	// 大括号分组
	{
		println("b")
		println("c")
	}
	println("d")

	// ==============
	// 不定参数
	// ==============

	println("==== 不定参数 ====")

	// ==============
	// 多重返回值
	// ==============

	println("==== 多返回值 ====")

	// ==============
	// 匿名函数与闭包
	// ==============

	closure()
}

type Tp struct {
	Name string
	Age  int
}

func f() *Tp {
	println("a")
	return &Tp{
		Name: "huzhou",
		Age:  20,
	}
}

func closure() {
	// Go的匿名函数是一个闭包
	// 闭包是可以包含自由（未绑定到特定对象）变量的代码块，这些变量不在这个代码块内或者
	// 任何全局上下文中定义，而是在定义代码块的环境中定义。

	j := 5

	// 一个匿名函数，闭包，返回一个函数
	fun := func() func() {
		var i = 10 // 闭包内封闭的变量 i
		return func() {
			fmt.Printf("i = %d, j = %d \n", i, j)
		}
	}()

	fun()
	j += 10
	fun()
}
