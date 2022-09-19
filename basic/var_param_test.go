package main

import (
	"fmt"
	"testing"
)

// 测试可变参数传递nil和是否解引用，观察是否导致可变参数变化
func TestVariableParam(t *testing.T) {
	// a(nil) // 传递nil，则可变参数的slice不为nil，而且长度为1，说明可变参数可以将nil包装仅可变参数对应的slice
	a() // 不传递任何参数，则可变参数为nil
}

func a(args ...any) {
	fmt.Println(args == nil) // false
	if args != nil {
		fmt.Println(len(args)) // 1
		for _, arg := range args {
			fmt.Println(arg)
		}
	}
	b(args) // 如果args为nil，则调用b方法后b方法的可变参数长度为1，又将nil包装仅可变参数的slice中，此时需要解引用
	// b(args...) // 如果args为nil，解引用后不会导致b的可变参数长度为1，而是nil
}

func b(args ...any) {
	fmt.Println(args == nil) // false
	if args != nil {
		fmt.Println(len(args)) // 1
		for _, arg := range args {
			fmt.Println(arg)
		}
	}
}
