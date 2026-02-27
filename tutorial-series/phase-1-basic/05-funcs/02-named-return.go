package main

import (
	"errors"
	"fmt"
)

// divide 命名返回值示例
// 命名返回值可以在函数体内直接使用，并且可以使用裸 return 语句
func divide(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("division by zero")
		return // 等价于 return result, err
	}
	result = a / b
	return // 自动返回 result 和 err
}

// divideVerbose 显式返回，更清晰但稍微冗长
func divideVerbose(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// calculateStats 命名返回值在复杂计算中很有用
func calculateStats(numbers []int) (sum, avg, min, max int) {
	if len(numbers) == 0 {
		return // 所有返回值都是零值
	}

	sum = numbers[0]
	min = numbers[0]
	max = numbers[0]

	for _, num := range numbers {
		sum += num
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	avg = sum / len(numbers)
	return // 自动返回所有命名的返回值
}

func main() {
	// 测试 divide 函数
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// 测试除以零的情况
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 测试统计函数
	numbers := []int{5, 2, 8, 1, 9, 3}
	sum, avg, min, max := calculateStats(numbers)
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Sum: %d, Avg: %d, Min: %d, Max: %d\n", sum, avg, min, max)
}
