package main

import (
	"fmt"
	"math/rand"
	"time"
)

func FlowMain() {
	// ==============
	// 条件
	// ==============

	// 条件语句，go关键字包括：if、else和else if
	// 条件语句不需要使用括号将条件包含起来();
	// 无论语句体内有几条语句，花括号{}都是必须存在的;
	// 左花括号{必须与if或者else处于同一行;
	// 在if之后，条件语句之前，可以添加变量初始化语句，使用;间隔;

	ifFlow()

	// ==============
	// 选择
	// ==============

	switchFlow()

	// ==============
	// 循环
	// ==============

	forFlow()

	// ==============
	// 跳转
	// ==============

	// 调整语句，关键字包括goto
	// 使用goto语句实现for循环的效果
	gotoWhile(5) // ~: 1 2 3 4 5
}

func forFlow() {
	// 循环语句，关键字包括for和range
	// Go语言只有for循环，不支持while和do while，因为for同样可以实现while的效果

	i2 := 5
	// > 基本的循环
	for i := 0; i < i2; i++ {
		fmt.Print(i, " ")
	}
	// ~: 0 1 2 3 4
	fmt.Println()

	// > 无条件的for循环
	for {
		i2++
		if i2 > 10 {
			break // 跳出循环
		}
	}
	fmt.Println(i2) // ~: 11

	// > for循环时使用多重赋值
	array1 := [...]int{0, 1, 2, 3, 4} // 申明一个数组
	// 颠倒顺组元素
	for i, j := 0, len(array1)-1; i < j; i, j = i+1, j-1 { // 使用多重赋值，不支持 i = 0, j = len(array1) - 1这种方式
		array1[i], array1[j] = array1[j], array1[i] // 交换数组的元素
	}
	fmt.Println(array1) // ~: [4 3 2 1 0]

	// > for循环中使用continue
	i3 := 5
	for i := 0; i < i3; i++ {
		if i == 2 {
			continue
		}
		fmt.Print(i, " ")
	}
	// ~: 0 1 3 4
	fmt.Println()

	// > for循环中使用break跳出多层循环
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if i == 1 && j == 2 {
				break
			}
			fmt.Print("i, j=", i, j, " ")
		}
	}
	// ~: j=0 0 i, j=0 1 i, j=0 2 i, j=0 3 i, j=1 0 i, j=1 1 i, j=2 0 i, j=2 1 i, j=2 2 i, j=2 3
	fmt.Println()
TAG:
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if i == 1 && j == 2 {
				break TAG // 直接跳出外层循环
			}
			fmt.Print("i, j=", i, j, " ")
		}
	}
	// ~: j=0 0 i, j=0 1 i, j=0 2 i, j=0 3 i, j=1 0 i, j=1 1
	fmt.Println()
}

func ifFlow() {
	// 随机产生一个数
	a := randInt(20)
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
}

func switchFlow() {
	// go的选择语句关键字包括：switch、case和select，select配合channel使用
	// 注意switch语句中每一个case分支不需要break语句来跳出，go默认执行分支会自动跳出

	i1 := randInt(3)
	// > 常规用法
	switch i1 {
	case 0:
		fmt.Println("i1 == 0") // switch不需要break
	case 1:
		fmt.Println("i1 == 1")
	case 2:
		fmt.Println("i1 == 2")
	default: // default分支可以省略
		fmt.Println("Unknown int i1")
	}

	// > 用法2：switch后可以不跟条件表达式，此时switch相当于简写的if语句
	switch {
	case i1 == 0:
		fmt.Println("i1 is 0.")
	case i1 > 0:
		fmt.Println("i1 is great than 0.")
		// default: // default语句可以省略
	}

	// > 如果需要case语句继续向下执行，需要使用fallthrough关键字
	switch i1 {
	case 0:
		fmt.Println("i1 is 0")
	case 1:
		fallthrough // 符合判断条件时，继续向下执行
	case 2:
		fmt.Println("i1 is great than 0")
	}
}

func gotoWhile(stop int) {
	i := 0
WHILE:
	i++
	fmt.Print(i, " ")
	if i < stop {
		goto WHILE
	}
}

func randInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func ifReturn(b int) int {
	if b > 1 {
		return 1
	} else {
		return b
	}
}
