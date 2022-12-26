package main

import (
	"fmt"
	"time"
)

func Count(ch chan<- int) {
	ch <- 1
	time.Sleep(100 * time.Millisecond)
}

func main() {
	// 创建一个 chan 数组
	chs := make([]chan int, 100)
	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan int) // 创建 channel
		go Count(chs[i])        // 并发地向 channel 写入数据
	}

	// 汇总读取到的数据
	var n int
	for _, ch := range chs {
		i := <-ch
		n = n + i
	}
	fmt.Println("received: ", n)
}
