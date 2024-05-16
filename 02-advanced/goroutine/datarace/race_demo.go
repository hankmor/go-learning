package main

import "fmt"

// data race detector
func main() {
	c := make(chan int)
	m := make(map[string]int)
	go func() {
		m["a"] = 1 // 访问map冲突
		c <- 1
	}()
	m["a"] = 2 // 访问map冲突
	<-c
	for k, v := range m {
		fmt.Printf("key = %v, val = %v\n", k, v)
	}
}
