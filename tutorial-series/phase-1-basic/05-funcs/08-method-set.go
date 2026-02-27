package main

import "fmt"

// Counter 计数器
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

// Incrementer 接口：只有 Increment 方法
type Incrementer interface {
	Increment()
}

// Getter 接口：只有 GetCount 方法
type Getter interface {
	GetCount() int
}

// CounterInterface 接口：包含两个方法
type CounterInterface interface {
	GetCount() int
	Increment()
}

// doIncrement 接收 Incrementer 接口
func doIncrement(i Incrementer) {
	i.Increment()
}

// doGet 接收 Getter 接口
func doGet(g Getter) int {
	return g.GetCount()
}

// doCounter 接收 CounterInterface 接口
func doCounter(c CounterInterface) {
	c.Increment()
	fmt.Println("Count:", c.GetCount())
}

// methodSetDemo 演示方法集规则
func methodSetDemo() {
	fmt.Println("=== Method Set Demo ===")

	// 值类型
	c1 := Counter{count: 0}
	fmt.Println("c1 initial:", c1.GetCount()) // ✓ 可以调用值接收者方法
	c1.Increment()                            // ✓ Go 自动转换为 (&c1).Increment()
	fmt.Println("c1 after increment:", c1.GetCount())

	// 指针类型
	c2 := &Counter{count: 0}
	fmt.Println("c2 initial:", c2.GetCount()) // ✓ 可以调用值接收者方法
	c2.Increment()                            // ✓ 可以调用指针接收者方法
	fmt.Println("c2 after increment:", c2.GetCount())
}

// interfaceDemo 演示接口实现的差异
func interfaceDemo() {
	fmt.Println("\n=== Interface Implementation Demo ===")

	c1 := Counter{count: 0}
	c2 := &Counter{count: 0}

	// 值类型实现了 Getter 接口（只有值接收者方法）
	fmt.Println("c1 implements Getter:", doGet(c1)) // ✓

	// 指针类型也实现了 Getter 接口
	fmt.Println("c2 implements Getter:", doGet(c2)) // ✓

	// 值类型不能作为 Incrementer（需要指针接收者方法）
	// doIncrement(c1) // ✗ 编译错误：Counter does not implement Incrementer

	// 指针类型实现了 Incrementer
	doIncrement(c2) // ✓
	fmt.Println("c2 after doIncrement:", c2.GetCount())

	// 值类型不能作为 CounterInterface
	// doCounter(c1) // ✗ 编译错误

	// 指针类型实现了 CounterInterface
	doCounter(c2) // ✓
}

// Person 演示嵌入类型的方法提升
type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() {
	fmt.Printf("Hi, I'm %s, %d years old.\n", p.Name, p.Age)
}

func (p *Person) Birthday() {
	p.Age++
	fmt.Printf("%s is now %d years old.\n", p.Name, p.Age)
}

// Employee 嵌入 Person
type Employee struct {
	Person // 嵌入
	Title  string
}

// Work 是 Employee 自己的方法
func (e Employee) Work() {
	fmt.Printf("%s is working as a %s.\n", e.Name, e.Title)
}

// embeddingDemo 演示嵌入类型的方法提升
func embeddingDemo() {
	fmt.Println("\n=== Embedding Demo ===")

	emp := Employee{
		Person: Person{Name: "Hank", Age: 18},
		Title:  "Engineer",
	}

	// 可以直接调用嵌入类型的方法
	emp.Introduce() // Person 的方法
	emp.Birthday()  // Person 的方法
	emp.Work()      // Employee 自己的方法

	fmt.Println("Final age:", emp.Age)
}

// 演示方法集的完整规则
func methodSetRules() {
	fmt.Println("\n=== Method Set Rules ===")
	fmt.Println("Type T:")
	fmt.Println("  - Can call methods with receiver T")
	fmt.Println("  - Can call methods with receiver *T (Go auto-converts)")
	fmt.Println("  - Method set for interface: only methods with receiver T")
	fmt.Println()
	fmt.Println("Type *T:")
	fmt.Println("  - Can call methods with receiver T")
	fmt.Println("  - Can call methods with receiver *T")
	fmt.Println("  - Method set for interface: methods with receiver T and *T")
}

func main() {
	// 1. 方法集基本演示
	methodSetDemo()

	// 2. 接口实现的差异
	interfaceDemo()

	// 3. 嵌入类型的方法提升
	embeddingDemo()

	// 4. 方法集规则总结
	methodSetRules()

	// 5. 实际建议
	fmt.Println("\n=== Best Practices ===")
	fmt.Println("1. 优先使用指针接收者，除非：")
	fmt.Println("   - 类型很小（如 time.Time）")
	fmt.Println("   - 类型是不可变的")
	fmt.Println("   - 类型是 map、slice、channel 等引用类型")
	fmt.Println()
	fmt.Println("2. 保持一致性：同一类型的方法要么都用值接收者，要么都用指针接收者")
	fmt.Println()
	fmt.Println("3. 如果类型需要实现接口，考虑使用指针接收者")
}
