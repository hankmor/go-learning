package main

import "fmt"

func main() {
	// 陷阱 1：共享底层数组
	fmt.Println("=== Pitfall 1: Shared Underlying Array ===")

	arr := [5]int{1, 2, 3, 4, 5}
	s1 := arr[1:3] // [2, 3]
	s2 := arr[2:4] // [3, 4]

	fmt.Println("Original array:", arr)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	// 修改 s1 会影响 s2
	s1[1] = 100
	fmt.Println("\nAfter s1[1] = 100:")
	fmt.Println("Array:", arr)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2) // s2[0] 也变成了 100

	// 解决方案：使用 copy 创建独立副本
	fmt.Println("\nSolution: Use copy")
	s3 := make([]int, len(s1))
	copy(s3, s1)
	s3[0] = 200
	fmt.Println("s1:", s1)
	fmt.Println("s3:", s3) // s3 是独立的

	// 陷阱 2：append 可能改变底层数组
	fmt.Println("\n=== Pitfall 2: Append May Modify Underlying Array ===")

	s1 = []int{1, 2, 3}
	s2 = s1[:2] // [1, 2], 但容量是 3

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Printf("s2 len=%d cap=%d\n", len(s2), cap(s2))

	// append 到 s2，如果容量足够，会修改底层数组
	s2 = append(s2, 100)
	fmt.Println("\nAfter append(s2, 100):")
	fmt.Println("s1:", s1) // s1 也被修改了！
	fmt.Println("s2:", s2)

	// 解决方案：使用三索引切片限制容量
	fmt.Println("\nSolution: Use three-index slice")
	s1 = []int{1, 2, 3}
	s2 = s1[:2:2] // len=2, cap=2
	fmt.Printf("s2 len=%d cap=%d\n", len(s2), cap(s2))

	s2 = append(s2, 200) // 会分配新的底层数组
	fmt.Println("After append(s2, 200):")
	fmt.Println("s1:", s1) // s1 不受影响
	fmt.Println("s2:", s2)

	// 陷阱 3：循环中的切片扩容
	fmt.Println("\n=== Pitfall 3: Slice Growth in Loop ===")

	s := []int{1, 2, 3}
	fmt.Println("Original:", s)

	// 错误：在循环中修改切片可能导致无限循环
	// for _, v := range s {
	//     s = append(s, v) // 危险！
	// }

	// 正确：先确定循环次数
	length := len(s)
	for i := 0; i < length; i++ {
		s = append(s, s[i])
	}
	fmt.Println("After doubling:", s)

	// 陷阱 4：切片作为函数参数
	fmt.Println("\n=== Pitfall 4: Slice as Function Parameter ===")

	s = []int{1, 2, 3}
	fmt.Println("Before modifySlice:", s)

	modifySlice(s)
	fmt.Println("After modifySlice:", s) // 元素被修改

	fmt.Println("\nBefore appendSlice:", s)
	appendSlice(s)
	fmt.Println("After appendSlice:", s) // 长度没变！

	// 正确的做法：返回新切片
	s = appendSliceCorrect(s)
	fmt.Println("After appendSliceCorrect:", s)

	// 陷阱 5：nil 切片 vs 空切片
	fmt.Println("\n=== Pitfall 5: Nil Slice vs Empty Slice ===")

	var nilSlice []int
	emptySlice := []int{}

	fmt.Printf("nilSlice: %v, len=%d, cap=%d, is nil: %t\n",
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("emptySlice: %v, len=%d, cap=%d, is nil: %t\n",
		emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)

	// nil 切片可以直接 append
	nilSlice = append(nilSlice, 1)
	fmt.Println("After append to nil slice:", nilSlice)

	// 陷阱 6：切片的比较
	fmt.Println("\n=== Pitfall 6: Slice Comparison ===")

	s1 = []int{1, 2, 3}
	s2 = []int{1, 2, 3}

	// 切片不能直接比较（除了和 nil 比较）
	// fmt.Println(s1 == s2) // 编译错误

	// 只能和 nil 比较
	fmt.Println("s1 == nil:", s1 == nil)

	// 如果需要比较，要手动实现
	fmt.Println("s1 equals s2:", sliceEqual(s1, s2))

	// 陷阱 7：切片的零值
	fmt.Println("\n=== Pitfall 7: Zero Value of Slice ===")

	var s4 []int
	fmt.Printf("Zero value slice: %v, is nil: %t\n", s4, s4 == nil)

	// 可以安全地对 nil 切片进行 len、cap、range、append
	fmt.Println("len(nil slice):", len(s4))
	fmt.Println("cap(nil slice):", cap(s4))

	for range s4 {
		fmt.Println("This won't print")
	}

	s4 = append(s4, 1)
	fmt.Println("After append:", s4)
}

// modifySlice 修改切片元素（会影响原切片）
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 100
	}
}

// appendSlice 追加元素（不会影响原切片的长度）
func appendSlice(s []int) {
	s = append(s, 4)
	fmt.Println("Inside appendSlice:", s)
}

// appendSliceCorrect 正确的追加方式：返回新切片
func appendSliceCorrect(s []int) []int {
	return append(s, 4)
}

// sliceEqual 比较两个切片是否相等
func sliceEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
