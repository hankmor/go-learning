package goroutine

import (
	"fmt"
	"runtime"
	"sync"
)

func ShowLetters() {
	// 设置给调度器分配的逻辑处理器的数量，这里设置为只有一个处理器执行
	runtime.GOMAXPROCS(1)

	// 申明一个等待组，即一个信号量（计数器），用来控制每个 goroutine 的执行结束（调用 Done 方法则 goroutine 执行结束 WaitGroup 调用 wait 方法
	// 后会等待设置数量的 goroutine 调用 Done，达到数量则继续向后执行，否则阻塞）
	var wg sync.WaitGroup
	// 设置信号量的数量，调用一次 Done 则数量减1，如果数量减为0，则 wait 方法返回，程序可以继续向后执行
	wg.Add(2)

	// 开启一个 goroutine 来执行匿名函数，打印小写字母表
	go func() {
		println("Show lower letter table")
		// 匿名函数执行完成，信号量减1
		defer wg.Done()

		// 打印字母表3次
		for i := 0; i < 3; i++ {
			for char := 'a'; char < 'a'+26; char++ {
				// 打印字母
				fmt.Printf("%c ", char)
			}
		}
		println()
	}()

	// 在开启一个 goroutine 来执行匿名函数，打印大写字母表
	go func() {
		println("Show upper letter table")
		defer wg.Done()

		for i := 0; i < 3; i++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		println()
	}()

	println("Waiting to finish")
	// 设置数量的计数器没有减为 0，则阻塞，否则 wait 方法从阻塞返回
	wg.Wait()

	println("Termination program")
}
