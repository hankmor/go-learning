package main

import (
	"fmt"
	"time"
)

// select 语句中表达式的求值顺序

func getReadonlyChan() <-chan int {
	fmt.Println("get a readonly chan")
	var c = make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		c <- 1
	}()
	return c
}

func getWriteOnlyChan() chan<- int {
	fmt.Println("get a write only chan")
	var c = make(chan int)
	return c
}

func getNumber() int {
	fmt.Println("get a number")
	return 2
}

func getSlice() []int {
	fmt.Println("get a slice")
	var s = make([]int, 5)
	return s
}

func main() {
	// 1、先求值 chan
	// 2、getSlice 延迟求值，选择该 case 后再求值
	// 3、getNumber 向 chan 写入数据，先求值
	select {
	case getSlice()[0] = <-getReadonlyChan():
		fmt.Println("enter case 1")
	case getWriteOnlyChan() <- getNumber():
		fmt.Println("enter case 2")
	}

	/*output:
	get a readonly chan
	get a write only chan
	get a number
	get a slice
	enter case 1
	*/
}
