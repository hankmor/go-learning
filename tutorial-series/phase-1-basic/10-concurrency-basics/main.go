package main

import (
	"fmt"
	"time"
)

// 一个模拟耗时任务的函数
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d started job %d\n", id, j)
        time.Sleep(time.Second) // 模拟耗时 1 秒
        fmt.Printf("Worker %d finished job %d\n", id, j)
        results <- j * 2 // 将结果发送回 results 通道
    }
}

func main() {
    // 创建两个通道：任务通道和结果通道
    // 设置缓冲区大小为 10，防止阻塞
    jobs := make(chan int, 10)
    results := make(chan int, 10)

    // 启动 3 个 Goroutine (Worker)
    // 它们会并发地从 jobs 通道抢任务做
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // 发送 5 个任务
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs) // 关闭任务通道，告知 Worker 没有新任务了

    // 接收并打印结果
    for a := 1; a <= 5; a++ {
        <-results
    }
    
    fmt.Println("All jobs finished!")
}
