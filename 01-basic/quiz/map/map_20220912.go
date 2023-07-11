package main

import "fmt"

func main() {
	m := make(map[int]int, 3) // 虽然指定了初始长度，但是 map 中没有任何元素，其 len 为0
	x := len(m)               // 刚创建map，没有任何元素，len 为 0
	m[1] = m[1]               // 给map的第2个元素赋值，现在有一个元素了，map 的 len 变为 1
	y := len(m)               // 1
	fmt.Println(x, y)         // 0 1
	fmt.Printf("%v\n", m)
}
