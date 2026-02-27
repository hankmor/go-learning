package main

import "fmt"

func main() {
	// 1. 遍历切片
	fmt.Println("=== Range over slice ===")
	nums := []int{1, 2, 3, 4, 5}
	for i, v := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// 2. 只要索引
	fmt.Println("\n=== Only index ===")
	for i := range nums {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 3. 只要值（使用 _ 忽略索引）
	fmt.Println("\n=== Only value ===")
	for _, v := range nums {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// 4. 遍历数组
	fmt.Println("\n=== Range over array ===")
	arr := [3]string{"a", "b", "c"}
	for i, v := range arr {
		fmt.Printf("%d: %s\n", i, v)
	}

	// 5. 遍历字符串（按 rune 遍历，不是 byte）
	fmt.Println("\n=== Range over string ===")
	str := "Hello, 世界"
	for i, ch := range str {
		fmt.Printf("Index: %d, Char: %c (Unicode: %U)\n", i, ch, ch)
	}

	// 6. 遍历 map
	fmt.Println("\n=== Range over map ===")
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// 7. 只要 key
	fmt.Println("\n=== Only keys from map ===")
	for key := range m {
		fmt.Print(key, " ")
	}
	fmt.Println()

	// 8. 遍历 channel
	fmt.Println("\n=== Range over channel ===")
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // 必须关闭，否则会死锁

	for v := range ch {
		fmt.Println("Received:", v)
	}

	// 9. 使用 range 复制切片
	fmt.Println("\n=== Copy slice using range ===")
	original := []int{1, 2, 3}
	copied := make([]int, len(original))
	for i, v := range original {
		copied[i] = v
	}
	fmt.Println("Original:", original)
	fmt.Println("Copied:", copied)
}
