package main

import (
	"fmt"
	"sync"
)

var cnt1 int
var sg sync.WaitGroup

func main() {
	sg.Add(2)
	ch := make(chan int)
	go countAdd1(ch) // 增加
	go func() {      // 读取
		for v := range ch {
			fmt.Println("received: ", v)
		}
		sg.Done() // 读取任务完成
	}()
	sg.Wait()
}

func countAdd1(ch chan<- int) {
	for i := 0; i < 1000; i++ {
		cnt1++
		ch <- cnt1 // 写回通道
	}
	close(ch) // 增加完成，关闭通道
	sg.Done() // 任务完成
}
