package main

import "fmt"

// [T int | float64] 定义了类型参数 T
// T 被限制为 int 或 float64（类型约束）
func Min[T int | float64](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// Stack 是一个泛型结构体，T 可以是任何类型 (any)
type Stack[T any] struct {
    elements []T
}

func (s *Stack[T]) Push(v T) {
    s.elements = append(s.elements, v)
}

func (s *Stack[T]) Pop() T {
    if len(s.elements) == 0 {
        var zero T 
        return zero // 返回 T 类型的零值
    }
    v := s.elements[len(s.elements)-1]
    s.elements = s.elements[:len(s.elements)-1]
    return v
}

func main() {
    // 隐式类型推导（推荐）：编译器自动根据参数推导出 T 是 float64
    fmt.Println(Min(3.14, 1.59))

    // 实例化一个 int 类型的栈
    sInt := Stack[int]{}
    sInt.Push(10)
    sInt.Push(20)
    fmt.Println(sInt.Pop()) // 20

    // 实例化一个 string 类型的栈
    sStr := Stack[string]{}
    sStr.Push("hello")
    fmt.Println(sStr.Pop()) // hello
}
