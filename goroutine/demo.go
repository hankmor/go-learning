package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Demo() {
	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU())
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 睡眠1s
			time.Sleep(time.Second)
			// 直接使用变量i，全部输出的i为10
			fmt.Printf("success: %d\n", i)
		}()
	}
	wg.Wait()
	println("finished")
}

func Demo1() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		ni := i // 拷贝到新的变量中
		go func() {
			defer wg.Done()
			// 睡眠1s
			time.Sleep(time.Second)
			// 全部输出的i为10
			fmt.Printf("success: %d\n", ni)
		}()
	}
	wg.Wait()
	println("finished")
}
