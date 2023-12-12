package main

import (
	"fmt"
	"runtime"
	"sync"
)

// ==================== 测试多线程模式下 cpu 的利用率 ====================

// 编写一个死循环，消耗 cpu
func loop() {
	x := 0
	for {
		x = x ^ 1
	}
}

func main() {
	//获取 cpu 核心数
	cpuNum := runtime.NumCPU()
	fmt.Printf("cpu number: %d\n", cpuNum)
	//启动多个 gorutine，然后观察 cpu 消耗情况
	wg := sync.WaitGroup{}
	wg.Add(cpuNum)
	for i := 0; i < cpuNum; i++ {
		fmt.Printf("gorutine %d is running...\n", i)
		go loop()
	}
	wg.Wait()
}

// 通过 Activity Monitor 观察我的 macOS 上的 cpu 消耗情况，8 核心cpu资源消耗达到 700% 以上，说明 gorutine 在多线程下可以充分利用
// cpu资源
