// 申明main包，Go的main函数必须放到main包中
package main

import (
	"fmt"
	"log"
	"rsc.io/quote"
	// 自定义模块
	"koobyte.com/greetings"
)

func main() {
	fmt.Println("Hello Go!")
	/*
		使用额外的模块，rsc.io/quote模块收集了很多精辟的谚语
		1、搜索模块：https://pkg.go.dev/
		2、引入模块
		3、解析模块依赖命令：go mod tidy，模块查找慢需要设置代理:https://goproxy.io/zh/
		4、执行名； go run .
	*/
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
	/*
		输出：
		Hello Go!
		Don't communicate by sharing memory, share memory by communicating.
		Hello, world.
		I can eat glass and it doesn't hurt me.
		If a program is too slow, it must have a loop.
	*/

	// 使用自定义模块

	/*
		1、模块未发布，而是存储在本地，需要使用go mod命令重新映射：
		go mod edit -replace koobyte.com/greetings=../greetings
		这个命令告诉go从本地目录加载模块，然后会在go.mod中生成一条映射指令：
		replace koobyte.com/greetings => ../greetings
		2、然后再使用go mod tidy 命令加载模块，此时会打印
		go: found koobyte.com/greetings in koobyte.com/greetings v0.0.0-00010101000000-000000000000
		后边的版本是系统自动生成的
	*/
	message := greetings.Hello("koobyte.com")
	fmt.Println(message)

	// 调用带异常信息的方法，该方法返回两个值

	message, err := greetings.Hello1("")
	Print(message, err)

	message1, err1 := greetings.Hello1("Sam")
	Print(message1, err1)
}

func Print(msg string, err error) {
	// 日志前缀
	log.SetPrefix("greetings: ")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(msg)
	}
}
