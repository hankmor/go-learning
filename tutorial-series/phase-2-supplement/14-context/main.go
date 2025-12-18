package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
    // 1. 创建一个带超时的 Context
    // 规定任务必须在 2 秒内完成，否则 ctx 会收到取消信号
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    // 关键：函数结束时必须调用 cancel，防止内存泄漏
    defer cancel()

    fmt.Println("Start working...")
    
    // 2. 将 ctx 传递给耗时任务
    doSlowWork(ctx)
}

func doSlowWork(ctx context.Context) {
    // 模拟一个需要 3 秒才能完成的任务
    // Select 会等待 done 信号或者 work 完成，以此来实现超时抢占
    select {
    case <-time.After(3 * time.Second): // 模拟业务逻辑
        fmt.Println("Work done successfully!")
    case <-ctx.Done(): // 监听 Context 的取消信号
        // 如果超时了，ctx.Err() 会返回 DeadlineExceeded
        fmt.Println("Work cancelled:", ctx.Err())
    }
}
