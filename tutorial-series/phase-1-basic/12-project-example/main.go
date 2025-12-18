package main

import (
	"fmt"
	"os"
	"strconv"

	// 导入刚才定义的包
	"example.com/calc/internal/mathop"
)

func main() {
    // 读取命令行参数
    if len(os.Args) < 3 {
        fmt.Println("Usage: calc <num1> <num2>")
        return
    }

    // 解析参数
    a, _ := strconv.Atoi(os.Args[1])
    b, _ := strconv.Atoi(os.Args[2])

    // 调用业务逻辑
    sum := mathop.Add(a, b)
    fmt.Printf("%d + %d = %d\n", a, b, sum)
}
