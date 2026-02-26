package main

import "fmt"

// Example_reverse 字符串反转示例
func Example_reverse() {
	input := "The quick brown fox"
	rev, _ := Reverse(input)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)

	// 中文测试
	ch, _ := Reverse("中国加油")
	fmt.Printf("中文: %q\n", ch)

	// Output:
	// original: "The quick brown fox"
	// reversed: "xof nworb kciuq ehT"
	// 中文: "油加国中"
}

// Example_parseJSON JSON 解析示例
func Example_parseJSON() {
	// 测试正常的 JSON
	result, err := ParseJSON(`{"name":"老墨","age":30}`)
	if err != nil {
		fmt.Printf("解析失败: %v\n", err)
	} else {
		fmt.Printf("解析成功: %+v\n", result)
	}

	// 测试嵌套深度限制
	deepJSON := `{"a":{"b":{"c":{"d":{"e":"value"}}}}}`
	result2, err2 := ParseJSON(deepJSON, MaxDepth(3))
	if err2 != nil {
		fmt.Printf("嵌套过深: %v\n", err2)
	} else {
		fmt.Printf("解析成功: %+v\n", result2)
	}

	// 测试数组
	result3, err3 := ParseJSON(`[1,2,3,4,5]`)
	if err3 != nil {
		fmt.Printf("解析失败: %v\n", err3)
	} else {
		fmt.Printf("解析成功: %+v\n", result3)
	}

	// Output:
	// 解析成功: map[age:30 name:老墨]
	// 嵌套过深: JSON 嵌套深度 5 超过最大限制 3
	// 解析成功: [1 2 3 4 5]
}
