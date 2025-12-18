package main

import (
	"fmt"
	"runtime"
)

func main() {
    // 分配一大块内存
    a := make([]byte, 100*1024*1024) // 100MB
    _ = a

    fmt.Println("Before GC")
    printMemStats()

    // 手动强制 GC
    runtime.GC()

    fmt.Println("\nAfter GC")
    printMemStats()
}

func printMemStats() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Alloc = %v MiB", m.Alloc / 1024 / 1024)
}
