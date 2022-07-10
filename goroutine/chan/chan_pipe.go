package main

import "fmt"

func main() {
	// 无缓冲通道，产生整数
	numbers := make(chan int)
	// 无缓冲通道，整数平方
	squared := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			numbers <- i
		}
		close(numbers) // 全部发送成功后关闭
	}()

	go func() {
		for v := range numbers {
			squared <- v * v
		}
		close(squared) // 全部发送成功后关闭
	}()

	for v := range squared { // 循环读取通道的值
		fmt.Println(v)
	}
}
