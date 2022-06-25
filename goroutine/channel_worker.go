package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 使用带缓冲的通道来处理多个工作

var (
	wsg            sync.WaitGroup
	goroutineNum   = 4
	workWaitingNum = 10
)

func MultiWork() {
	// 带缓冲的通道，可以等待执行的工作数量
	wc := make(chan string, workWaitingNum)
	wsg.Add(goroutineNum)

	for i := 0; i < goroutineNum; i++ {
		go Work(wc, i)
	}

	// 创建多个工作
	for i := 0; i < 20; i++ {
		wc <- fmt.Sprintf("task: %d", i)
	}

	// 所有工作完成，关闭通道
	close(wc)

	wsg.Wait()
}

func Work(wc chan string, worker int) {
	defer wsg.Done()

	for {
		task, ok := <-wc // 获取一个任务

		if !ok {
			println("Worker ", worker, " shutdown")
			return
		}

		fmt.Printf("Worker %d started: %s\n", worker, task)

		// 模拟工作时间
		time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)

		fmt.Printf("Worker %d completed: %s\n", worker, task)
	}
}

func main() {
	MultiWork()
}
