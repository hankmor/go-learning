package main

import "fmt"

func main() {
	t1()
	t2()
	fmt.Println()
	fmt.Println(t3()) // 2
}

func t1() {
	i := 0
	defer fmt.Println(i) // 0
	i++
	return
}

func t2() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i) // 3210
	}
}

func t3() (i int) {
	defer func() { i++ }() // 先执行 i++，再返回
	return 1               // 返回2
}

func sumWithDefer() {
	defer func() {
		sum(100)
	}()
}
