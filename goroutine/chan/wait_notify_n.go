package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int) {
	fmt.Printf("working for %d miliseconds\n", i*100)
	time.Sleep(time.Millisecond * time.Duration(i*100))
}

func spawn2(n int, f func(i int)) chan struct{} {
	job := make(chan int)
	quit := make(chan struct{})
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			var name = fmt.Sprintf("worker-%d", i)
			go func(j int) {
				job <- j
			}(i)
			defer wg.Done()
			for {
				j, ok := <-job
				if !ok { // 关闭说明所有worker完成，退出
					fmt.Println(name, "done")
					return
				}
				f(j) // 调用worker
			}
		}(i)
	}
	go func() {
		<-quit // 等待退出信号
		fmt.Println("all workers done")
		close(job) // 所有 worker 完成，关闭 job chan
		wg.Wait()
		quit <- struct{}{} // 继续向 quit 写入完成信号，表示所有 worker 完成
	}()
	return quit
}

func main() {
	q := spawn2(5, worker)
	println("spawn a group of worker")

	time.Sleep(5 * time.Second)
	println("notify the worker group to exit")
	q <- struct{}{} // 发送退出信号

	timer := time.NewTimer(time.Second * 5)
	defer timer.Stop()
	select {
	case <-timer.C:
		println("wait group worker exit timeout")
	case <-q:
		println("group workers done")
	}
	println("main exit")
}
