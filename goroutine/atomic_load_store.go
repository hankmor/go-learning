package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	canBreak int64 = 0
	sg1      sync.WaitGroup
)

// BreakFor 使用 atomic 包下的 LoadXxx 和 StoreXxx 方法来同步，LoadXxx 和 StoreXxx 不能被不同的 goroutine 同时执行
func BreakFor() {
	sg1.Add(2)

	go DeadFor("A")
	go DeadFor("B")

	time.Sleep(1 * time.Second)     // 睡眠1s
	atomic.StoreInt64(&canBreak, 1) // 原子的将 canBreak 改为1

	sg1.Wait()
	fmt.Println("All work done.")
}

func DeadFor(name string) {
	defer sg1.Done()
	for {
		fmt.Printf("Doing work: %v\n", name)
		time.Sleep(100 * time.Millisecond) // 睡眠
		// 原子的加载 canBreak 的值，为1时退出循环
		if atomic.LoadInt64(&canBreak) == 1 {
			fmt.Printf("Break work: %v\n", name)
			break
		}
	}
}
