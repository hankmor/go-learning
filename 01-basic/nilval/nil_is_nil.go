package main

import (
	"fmt"
	"strings"
)

var a *int

var b any

func main() {
	fmt.Println(a)      // nil
	fmt.Println(b)      // nil
	fmt.Println(a == b) // false，都是nil，为什么这里为false？具体看 https://go.dev/doc/faq#nil_error

	println(a) // 0x0 使用内置的print函数可以打印变量的值
	println(b) // (0x0,0x0)，go内部接口存储分为两个部分，类型T和值V, T和V都是未设置时接口才为nil
	var i = strings.NewReader("hello")
	println(i)

	fmt.Println("assign value...")
	var x *int
	var y any = x
	println(x)        // 0x0
	println(x == nil) // true
	println(y)        // (0x1096460,0x0)
	println(y == nil) // false
	println(x == y)   // true，这里为true

	fmt.Println("return non nil error...")
	ok = true
	e := returnsError()
	fmt.Println(e) // my error: bad error
	println(e)     // (0x10ca7d0,0x11490d0) 类型和值都不为nil
	ok = false
	e = returnsError()
	fmt.Println(e)    // <nil>
	println(e)        // (0x10ca7d0,0x0) 值为 0x0
	println(e == nil) // false
}

type MyError struct {
	error
}

func (m *MyError) Error() string {
	return "my error: " + m.error.Error()
}

var ok = true
var ErrBad = &MyError{fmt.Errorf("bad error")}

func returnsError() error {
	var p *MyError = nil
	if ok {
		p = ErrBad
	}
	return p // Will always return a non-nil error.
}
