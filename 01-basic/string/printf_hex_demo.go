package main

import "fmt"

func main() {
	// 示例字符串（包含一些字节）
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("=== 区别演示 ===")
	fmt.Println()

	// %x：连续输出十六进制，没有空格
	fmt.Println("Printf with %x (no spaces):")
	fmt.Printf("%x\n", sample)
	fmt.Println()

	// % x：在每个字节之间添加空格
	fmt.Println("Printf with % x (with spaces):")
	fmt.Printf("% x\n", sample)
	fmt.Println()

	// 更明显的例子
	fmt.Println("=== 更明显的例子 ===")
	fmt.Println()

	data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f} // "Hello"

	fmt.Println("Data:", string(data))
	fmt.Println()

	fmt.Printf("%%x format:  %x\n", data)
	fmt.Printf("%% x format: % x\n", data)
	fmt.Println()

	// 对比：没有空格 vs 有空格
	fmt.Println("=== 对比 ===")
	fmt.Printf("Without space: %x\n", data)
	fmt.Printf("With space:    % x\n", data)
	fmt.Println()

	// 实际应用场景
	fmt.Println("=== 实际应用 ===")
	ipBytes := []byte{192, 168, 1, 1}
	fmt.Printf("IP bytes (no space): %x\n", ipBytes)
	fmt.Printf("IP bytes (space):    % x\n", ipBytes)
	fmt.Println()

	// 更长的数据
	longData := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	fmt.Println("Long data:")
	fmt.Printf("No space: %x\n", longData)
	fmt.Printf("Space:    % x\n", longData)
}
