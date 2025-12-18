package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
    // 查看当前系统的 P 数量（通常等于 CPU 核数）
    fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

    var wg sync.WaitGroup
    count := 1000 * 1000 // 100 万

    start := time.Now()
    for i := 0; i < count; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // 假装做点复杂的数学运算
            _ = 1 + 1
        }()
    }
    wg.Wait()
    
    fmt.Printf("Finished %d goroutines in %v\n", count, time.Since(start))
}
