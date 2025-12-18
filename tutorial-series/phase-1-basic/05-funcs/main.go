package main

import "fmt"

// add 是一个普通函数，接收两个 int 参数，返回一个 int 结果
func add(a, b int) int {
    return a + b
}

// swap 演示多返回值，交换输入的两个字符串
func swap(x, y string) (string, string) {
    return y, x
}

// 定义一个简单的结构体 User
type User struct {
    Name string
    Age  int
}

// SayHello 是绑定到 User 结构体的方法
// (u User) 称为接收者（Receiver）
func (u User) SayHello() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", u.Name, u.Age)
}

// 只有指针接收者才能修改结构体内部的值
func (u *User) Grow() {
    u.Age++
}

func main() {
    // 1. 调用普通函数
    sum := add(10, 20)
    fmt.Println("10 + 20 =", sum)

    // 2. 调用多返回值函数
    a, b := swap("hello", "world")
    fmt.Println(a, b) // world hello

    // 3. 调用方法
    user := User{Name: "Hank", Age: 18}
    user.SayHello()
    
    // 调用指针接收者方法修改状态
    user.Grow()
    fmt.Printf("Age after grow: %d\n", user.Age) // 19
}
