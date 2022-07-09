package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func main() {
	wg1.Add(1)

	ch := make(chan int)
	timer := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		ch <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		timer <- 1
	}()

L:
	// for {
	// 没有default，必须等到ch 或者 timer能够读取到数据
	select {
	case <-ch:
		fmt.Println("ch")
		wg1.Done()
		break L
	case <-timer:
		fmt.Println("timer")
		wg1.Done()
		break L
		// default:
		// time.Sleep(1 * time.Second)
		// 	fmt.Println("default")
	}
	time.Sleep(1 * time.Second)
	// }

	wg1.Wait()
}
