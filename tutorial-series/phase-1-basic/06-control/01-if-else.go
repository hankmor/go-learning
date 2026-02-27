package main

import "fmt"

func main() {
	// 1. 基本 if
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	// 2. if-else
	if x > 15 {
		fmt.Println("x is large")
	} else {
		fmt.Println("x is not that large")
	}

	// 3. if-else if-else 链
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	// 4. if 带初始化语句
	if y := x * 2; y > 15 {
		fmt.Println("y is", y) // y 的作用域仅限于 if-else 块
	} else {
		fmt.Println("y is not greater than 15, y =", y)
	}
	// fmt.Println(y) // 错误：y 在这里不可用

	// 5. 常用于错误处理
	if err := doSomething(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success")
	}

	// 6. 常用于 map 查找
	m := map[string]int{"a": 1, "b": 2}
	if value, ok := m["a"]; ok {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found")
	}

	// 7. 常用于类型断言
	var i interface{} = "hello"
	if str, ok := i.(string); ok {
		fmt.Println("String:", str)
	} else {
		fmt.Println("Not a string")
	}
}

// doSomething 模拟一个可能返回错误的函数
func doSomething() error {
	// 模拟成功
	return nil
}
