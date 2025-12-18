package main

import (
	"errors"
	"fmt"
)

// 定义一个除法函数，返回 result 和 error
func divide(a, b int) (int, error) {
    if b == 0 {
        // 使用 errors.New 创建一个简单的错误对象
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

// 演示 panic 和 recover
func safeCall() {
    // defer 必须在 panic 发生前定义
    defer func() {
        // recover() 捕获 panic，如果返回值不为 nil，说明发生了 panic
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    panic("Something went wrong terribly!")
    fmt.Println("This line will not execute")
}

func main() {
    // 1. 标准错误处理
    res, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", res)
    }

    // 2. 演示从 panic 中恢复
    fmt.Println("Starting safeCall...")
    safeCall()
    fmt.Println("Program continues...")
}
