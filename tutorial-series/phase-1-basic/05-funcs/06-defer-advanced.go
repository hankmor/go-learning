package main

import (
	"fmt"
	"os"
	"time"
)

// deferOrder 演示 defer 的执行顺序（LIFO）
func deferOrder() {
	fmt.Println("=== Defer Order Demo ===")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("4")
	// 输出: 4 3 2 1
}

// deferWithArgs 演示 defer 的参数立即求值
func deferWithArgs() {
	fmt.Println("\n=== Defer Arguments Evaluation ===")
	x := 1
	defer fmt.Println("Deferred x:", x) // x 的值在这里就确定了

	x = 2
	fmt.Println("Current x:", x)
	// 输出: Current x: 2
	//      Deferred x: 1
}

// deferModifyReturn 演示 defer 修改命名返回值
func deferModifyReturn() (result int) {
	defer func() {
		result++ // 修改命名返回值
	}()
	return 5 // result = 5, 然后执行 defer, result 变成 6
}

// deferWithClosure 演示 defer 中的闭包
func deferWithClosure() {
	fmt.Println("\n=== Defer with Closure ===")
	x := 1
	defer func() {
		fmt.Println("Deferred x (closure):", x) // 闭包捕获变量，使用最终值
	}()

	x = 2
	fmt.Println("Current x:", x)
	// 输出: Current x: 2
	//      Deferred x (closure): 2
}

// measureTime 使用 defer 测量函数执行时间
func measureTime(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

// slowOperation 模拟耗时操作
func slowOperation() {
	defer measureTime("slowOperation")() // 注意这里有两个括号

	// 模拟耗时操作
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Operation completed")
}

// readFile 使用 defer 确保资源释放
func readFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close() // 确保函数返回前关闭文件

	// 读取文件内容...
	// 即使发生错误，defer 也会执行
	return nil
}

// deferInLoop 演示 defer 在循环中的陷阱
func deferInLoop() {
	fmt.Println("\n=== Defer in Loop (Wrong) ===")
	// 错误示例：defer 会累积，直到函数返回才执行
	for i := 0; i < 3; i++ {
		defer fmt.Println("Loop defer:", i)
	}
	// 输出: Loop defer: 2
	//      Loop defer: 1
	//      Loop defer: 0
}

// deferInLoopCorrect 正确的做法：使用匿名函数
func deferInLoopCorrect() {
	fmt.Println("\n=== Defer in Loop (Correct) ===")
	for i := 0; i < 3; i++ {
		func() {
			defer fmt.Println("Loop defer (correct):", i)
			// 其他操作...
		}()
	}
	// 每次循环都会立即执行 defer
}

// panicRecover 演示 defer 与 panic/recover
func panicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("About to panic...")
	panic("something went wrong")
	fmt.Println("This will not be printed")
}

// multipleDefers 演示多个 defer 的实际应用
func multipleDefers() error {
	fmt.Println("\n=== Multiple Defers ===")

	// 模拟打开多个资源
	fmt.Println("Opening resource 1")
	defer fmt.Println("Closing resource 1")

	fmt.Println("Opening resource 2")
	defer fmt.Println("Closing resource 2")

	fmt.Println("Opening resource 3")
	defer fmt.Println("Closing resource 3")

	fmt.Println("Doing work...")

	// 资源会按照相反的顺序关闭
	return nil
}

func main() {
	// 1. defer 执行顺序
	deferOrder()

	// 2. defer 参数立即求值
	deferWithArgs()

	// 3. defer 修改返回值
	result := deferModifyReturn()
	fmt.Println("\nResult from deferModifyReturn:", result) // 6

	// 4. defer 中的闭包
	deferWithClosure()

	// 5. 使用 defer 测量时间
	fmt.Println()
	slowOperation()

	// 6. defer 在循环中
	deferInLoop()
	deferInLoopCorrect()

	// 7. defer 与 panic/recover
	fmt.Println()
	panicRecover()
	fmt.Println("Program continues after panic recovery")

	// 8. 多个 defer
	multipleDefers()
}
