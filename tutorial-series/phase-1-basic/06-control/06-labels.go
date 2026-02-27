package main

import "fmt"

func main() {
	// 1. 使用标签跳出外层循环
	fmt.Println("=== Break outer loop with label ===")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			if i == 1 && j == 1 {
				break outer // 跳出外层循环
			}
		}
	}
	fmt.Println("Exited outer loop")

	// 2. 使用标签 continue 到外层循环
	fmt.Println("\n=== Continue outer loop with label ===")
outer2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue outer2 // 继续外层循环的下一次迭代
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// 3. 在二维切片中查找值
	fmt.Println("\n=== Find value in 2D slice ===")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	target := 5
	found := false
	var row, col int

search:
	for i, r := range matrix {
		for j, v := range r {
			if v == target {
				row, col = i, j
				found = true
				break search
			}
		}
	}

	if found {
		fmt.Printf("Found %d at position (%d, %d)\n", target, row, col)
	} else {
		fmt.Printf("%d not found\n", target)
	}

	// 4. 多层嵌套循环的复杂控制
	fmt.Println("\n=== Complex nested loop control ===")
outerLoop:
	for i := 0; i < 5; i++ {
	middleLoop:
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				if i+j+k == 6 {
					fmt.Printf("Found: i=%d, j=%d, k=%d\n", i, j, k)
					continue middleLoop
				}
				if i+j+k > 10 {
					break outerLoop
				}
			}
		}
	}

	// 5. 标签的作用域
	fmt.Println("\n=== Label scope ===")
	// 标签必须在同一个函数内
	// 标签后面必须跟一个语句
label1:
	for i := 0; i < 3; i++ {
		if i == 1 {
			goto label2
		}
		fmt.Println("Before label2:", i)
	}

label2:
	fmt.Println("At label2")

	// 6. 实际应用：解析嵌套数据
	fmt.Println("\n=== Practical example: Parse nested data ===")
	parseNestedData()
}

// parseNestedData 演示在实际场景中使用标签
func parseNestedData() {
	data := [][][]int{
		{{1, 2}, {3, 4}},
		{{5, 6}, {7, 8}},
		{{9, 10}, {11, 12}},
	}

	target := 7
	var x, y, z int
	found := false

findTarget:
	for i, layer := range data {
		for j, row := range layer {
			for k, val := range row {
				if val == target {
					x, y, z = i, j, k
					found = true
					break findTarget
				}
			}
		}
	}

	if found {
		fmt.Printf("Found %d at position [%d][%d][%d]\n", target, x, y, z)
	}
}
