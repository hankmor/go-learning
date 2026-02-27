package main

import (
	"errors"
	"fmt"
)

func main() {
	// 1. 基本 goto 用法
	fmt.Println("=== Basic goto ===")
	i := 0
loop:
	fmt.Println(i)
	i++
	if i < 3 {
		goto loop
	}

	// 2. 错误处理中的 goto
	fmt.Println("\n=== goto for error handling ===")
	if err := processWithGoto(); err != nil {
		fmt.Println("Error:", err)
	}

	// 3. 跳出多层嵌套
	fmt.Println("\n=== goto to break nested loops ===")
	breakNestedLoops()

	// 4. goto 的限制
	fmt.Println("\n=== goto limitations ===")
	demonstrateGotoLimitations()
}

// processWithGoto 演示在错误处理中使用 goto
func processWithGoto() error {
	var file, db, cache interface{}
	var err error

	// 模拟打开资源
	file, err = openFile()
	if err != nil {
		goto cleanup
	}

	db, err = openDB()
	if err != nil {
		goto cleanup
	}

	cache, err = openCache()
	if err != nil {
		goto cleanup
	}

	// 处理数据
	err = processData()
	if err != nil {
		goto cleanup
	}

	fmt.Println("All operations successful")
	return nil

cleanup:
	// 统一的清理逻辑
	fmt.Println("Cleaning up resources...")
	if cache != nil {
		fmt.Println("Closing cache")
	}
	if db != nil {
		fmt.Println("Closing database")
	}
	if file != nil {
		fmt.Println("Closing file")
	}
	return err
}

// breakNestedLoops 使用 goto 跳出多层嵌套循环
func breakNestedLoops() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				fmt.Printf("i=%d, j=%d, k=%d\n", i, j, k)
				if i == 1 && j == 1 && k == 1 {
					goto done
				}
			}
		}
	}
done:
	fmt.Println("Exited all loops")
}

// demonstrateGotoLimitations 演示 goto 的限制
func demonstrateGotoLimitations() {
	// 1. goto 不能跳过变量声明
	// goto skip // 错误：跳过了 x 的声明
	// x := 10
	// skip:
	// fmt.Println(x)

	// 2. goto 不能跳入其他作用域
	if true {
		// goto inner // 错误：不能从外部跳入 if 块
	}
	// inner:
	// fmt.Println("Inside if")

	// 3. 正确的用法：在同一作用域内跳转
	x := 10
	if x > 5 {
		goto valid
	}
	fmt.Println("x <= 5")
	return

valid:
	fmt.Println("x > 5")
}

// 模拟函数
func openFile() (interface{}, error) {
	fmt.Println("Opening file...")
	return "file", nil
}

func openDB() (interface{}, error) {
	fmt.Println("Opening database...")
	// 模拟错误
	return nil, errors.New("database connection failed")
}

func openCache() (interface{}, error) {
	fmt.Println("Opening cache...")
	return "cache", nil
}

func processData() error {
	fmt.Println("Processing data...")
	return nil
}
