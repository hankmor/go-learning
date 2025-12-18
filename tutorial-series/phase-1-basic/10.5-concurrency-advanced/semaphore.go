package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	fmt.Println("=== Semaphore Demo ===")

	// 创建一个权重为 3 的信号量（最多 3 个并发）
	sem := semaphore.NewWeighted(3)
	ctx := context.Background()

	// 启动 10 个任务
	for i := 1; i <= 10; i++ {
		// 获取信号量（如果已满，会阻塞）
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Println("Failed to acquire semaphore:", err)
			break
		}

		go func(id int) {
			defer sem.Release(1) // 释放信号量

			fmt.Printf("Task %d started\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Task %d completed\n", id)
		}(i)
	}

	// 等待所有任务完成
	// 通过获取全部权重来确保所有任务都已释放
	sem.Acquire(ctx, 3)
	fmt.Println("All tasks completed")
}
