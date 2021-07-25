package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// ==============
	// 条件
	// ==============

	// 条件语句，go关键字包括：if、else和else if
	// 条件语句不需要使用括号将条件包含起来();
	// 无论语句体内有几条语句，花括号{}都是必须存在的;
	// 左花括号{必须与if或者else处于同一行;
	// 在if之后，条件语句之前，可以添加变量初始化语句，使用;间隔;

	// 随机产生一个数
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(20)
	fmt.Println(a)
	if a > 10 { // if语句不需要圆括号，但是必须要大括号
		fmt.Println("a > 10")
	} else if a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a < 10")
	}

	// 条件语句可以初始化，分好分隔
	if b := 0; b < 1 {
		fmt.Println("b < 1")
	}

	r := ifReturn(1)
	fmt.Println(r) // ~: 1

	// ==============
	// 选择
	// ==============

	// go的选择语句关键字包括：switch、case和select，select配合channel使用

	// ==============
	// 循环
	// ==============

	// 循环语句，关键字包括for和range

	// ==============
	// 跳转
	// ==============

	// 调整语句，关键字包括goto

}

func ifReturn(b int) int {
	if b > 1 {
		return 1
	} else {
		return b
	}
}
