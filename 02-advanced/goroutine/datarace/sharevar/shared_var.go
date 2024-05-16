package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

// 据竞争也可能发生在原始类型的变量上（ bool 、 int 、 int64 等）
// 导致问题的原因在于内存访问的非原子性、编译器优化的干扰或访问处理器内存的重新排序问题
// 解决办法是加锁或者使用 atomic 包，

type Watchdog struct{ last int64 }

func (w *Watchdog) KeepAlive() {
	w.last = time.Now().UnixNano() // First conflicting access.
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			// Second conflicting access.
			if w.last < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}

func (w *Watchdog) SafeKeepAlive() {
	atomic.StoreInt64(&w.last, time.Now().UnixNano())
}

func (w *Watchdog) SafeStart() {
	go func() {
		for {
			time.Sleep(time.Second)
			if atomic.LoadInt64(&w.last) < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}
