package main

import "fmt"

// Counter 计数器结构体
type Counter struct {
	count int
}

// GetCount 值接收者方法
func (c Counter) GetCount() int {
	return c.count
}

// Increment 指针接收者方法
func (c *Counter) Increment() {
	c.count++
}

// IncrementBy 指针接收者方法，带参数
func (c *Counter) IncrementBy(n int) {
	c.count += n
}

// Reset 指针接收者方法
func (c *Counter) Reset() {
	c.count = 0
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 计算面积（值接收者）
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算周长（值接收者）
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Scale 缩放（指针接收者）
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Calculator 计算器结构体
type Calculator struct {
	value int
}

// Add 加法
func (c *Calculator) Add(n int) *Calculator {
	c.value += n
	return c // 返回自身，支持链式调用
}

// Subtract 减法
func (c *Calculator) Subtract(n int) *Calculator {
	c.value -= n
	return c
}

// Multiply 乘法
func (c *Calculator) Multiply(n int) *Calculator {
	c.value *= n
	return c
}

// Result 获取结果
func (c *Calculator) Result() int {
	return c.value
}

// methodValueDemo 演示方法值
func methodValueDemo() {
	fmt.Println("\n=== Method Value Demo ===")
	calc := &Calculator{value: 10}

	// 方法值：绑定到特定实例
	addMethod := calc.Add
	addMethod(5)
	fmt.Println("After Add(5):", calc.value) // 15

	// 可以将方法值赋值给变量
	add := calc.Add
	subtract := calc.Subtract

	add(10)
	fmt.Println("After Add(10):", calc.value) // 25

	subtract(5)
	fmt.Println("After Subtract(5):", calc.value) // 20
}

// methodExpressionDemo 演示方法表达式
func methodExpressionDemo() {
	fmt.Println("\n=== Method Expression Demo ===")
	calc := &Calculator{value: 10}

	// 方法表达式：需要显式传递接收者
	addExpr := (*Calculator).Add
	addExpr(calc, 5)
	fmt.Println("After Add(5):", calc.value) // 15

	// 方法表达式可以用于任何该类型的实例
	calc2 := &Calculator{value: 20}
	addExpr(calc2, 10)
	fmt.Println("calc2 value:", calc2.value) // 30
}

func main() {
	// 1. 基本方法调用
	fmt.Println("=== Basic Method Calls ===")
	counter := Counter{count: 0}
	fmt.Println("Initial count:", counter.GetCount())

	counter.Increment()
	fmt.Println("After Increment:", counter.GetCount())

	counter.IncrementBy(5)
	fmt.Println("After IncrementBy(5):", counter.GetCount())

	// 2. 值接收者 vs 指针接收者
	fmt.Println("\n=== Value vs Pointer Receiver ===")
	c1 := Counter{count: 0}
	c1.Increment() // Go 自动转换为 (&c1).Increment()
	fmt.Println("c1 count:", c1.GetCount())

	c2 := &Counter{count: 0}
	c2.Increment()
	fmt.Println("c2 count:", c2.GetCount())

	// 3. 矩形示例
	fmt.Println("\n=== Rectangle Demo ===")
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	rect.Scale(2)
	fmt.Printf("After Scale(2) - Width: %.2f, Height: %.2f\n", rect.Width, rect.Height)
	fmt.Printf("New Area: %.2f\n", rect.Area())

	// 4. 链式调用
	fmt.Println("\n=== Method Chaining ===")
	calc := &Calculator{value: 10}
	result := calc.Add(5).Multiply(2).Subtract(10).Result()
	fmt.Println("Result:", result) // (10+5)*2-10 = 20

	// 5. 方法值
	methodValueDemo()

	// 6. 方法表达式
	methodExpressionDemo()
}
