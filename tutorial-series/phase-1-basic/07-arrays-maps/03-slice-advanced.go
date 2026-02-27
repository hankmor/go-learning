package main

import "fmt"

func main() {
	// 1. 切片扩容机制
	fmt.Println("=== Slice Growth Mechanism ===")

	s := make([]int, 0, 1)
	fmt.Printf("Initial: len=%d cap=%d\n", len(s), cap(s))

	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("After append %d: len=%d cap=%d\n", i, len(s), cap(s))
	}
	// 观察容量变化：1 -> 2 -> 4 -> 8 -> 16

	// 2. 删除元素
	fmt.Println("\n=== Delete Elements ===")

	s = []int{1, 2, 3, 4, 5}
	fmt.Println("Original:", s)

	// 删除索引 2 的元素（保持顺序）
	index := 2
	s = append(s[:index], s[index+1:]...)
	fmt.Println("After delete index 2:", s) // [1, 2, 4, 5]

	// 删除索引 1 的元素（不保持顺序，更高效）
	s = []int{1, 2, 3, 4, 5}
	index = 1
	s[index] = s[len(s)-1] // 用最后一个元素替换
	s = s[:len(s)-1]       // 截断最后一个元素
	fmt.Println("After fast delete index 1:", s) // [1, 5, 3, 4]

	// 3. 插入元素
	fmt.Println("\n=== Insert Elements ===")

	s = []int{1, 2, 4, 5}
	fmt.Println("Original:", s)

	// 在索引 2 插入 3
	index = 2
	value := 3
	s = append(s[:index], append([]int{value}, s[index:]...)...)
	fmt.Println("After insert 3 at index 2:", s) // [1, 2, 3, 4, 5]

	// 更高效的插入方式
	s = []int{1, 2, 4, 5}
	index = 2
	value = 3
	s = append(s, 0)       // 扩展切片
	copy(s[index+1:], s[index:]) // 移动元素
	s[index] = value       // 插入新元素
	fmt.Println("After efficient insert:", s)

	// 4. 过滤元素
	fmt.Println("\n=== Filter Elements ===")

	s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original:", s)

	// 过滤出偶数
	evens := make([]int, 0, len(s)/2)
	for _, v := range s {
		if v%2 == 0 {
			evens = append(evens, v)
		}
	}
	fmt.Println("Evens:", evens)

	// 原地过滤（不分配新内存）
	s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := 0
	for _, v := range s {
		if v%2 == 0 {
			s[n] = v
			n++
		}
	}
	s = s[:n]
	fmt.Println("In-place filtered:", s)

	// 5. 反转切片
	fmt.Println("\n=== Reverse Slice ===")

	s = []int{1, 2, 3, 4, 5}
	fmt.Println("Original:", s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println("Reversed:", s)

	// 6. 去重
	fmt.Println("\n=== Remove Duplicates ===")

	s = []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	fmt.Println("Original:", s)

	// 使用 map 去重
	seen := make(map[int]bool)
	unique := make([]int, 0, len(s))
	for _, v := range s {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}
	fmt.Println("Unique:", unique)

	// 7. 合并切片
	fmt.Println("\n=== Merge Slices ===")

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	// 方式 1：使用 append
	merged := append(s1, s2...)
	fmt.Println("Merged with append:", merged)

	// 方式 2：使用 copy（更高效）
	s1 = []int{1, 2, 3}
	s2 = []int{4, 5, 6}
	merged = make([]int, len(s1)+len(s2))
	copy(merged, s1)
	copy(merged[len(s1):], s2)
	fmt.Println("Merged with copy:", merged)

	// 8. 切片作为栈
	fmt.Println("\n=== Slice as Stack ===")

	var stack []int

	// Push
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	fmt.Println("Stack after pushes:", stack)

	// Pop
	if len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("Popped: %d, Stack: %v\n", top, stack)
	}

	// Peek
	if len(stack) > 0 {
		top := stack[len(stack)-1]
		fmt.Printf("Top: %d, Stack: %v\n", top, stack)
	}

	// 9. 切片作为队列
	fmt.Println("\n=== Slice as Queue ===")

	var queue []int

	// Enqueue
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	fmt.Println("Queue after enqueues:", queue)

	// Dequeue
	if len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		fmt.Printf("Dequeued: %d, Queue: %v\n", front, queue)
	}

	// 10. 二维切片
	fmt.Println("\n=== 2D Slice ===")

	// 创建二维切片
	rows, cols := 3, 4
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	// 填充数据
	count := 1
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = count
			count++
		}
	}

	// 打印矩阵
	fmt.Println("Matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
