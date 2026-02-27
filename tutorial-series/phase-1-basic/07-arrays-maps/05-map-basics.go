package main

import "fmt"

func main() {
	// 1. Map 的创建方式
	fmt.Println("=== Map Creation ===")

	// 使用 make
	m1 := make(map[string]int)
	m1["a"] = 1
	m1["b"] = 2
	fmt.Println("m1:", m1)

	// 使用字面量
	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println("m2:", m2)

	// 指定初始容量（性能优化）
	m3 := make(map[string]int, 100)
	fmt.Printf("m3 len: %d\n", len(m3))

	// nil map（不能写入）
	var m4 map[string]int
	fmt.Printf("m4: %v, is nil: %t\n", m4, m4 == nil)
	// m4["a"] = 1 // panic: assignment to entry in nil map

	// 2. Map 的基本操作
	fmt.Println("\n=== Map Operations ===")

	m := make(map[string]int)

	// 添加/修改
	m["Alice"] = 95
	m["Bob"] = 88
	m["Charlie"] = 92
	fmt.Println("After adding:", m)

	// 读取
	score := m["Alice"]
	fmt.Println("Alice's score:", score)

	// 读取不存在的键（返回零值）
	score = m["David"]
	fmt.Println("David's score (not exist):", score) // 0

	// 检查键是否存在
	if score, ok := m["Alice"]; ok {
		fmt.Println("Alice found:", score)
	} else {
		fmt.Println("Alice not found")
	}

	if score, ok := m["David"]; ok {
		fmt.Println("David found:", score)
	} else {
		fmt.Println("David not found")
	}

	// 删除
	delete(m, "Bob")
	fmt.Println("After delete Bob:", m)

	// 删除不存在的键（不会报错）
	delete(m, "NonExist")

	// 获取长度
	fmt.Println("Map length:", len(m))

	// 3. 遍历 Map
	fmt.Println("\n=== Traverse Map ===")

	m = map[string]int{
		"Alice":   95,
		"Bob":     88,
		"Charlie": 92,
		"David":   85,
	}

	// 遍历键值对
	fmt.Println("Key-Value pairs:")
	for key, value := range m {
		fmt.Printf("%s: %d\n", key, value)
	}

	// 只遍历键
	fmt.Println("\nOnly keys:")
	for key := range m {
		fmt.Print(key, " ")
	}
	fmt.Println()

	// 只遍历值
	fmt.Println("\nOnly values:")
	for _, value := range m {
		fmt.Print(value, " ")
	}
	fmt.Println()

	// 注意：遍历顺序是随机的
	fmt.Println("\nMultiple iterations (order may differ):")
	for i := 0; i < 3; i++ {
		fmt.Printf("Iteration %d: ", i+1)
		for key := range m {
			fmt.Print(key, " ")
		}
		fmt.Println()
	}

	// 4. Map 的键类型
	fmt.Println("\n=== Map Key Types ===")

	// 整数键
	intMap := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	fmt.Println("Int map:", intMap)

	// 结构体键（结构体必须可比较）
	type Point struct {
		X, Y int
	}
	pointMap := map[Point]string{
		{1, 2}: "A",
		{3, 4}: "B",
	}
	fmt.Println("Point map:", pointMap)

	// 数组键（数组可以作为键）
	arrayMap := map[[2]int]string{
		{1, 2}: "A",
		{3, 4}: "B",
	}
	fmt.Println("Array map:", arrayMap)

	// 切片不能作为键（切片不可比较）
	// sliceMap := map[[]int]string{} // 编译错误

	// 5. Map 的值类型
	fmt.Println("\n=== Map Value Types ===")

	// 值是切片
	sliceValueMap := make(map[string][]int)
	sliceValueMap["a"] = []int{1, 2, 3}
	sliceValueMap["a"] = append(sliceValueMap["a"], 4)
	fmt.Println("Slice value map:", sliceValueMap)

	// 值是 Map
	mapValueMap := make(map[string]map[string]int)
	mapValueMap["user1"] = make(map[string]int)
	mapValueMap["user1"]["score"] = 95
	fmt.Println("Map value map:", mapValueMap)

	// 值是结构体
	type Student struct {
		Name  string
		Score int
	}
	studentMap := map[int]Student{
		1: {"Alice", 95},
		2: {"Bob", 88},
	}
	fmt.Println("Struct value map:", studentMap)

	// 6. 清空 Map
	fmt.Println("\n=== Clear Map ===")

	m = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Before clear:", m)

	// 方式 1：逐个删除
	for key := range m {
		delete(m, key)
	}
	fmt.Println("After clear (delete):", m)

	// 方式 2：重新赋值
	m = map[string]int{"a": 1, "b": 2, "c": 3}
	m = make(map[string]int)
	fmt.Println("After clear (reassign):", m)

	// 方式 3：赋值为 nil
	m = map[string]int{"a": 1, "b": 2, "c": 3}
	m = nil
	fmt.Printf("After clear (nil): %v, is nil: %t\n", m, m == nil)
}
