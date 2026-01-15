package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
    var count int64
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // 原子加 1，绝对线程安全，且比 Mutex 快得多
            atomic.AddInt64(&count, 1)
        }()
    }
    wg.Wait()
    fmt.Println("Count:", count) // 1000
}
