package main

import (
	"fmt"
	"time"
)

func Worker(j int) {
	fmt.Println("invoking worker")
	time.Sleep(time.Second * time.Duration(j))
}

func spawn(f func(int)) chan string {
	quit := make(chan string)
	go func() {
		var job = make(chan int)
		go func() {
			time.Sleep(time.Second * time.Duration(2))
			fmt.Println("write to job chan")
			job <- 2
		}()
		for {
			fmt.Println("enter for")
			select {
			case j := <-job:
				f(j)
			case <-quit:
				quit <- "ok"
			}
		}
	}()
	return quit
}

func main() {
	quit := spawn(Worker)
	println("spawn a worker goroutine")

	time.Sleep(5 * time.Second)

	println("notify the worker to exit")
	quit <- "exit"

	timer := time.NewTimer(time.Second * 10)
	defer timer.Stop()
	select {
	case status := <-quit:
		println("worker done:", status)
	case <-timer.C:
		println("wait worker exit timeout")
	}
	time.Sleep(time.Second * time.Duration(20))
	fmt.Println("main exit")
}
