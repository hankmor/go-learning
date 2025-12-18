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
    buf := bufPool.Get().([]byte) 
    
    // 用完它...
    
    // 2. Put(): 还回去，下次给别人用
    // 注意：还之前最好重置一下状态（比如清空）
    bufPool.Put(buf)

    // 再次 Get，就不会触发 New，而是直接复用刚才那个
    buf2 := bufPool.Get().([]byte)
    _ = buf2
}
