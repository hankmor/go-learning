package main

import (
	"fmt"
	"sync"
)

var cnt1 int

func main() {
	// ch := make(chan int)
	// for i := 0; i < 10; i++ {
	// 	go countAdd1(ch) // 加1
	// }
	// time.Sleep(1 * time.Second)
	// for v := range ch {
	// 	fmt.Println(v)
	// }
	// close(ch)
	// fatal error: all goroutines are asleep - deadlock!

	var sg sync.WaitGroup
	sg.Add(2)
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go countAdd1(ch) // 加1
		if i == 9 {
			sg.Done()
		}
	}
	// time.Sleep(1 * time.Second)
	for v := range ch {
		fmt.Println("received: ", v)
		if v == 10 {
			close(ch)
			sg.Done()
		}
	}
	sg.Wait()
}

func countAdd1(ch chan<- int) {
	cnt1++
	ch <- cnt1 // 写回通道
}
