package main

import "fmt"

// sum 可变参数函数
// 使用 ... 语法可以接收任意数量的参数
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// printf 风格的函数，第一个参数是格式字符串，后面是可变参数
func printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// concat 连接字符串，可以指定分隔符
func concat(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i]
	}
	return result
}

// max 返回最大值，至少需要一个参数
func max(first int, rest ...int) int {
	maxVal := first
	for _, num := range rest {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}

func main() {
	// 1. 基本用法
	fmt.Println("sum(1, 2, 3) =", sum(1, 2, 3))
	fmt.Println("sum(1, 2, 3, 4, 5) =", sum(1, 2, 3, 4, 5))

	// 2. 传递切片，使用 ... 展开
	numbers := []int{10, 20, 30, 40}
	fmt.Println("sum(numbers...) =", sum(numbers...))

	// 3. 可变参数可以为空
	fmt.Println("sum() =", sum())

	// 4. 自定义 printf
	printf("Name: %s, Age: %d\n", "Hank", 18)

	// 5. 字符串连接
	result := concat(", ", "apple", "banana", "orange")
	fmt.Println("Fruits:", result)

	// 6. 至少一个参数的可变参数
	fmt.Println("max(5, 2, 8, 1, 9) =", max(5, 2, 8, 1, 9))
	fmt.Println("max(42) =", max(42))
}
