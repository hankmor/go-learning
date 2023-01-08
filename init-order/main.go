package main

import (
	"fmt"
	_ "init-order/pkg1" // 引入 pkg1 包，测试 main 包与之的初始化顺序
	_ "init-order/pkg2" // 引入 pkg1 包，测试 main 包与之的初始化顺序
)

// 测试包级别的 init 的执行顺序：
// 1) 依赖包的初始化顺序按照导入的先后顺序进行，先导入的包先初始化
// 2) 包内部先初始化包级常量，再初始化包级变量，再调用 init 方法，即：常量 -> 变量 -> init函数
// 3) 同一个包的同一个文件中的多个 init 方法按照申明顺序一次执行
// 4) 同一个包中，不同文件的 init 方法，按照 go 对文件的编译顺序调用，先编译的先执行
// 5) 不要依赖 init 方法的执行顺序
// 6) main 包初始化同样按照 常量 -> 变量 -> init函数的顺序进行，最后才会调用 main 方法

// 变量定义，先于 const 申明，测试初始化顺序
var (
	_  = constInitCheck()
	v1 = varInitCheck("v1", 4)
	v2 = varInitCheck("v2", 5)
)

// 常量定义，后于 var 申明，测试初始化顺序
const (
	c1 = 1
	c2 = 2
)

func init() {
	fmt.Println("main: invoke init method 1")
}

// 检查常量初始化
func constInitCheck() int {
	fmt.Println("main: invoke constInitCheck...")
	if c1 != 1 { // 尝试捕捉常量未初始化时的状态
		fmt.Println("main: const c1 is init")
	}
	if c2 != 2 {
		fmt.Println("main: const c2 is init")
	}
	return 3
}

// 检查变量初始化
func varInitCheck(name string, v int) int {
	fmt.Printf("main: var %s init with value %d \n", name, v)
	return v
}

func main() {
	fmt.Println("invoke main method")
	/*output:
	pkg1: invoke constInitCheck...
	pkg1: var v1 init with value 1
	pkg1: var v2 init with value 2
	pkg1: invoke init method
	pkg1: invoke init method 2
	pkg2: invoke constInitCheck...
	pkg2: var v1 init with value 1
	pkg2: var v2 init with value 2
	pkg2: invoke init method
	pkg2: invoke init method 2
	main: invoke constInitCheck...
	main: var v1 init with value 4
	main: var v2 init with value 5
	main: invoke init method 1
	main: invoke init method 2
	invoke main method
	*/
}

func init() {
	fmt.Println("main: invoke init method 2")
}
