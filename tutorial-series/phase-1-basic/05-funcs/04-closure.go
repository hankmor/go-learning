package main

import "fmt"

// makeAdder 返回一个闭包函数
// 闭包会捕获外部变量 x
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// makeMultiplier 返回一个乘法闭包
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// makeCounter 返回一个计数器闭包
// 每次调用都会递增内部计数器
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// fibonacci 返回一个斐波那契数列生成器
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

// filter 高阶函数：接收一个过滤函数，返回满足条件的元素
func filter(nums []int, fn func(int) bool) []int {
	var result []int
	for _, num := range nums {
		if fn(num) {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	// 1. 匿名函数直接调用
	func(msg string) {
		fmt.Println(msg)
	}("Hello, Anonymous!")

	// 2. 将匿名函数赋值给变量
	greet := func(name string) string {
		return "Hello, " + name
	}
	fmt.Println(greet("Hank"))

	// 3. 闭包：捕获外部变量
	counter := 0
	increment := func() int {
		counter++ // 捕获并修改外部变量
		return counter
	}

	fmt.Println("Counter:", increment()) // 1
	fmt.Println("Counter:", increment()) // 2
	fmt.Println("Counter:", increment()) // 3

	// 4. 使用 makeAdder 创建不同的加法器
	add5 := makeAdder(5)
	add10 := makeAdder(10)

	fmt.Println("add5(3) =", add5(3))   // 8
	fmt.Println("add10(3) =", add10(3)) // 13

	// 5. 使用 makeMultiplier
	double := makeMultiplier(2)
	triple := makeMultiplier(3)

	fmt.Println("double(5) =", double(5)) // 10
	fmt.Println("triple(5) =", triple(5)) // 15

	// 6. 计数器闭包
	counter1 := makeCounter()
	counter2 := makeCounter()

	fmt.Println("counter1:", counter1()) // 1
	fmt.Println("counter1:", counter1()) // 2
	fmt.Println("counter2:", counter2()) // 1
	fmt.Println("counter1:", counter1()) // 3

	// 7. 斐波那契数列生成器
	fib := fibonacci()
	fmt.Print("Fibonacci: ")
	for i := 0; i < 10; i++ {
		fmt.Print(fib(), " ")
	}
	fmt.Println()

	// 8. 高阶函数：filter
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 过滤出偶数
	evens := filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Evens:", evens)

	// 过滤出大于 5 的数
	greaterThan5 := filter(numbers, func(n int) bool {
		return n > 5
	})
	fmt.Println("Greater than 5:", greaterThan5)
}
