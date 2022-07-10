package main

import (
	"fmt"
	"time"
)

var cnt1 int

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go countAdd1(ch) // 加1
	}
	time.Sleep(5 * time.Second)
	for v := range ch {
		fmt.Println(v)
	}
}

func countAdd1(ch chan int) {
	cnt1++
	ch <- cnt1 // 写回通道
}
