package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// 1. Store: 存储键值对
	m.Store("name", "Hank")
	m.Store("age", 18)

	// 2. Load: 读取
	val, ok := m.Load("name")
	if ok {
		fmt.Printf("Found name: %v\n", val)
	}

	// 3. LoadOrStore: 如果存在则读取，不存在则存储
	// 返回值 actual 是最终存的值，loaded 为 true 表示是读出来的，false 表示是刚存进去的
	actual, loaded := m.LoadOrStore("city", "Beijing")
	fmt.Printf("City: %v, Loaded: %v\n", actual, loaded)

	actual, loaded = m.LoadOrStore("city", "Shanghai")
	fmt.Printf("City: %v, Loaded: %v\n", actual, loaded) // 应该是 Beijing, true

	// 4. Delete: 删除
	m.Delete("age")

	// 5. Range: 遍历
	fmt.Println("Iterating map:")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true // 返回 true 继续遍历，false 停止
	})
}
