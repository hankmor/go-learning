package main

import "fmt"

// 定义一个接口：只要实现了 Speak 方法的类型，都满足这个接口
type Speaker interface {
    Speak() string
}

// 定义结构体 Dog
type Dog struct {
    Name string
}

// Dog 实现 Speaker 接口
// 注意：Go 中没有 implements 关键字，这是隐式实现的
func (d Dog) Speak() string {
    return "Woof!"
}

// 定义结构体 Cat
type Cat struct {
    Name string
}

// Cat 实现 Speaker 接口
func (c Cat) Speak() string {
    return "Meow!"
}

// 多态演示：接收任何 Speaker
func introduce(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    d := Dog{Name: "Buddy"}
    c := Cat{Name: "Kitty"}

    introduce(d) // Woof!
    introduce(c) // Meow!

    // 接口类型断言
    var s Speaker = d
    if dog, ok := s.(Dog); ok {
        fmt.Printf("It's a dog named %s\n", dog.Name)
    }
}
