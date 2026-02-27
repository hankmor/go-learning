package main

import "fmt"

func main() {
	// 1. 切片的创建方式
	fmt.Println("=== Slice Creation ===")

	// 从数组创建
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:4]
	fmt.Printf("slice1: %v, len: %d, cap: %d\n", slice1, len(slice1), cap(slice1))

	// 使用字面量
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice2: %v, len: %d, cap: %d\n", slice2, len(slice2), cap(slice2))

	// 使用 make
	slice3 := make([]int, 5)      // len=5, cap=5
	slice4 := make([]int, 3, 5)   // len=3, cap=5
	fmt.Printf("slice3: %v, len: %d, cap: %d\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("slice4: %v, len: %d, cap: %d\n", slice4, len(slice4), cap(slice4))

	// nil 切片 vs 空切片
	var nilSlice []int
	emptySlice := []int{}
	emptySlice2 := make([]int, 0)

	fmt.Printf("\nnilSlice: %v, len: %d, cap: %d, is nil: %t\n",
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("emptySlice: %v, len: %d, cap: %d, is nil: %t\n",
		emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
	fmt.Printf("emptySlice2: %v, len: %d, cap: %d, is nil: %t\n",
		emptySlice2, len(emptySlice2), cap(emptySlice2), emptySlice2 == nil)

	// 2. 切片的切片操作
	fmt.Println("\n=== Slice Slicing ===")

	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Original:", s)

	s1 := s[1:3]   // [2, 3]
	s2 := s[:3]    // [1, 2, 3]
	s3 := s[2:]    // [3, 4, 5]
	s4 := s[:]     // [1, 2, 3, 4, 5]

	fmt.Println("s[1:3]:", s1)
	fmt.Println("s[:3]:", s2)
	fmt.Println("s[2:]:", s3)
	fmt.Println("s[:]:", s4)

	// 三索引切片：s[low:high:max]
	s5 := s[1:3:3] // [2, 3], len=2, cap=2
	fmt.Printf("s[1:3:3]: %v, len: %d, cap: %d\n", s5, len(s5), cap(s5))

	// 3. append 操作
	fmt.Println("\n=== Append Operations ===")

	s = []int{1, 2, 3}
	fmt.Println("Original:", s)

	// 追加单个元素
	s = append(s, 4)
	fmt.Println("After append(s, 4):", s)

	// 追加多个元素
	s = append(s, 5, 6, 7)
	fmt.Println("After append(s, 5, 6, 7):", s)

	// 追加另一个切片
	s2 = []int{8, 9, 10}
	s = append(s, s2...)
	fmt.Println("After append(s, s2...):", s)

	// 4. copy 操作
	fmt.Println("\n=== Copy Operations ===")

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	n := copy(dst, src)
	fmt.Printf("Copied %d elements: %v\n", n, dst)

	// 部分复制
	dst2 := make([]int, 3)
	n = copy(dst2, src)
	fmt.Printf("Copied %d elements to smaller slice: %v\n", n, dst2)

	// 目标切片更大
	dst3 := make([]int, 10)
	n = copy(dst3, src)
	fmt.Printf("Copied %d elements to larger slice: %v\n", n, dst3)

	// 5. 遍历切片
	fmt.Println("\n=== Traverse Slice ===")

	nums := []int{10, 20, 30, 40, 50}

	// 使用索引
	fmt.Println("Using index:")
	for i := 0; i < len(nums); i++ {
		fmt.Printf("nums[%d] = %d\n", i, nums[i])
	}

	// 使用 range
	fmt.Println("\nUsing range:")
	for i, v := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// 只要值
	fmt.Println("\nOnly values:")
	for _, v := range nums {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// 只要索引
	fmt.Println("\nOnly indices:")
	for i := range nums {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 6. 切片的长度和容量
	fmt.Println("\n=== Length and Capacity ===")

	s = make([]int, 3, 5)
	fmt.Printf("Initial: len=%d, cap=%d, %v\n", len(s), cap(s), s)

	s = append(s, 1)
	fmt.Printf("After append 1: len=%d, cap=%d, %v\n", len(s), cap(s), s)

	s = append(s, 2)
	fmt.Printf("After append 2: len=%d, cap=%d, %v\n", len(s), cap(s), s)

	s = append(s, 3) // 触发扩容
	fmt.Printf("After append 3 (expand): len=%d, cap=%d, %v\n", len(s), cap(s), s)
}
