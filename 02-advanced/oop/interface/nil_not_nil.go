package main

import "fmt"

type MyError struct {
	error
}

func returnError() error {
	var err *MyError = nil
	return err
}

func main() {
	var err *MyError = nil
	fmt.Println(err == nil) // true

	// 通过方法返回的 nil 指针在这里判断却不是 nil
	// 原因：返回的数据指针为 nil，但是有非 nil 的类型信息()*MyError，所以整体上不为 nil
	// 具体详细信息需要弄明白：接口的底层内部表示(iface、eface)
	var err1 = returnError()
	// 通过内置的 printXx 等函数可以输出接口的内部表示, nil 的内部表示为 (0x0,0x0)
	println("err1:", err1)   // err1: (0x10c6368,0x0)，前边为类型地址，后边为数据地址
	println("nil:", err)     // nil: 0x0
	fmt.Println(err1 == nil) // false
}
