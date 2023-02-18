package main

import (
	"fmt"
	"time"
)

// 从 nil 的chan读写会阻塞，但是有时候却可以巧妙的使用它

// 想要 2 个 goroutine 依次输出 2 3 后退出
func main() {
	// nilChan()
	// bad()
	betterWithNilChan()
}

func nilChan() {
	var a chan int
	go func() {
		fmt.Println("write to a nil chan")
		a <- 1
		fmt.Println("write success")
	}()
	go func() {
		fmt.Println("read from a nil chan")
		<-a
		fmt.Println("read success")
	}()
	select {} // 阻塞程序，不退出
	/*
	 fatal error: all goroutines are asleep - deadlock!
	*/
}

func bad() {
	var c1, c2 = make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- 2
		close(c1) // 关闭 chan，导致 select 读取时不会阻塞
	}()

	go func() {
		time.Sleep(time.Second * 3)
		c2 <- 3
		close(c2)
	}()

	var a, b bool
	for {
		select {
		case i := <-c1: // 先输出 2，然后关闭。关闭了的chan读取不会阻塞，此时 i 为零值 0，所以一直输出 0
			a = true
			fmt.Println(i)
		case j := <-c2:
			b = true
			fmt.Println(j)
		}
		time.Sleep(time.Millisecond * 100)
		if a && b {
			break
		}
	}
	/*output:
	2
	0
	0
	...
	3
	*/
}

func betterWithNilChan() {
	var c1, c2 = make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- 2
		close(c1) // 关闭 chan，导致 select 读取时不会阻塞
	}()

	go func() {
		time.Sleep(time.Second * 3)
		c2 <- 3
		close(c2)
	}()

	for {
		select {
		case i, ok := <-c1:
			if !ok {
				c1 = nil // 置为 nil，此时 select 从 nil的chan读取会阻塞
			} else {
				fmt.Println(i)
			}
		case j, ok := <-c2:
			if !ok {
				c2 = nil
			} else {
				fmt.Println(j)
			}
		}
		time.Sleep(time.Millisecond * 100)
		if c1 == nil && c2 == nil {
			break
		}
	}
	/*output:
	2
	3
	*/
}
