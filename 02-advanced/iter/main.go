package main

import (
	"fmt"
	"runtime"
	"slices"
)

// go version >= 1.23rc1

// 可迭代序列，形式为: func(func() bool)
type Seq0 func(yield func() bool)

// 定一个迭代器函数
func iter0[S ~[]E, E any](s S) Seq0 {
	return func(yield func() bool) {
		for range s {
			if !yield() { // yield函数如果返回 true 则还能继续迭代，否则迭代完成
				return
			}
		}
	}
}

func main() {
	fmt.Println(runtime.Version()) // go1.23rc1
	s1 := []int{1, 2, 3, 4, 5, 6, 7}

	// 自定义迭代器，返回一个没有人和元素的迭代器 Seq0
	var cnt int
	for range iter0(s1) {
		cnt++
	}
	fmt.Println("cnt: ", cnt) // 7

	// Values方法返回一个迭代器，从第一个元素开始
	for v := range slices.Values(s1) {
		fmt.Print(v)
	}
	// 1234567
	fmt.Println()

	// Backward返回slice的迭代器，包含 k，v 两个值，k为索引，v 为slice中的元素，顺序为倒叙，从最后一个元素向前迭代
	for k, v := range slices.Backward(s1) {
		fmt.Printf("%v - %v\n", k, v)
	}
	// 6 - 7
	// 5 - 6
	// 4 - 5
	// 3 - 4
	// 2 - 3
	// 1 - 2
	// 0 - 1
}
