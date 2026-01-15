package main

import (
	"fmt"
	"sync"
)

// 定义池子
var bufPool = sync.Pool{
	// New 函数：当池子里没存货时，调用它创建一个新的
	New: func() interface{} {
		fmt.Println("Creating new buffer")
		return make([]byte, 1024)
	},
}

func main() {
	// 1.Get(): 借一个对象
	// Get 返回的是 interface{}，需要断言成具体类型
	buf := bufPool.Get().([]byte)
	fmt.Printf("Got buffer of len: %d\n", len(buf)) // Creating new buffer

	// 用完它...模拟使用
	buf[0] = 1

	// 2. Put(): 还回去，下次给别人用
	// 注意：还之前最好重置一下状态（比如清空）
	// 这里简单演示，实际可能需要 buf = buf[:0]
	bufPool.Put(buf)

	// 再次 Get，就不会触发 New，而是直接复用刚才那个
	// 注意：sync.Pool 的复用策略由 runtime 决定，不保证一定复用，
	// 也不保证获取到的是刚才 Put 进去的那个（可能有多个 P）
	buf2 := bufPool.Get().([]byte)
	fmt.Printf("Got buffer again, len: %d\n", len(buf2))
	_ = buf2
}
