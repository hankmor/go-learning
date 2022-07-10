package main

import (
	"fmt"
	"time"
)

func main() {
	// 申明chan，类型为int
	var c chan int
	// 创建chan，无缓冲
	c = make(chan int)
	// 主程序未退出，则一直展示旋转动画
	go spinner(100 * time.Millisecond)
	// 开启 goroutine 执行匿名函数
	go func() {
		// 休眠2s，然后从chan读取数据
		time.Sleep(2 * time.Second)
		// \b退格符删除spinner的字符，然后打印读取
		fmt.Printf("\breading from chan...\n")
		i := <-c // 读取数据，读不到阻塞
		fmt.Println(i)
		close(c) // 关闭通道
	}()

	// 写入一条数据，成功读取后程序退出
	fmt.Println("writing to chan...")
	c <- 1 // 无缓冲通道，如果没有被读取，则阻塞

	fmt.Println("exit")
}

// 自旋，防止主程序退出，一个转动的动画效果
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
