package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 使用加锁的方式保证并发的正确性

var cnt = 0

func countAdd(lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()

	cnt++
}
func main() {
	var lock sync.Mutex
	for i := 0; i < 10; i++ {
		go countAdd(&lock)
	}
	// 让程序不退出
	for {
		// atomic.Lock()
		// c := cnt
		// fmt.Println(c)
		// atomic.Unlock()
		runtime.Gosched() // 让出CPU
		if cnt >= 10 {
			break
		}
	}
	fmt.Println(cnt)
}
