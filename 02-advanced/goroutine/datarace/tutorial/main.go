package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	go func() { i++ }()                // 协程写
	fmt.Println(i)                     // 主程读
	time.Sleep(100 * time.Millisecond) // 稍微等一下让 goroutine 也有机会跑（虽然这里主要是演示 race，不是逻辑正确）
}

// 运行: go run -race race_tutorial.go
