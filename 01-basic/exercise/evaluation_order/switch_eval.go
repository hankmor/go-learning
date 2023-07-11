package main

import "fmt"

// switch 中表达式的 "惰性求值"

func swexp(n int) int {
	fmt.Println(n)
	return n
}

func main() {
	switch swexp(2) { // 表达式求值：2
	case swexp(1), swexp(2), swexp(3): // 多个表达式，从左到右依次求值，得到 1， 2，此时 2 符合case条件，不再执行 swexp(3)
		fmt.Println("into case1")
		fallthrough // 直接进入下一个 case，不需要求值
	case swexp(4): // 不求值，而是通过 fallthrough 直接进入
		fmt.Println("into case2")
	}
	/*output:
	2
	1
	2
	into case1
	into case2
	*/
}
