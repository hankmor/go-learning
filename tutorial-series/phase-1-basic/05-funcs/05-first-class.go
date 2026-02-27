package main

import "fmt"

// Operation 函数类型定义
type Operation func(int, int) int

// calculate 接收一个操作函数作为参数
func calculate(a, b int, op Operation) int {
	return op(a, b)
}

// applyTwice 将函数应用两次
func applyTwice(x int, fn func(int) int) int {
	return fn(fn(x))
}

// compose 函数组合：返回 f(g(x))
func compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

// map 对切片中的每个元素应用函数
func mapInt(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = fn(num)
	}
	return result
}

// reduce 将切片归约为单个值
func reduce(nums []int, initial int, fn func(int, int) int) int {
	result := initial
	for _, num := range nums {
		result = fn(result, num)
	}
	return result
}

// getOperation 根据操作符返回对应的函数
func getOperation(op string) Operation {
	switch op {
	case "+":
		return func(a, b int) int { return a + b }
	case "-":
		return func(a, b int) int { return a - b }
	case "*":
		return func(a, b int) int { return a * b }
	case "/":
		return func(a, b int) int { return a / b }
	default:
		return nil
	}
}

func main() {
	// 1. 定义函数变量
	add := func(x, y int) int { return x + y }
	multiply := func(x, y int) int { return x * y }

	fmt.Println("calculate(10, 5, add) =", calculate(10, 5, add))           // 15
	fmt.Println("calculate(10, 5, multiply) =", calculate(10, 5, multiply)) // 50

	// 2. 使用 applyTwice
	double := func(x int) int { return x * 2 }
	fmt.Println("applyTwice(3, double) =", applyTwice(3, double)) // 12 (3*2*2)

	// 3. 函数组合
	addOne := func(x int) int { return x + 1 }
	square := func(x int) int { return x * x }

	// 先加1，再平方
	addThenSquare := compose(square, addOne)
	fmt.Println("addThenSquare(3) =", addThenSquare(3)) // 16 ((3+1)^2)

	// 先平方，再加1
	squareThenAdd := compose(addOne, square)
	fmt.Println("squareThenAdd(3) =", squareThenAdd(3)) // 10 (3^2+1)

	// 4. map 函数
	numbers := []int{1, 2, 3, 4, 5}
	doubled := mapInt(numbers, func(x int) int { return x * 2 })
	fmt.Println("doubled:", doubled) // [2 4 6 8 10]

	squared := mapInt(numbers, func(x int) int { return x * x })
	fmt.Println("squared:", squared) // [1 4 9 16 25]

	// 5. reduce 函数
	sum := reduce(numbers, 0, func(acc, x int) int { return acc + x })
	fmt.Println("sum:", sum) // 15

	product := reduce(numbers, 1, func(acc, x int) int { return acc * x })
	fmt.Println("product:", product) // 120

	// 6. 动态获取操作函数
	ops := []string{"+", "-", "*", "/"}
	for _, op := range ops {
		fn := getOperation(op)
		if fn != nil {
			result := fn(10, 2)
			fmt.Printf("10 %s 2 = %d\n", op, result)
		}
	}

	// 7. 函数切片
	operations := []Operation{
		func(a, b int) int { return a + b },
		func(a, b int) int { return a - b },
		func(a, b int) int { return a * b },
		func(a, b int) int { return a / b },
	}

	fmt.Println("\nApplying all operations to 10 and 2:")
	for i, op := range operations {
		fmt.Printf("Operation %d: %d\n", i+1, op(10, 2))
	}
}
