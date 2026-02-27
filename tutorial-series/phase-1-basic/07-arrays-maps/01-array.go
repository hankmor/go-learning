package main

import "fmt"

func main() {
	// 1. 数组的多种初始化方式
	fmt.Println("=== Array Initialization ===")

	// 指定长度和初始值
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println("arr1:", arr1)

	// 类型推导
	arr2 := [3]int{1, 2, 3}
	fmt.Println("arr2:", arr2)

	// 让编译器推导长度
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("arr3: %v, len: %d\n", arr3, len(arr3))

	// 指定索引初始化
	arr4 := [5]int{0: 10, 2: 20, 4: 30}
	fmt.Println("arr4:", arr4) // [10, 0, 20, 0, 30]

	// 部分初始化（其余为零值）
	arr5 := [5]int{1, 2}
	fmt.Println("arr5:", arr5) // [1, 2, 0, 0, 0]

	// 2. 数组的基本操作
	fmt.Println("\n=== Array Operations ===")

	arr := [5]int{10, 20, 30, 40, 50}

	// 访问元素
	fmt.Println("arr[0]:", arr[0])
	fmt.Println("arr[4]:", arr[4])

	// 修改元素
	arr[1] = 25
	fmt.Println("After modification:", arr)

	// 获取长度
	fmt.Println("Length:", len(arr))

	// 遍历数组
	fmt.Println("\nTraverse with index:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}

	fmt.Println("\nTraverse with range:")
	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// 3. 数组比较
	fmt.Println("\n=== Array Comparison ===")

	arr6 := [3]int{1, 2, 3}
	arr7 := [3]int{1, 2, 3}
	arr8 := [3]int{1, 2, 4}

	fmt.Println("arr6 == arr7:", arr6 == arr7) // true
	fmt.Println("arr6 == arr8:", arr6 == arr8) // false

	// 不同长度的数组不能比较
	// arr9 := [4]int{1, 2, 3, 4}
	// fmt.Println(arr6 == arr9) // 编译错误

	// 4. 多维数组
	fmt.Println("\n=== Multi-dimensional Array ===")

	// 二维数组
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("Matrix:")
	for i, row := range matrix {
		for j, val := range row {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, val)
		}
	}

	// 5. 数组作为函数参数
	fmt.Println("\n=== Array as Function Parameter ===")

	original := [3]int{1, 2, 3}
	fmt.Println("Original before:", original)

	// 值传递：不会修改原数组
	modifyArray(original)
	fmt.Println("Original after modifyArray:", original)

	// 指针传递：会修改原数组
	modifyArrayPtr(&original)
	fmt.Println("Original after modifyArrayPtr:", original)

	// 6. 数组的实际应用场景
	fmt.Println("\n=== Practical Use Cases ===")

	// 固定大小的缓冲区
	var buffer [1024]byte

	// 查找表
	daysInMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	fmt.Println("Days in March:", daysInMonth[2])

	// 矩阵运算
	identity := [3][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Println("Identity matrix:", identity)

	_ = buffer // 避免未使用变量警告
}

// modifyArray 值传递：会复制整个数组
func modifyArray(arr [3]int) {
	arr[0] = 100
	fmt.Println("Inside modifyArray:", arr)
}

// modifyArrayPtr 指针传递：可以修改原数组
func modifyArrayPtr(arr *[3]int) {
	arr[0] = 200
	fmt.Println("Inside modifyArrayPtr:", *arr)
}
