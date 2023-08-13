package main

import (
	"context"
	"fmt"
)

// 当前目录下执行 wire 可以生成 wire_gen.go，执行 go generate 可以重新生成该文件

func main() {
	s, f, err := initializeBaz(context.Background())
	f()
	fmt.Println(s, err)
}
