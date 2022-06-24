package oop

import (
	"bytes"
	"fmt"
	"io"
)

// 类型与类型值的奇怪错误

// debug改为false，异常：panic: runtime error: invalid memory address or nil pointer dereference
const debug = false

func Trap() {
	test()
}

func test() {
	// *bytes.Buffer 实现了 io.Writer 接口
	var buf *bytes.Buffer // 错误，指针类型不为nil值却为nil
	// var buf io.Writer // 正确的写法，此时类型为nil，值也为nil
	if debug {
		buf = new(bytes.Buffer) // 启用输出收集
	}
	fmt.Printf("%T，%v, %X, %t\n", buf, buf, &buf, buf != nil) // 打印类型和值，*bytes.Buffer，<nil>, C00000E028
	f(buf)                                                    // 微妙的错误：buf不为nil，但是其值为nil，导致f函数中w.Write出错
}

// *bytes.Buffer 赋值给 w 后，w的动态类型不为nil，值为nil，导致调用出错，此时 w 是一个包含空指针的非空接口
func f(w io.Writer) {
	fmt.Println(w != nil) // true
	if w != nil {
		_, err := w.Write([]byte("something"))
		if err != nil {
			panic(err)
		}
	}
}
