package main

import "fmt"

func main() {
	// 1. 标准 for 循环
	fmt.Println("=== Standard for loop ===")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 2. 类似 while 的形式
	fmt.Println("\n=== While-like loop ===")
	count := 0
	for count < 5 {
		fmt.Print(count, " ")
		count++
	}
	fmt.Println()

	// 3. 无限循环
	fmt.Println("\n=== Infinite loop with break ===")
	i := 0
	for {
		if i >= 5 {
			break
		}
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	// 4. continue 跳过本次迭代
	fmt.Println("\n=== Continue demo ===")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // 跳过偶数
		}
		fmt.Print(i, " ") // 只打印奇数
	}
	fmt.Println()

	// 5. break 跳出循环
	fmt.Println("\n=== Break demo ===")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 跳出循环
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 6. 嵌套循环
	fmt.Println("\n=== Nested loops ===")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}

	// 7. 省略初始化和后置语句
	fmt.Println("\n=== Omit init and post ===")
	sum := 0
	for ; sum < 10; {
		sum += 2
		fmt.Print(sum, " ")
	}
	fmt.Println()
}
