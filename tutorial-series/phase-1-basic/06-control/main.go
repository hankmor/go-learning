package main

import "fmt"

func main() {
    // 1. if 语句
    x := 10
    // 条件判断不需要小括号 ()
    if x > 5 {
        fmt.Println("x is large")
    }

    // if 支持初始化语句：先执行初始化，再判断
    if y := x * 2; y > 15 {
        fmt.Println("y is", y) // y 的作用域仅限于 if 块
    }

    // 2. switch 语句
    day := "Mon"
    switch day {
    case "Mon":
        fmt.Println("Start of week")
        // Go 默认不需要 break，自动终止
    case "Fri":
        fmt.Println("Weekend is coming")
    default:
        fmt.Println("Other day")
    }

    // 3. for 循环：Go 唯一的循环结构
    // 形式一：类似于 C/Java 的 for
    for i := 0; i < 3; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println()

    // 形式二：类似于 while
    count := 3
    for count > 0 {
        fmt.Print(count, " ")
        count--
    }
    fmt.Println()

    // 4. defer 延迟执行
    // 常用于资源释放，函数返回前才会执行
    defer fmt.Println("Exiting main function...") 
    fmt.Println("Doing some work...")
}
