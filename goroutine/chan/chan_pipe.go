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
		// 从 numbers 中读取数据，然后平方后再写入 squared
		for v := range numbers {
			squared <- v * v
		}
		close(squared) // 全部发送成功后关闭
	}()

	// 从 squared 通道读取值
	for v := range squared {
		fmt.Println(v)
	}
}
