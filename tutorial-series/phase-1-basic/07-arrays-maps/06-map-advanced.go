package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

func main() {
	// 1. 统计词频
	fmt.Println("=== Word Frequency ===")

	text := "hello world hello go world go go"
	wordCount := make(map[string]int)

	for _, word := range strings.Fields(text) {
		wordCount[word]++
	}

	fmt.Println("Word count:", wordCount)

	// 2. 字符频率统计
	fmt.Println("\n=== Character Frequency ===")

	str := "hello"
	charCount := make(map[rune]int)

	for _, ch := range str {
		charCount[ch]++
	}

	for ch, count := range charCount {
		fmt.Printf("%c: %d\n", ch, count)
	}

	// 3. 分组
	fmt.Println("\n=== Grouping ===")

	type Student struct {
		Name  string
		Grade string
	}

	students := []Student{
		{"Alice", "A"},
		{"Bob", "B"},
		{"Charlie", "A"},
		{"David", "B"},
		{"Eve", "A"},
	}

	// 按成绩分组
	gradeGroups := make(map[string][]Student)
	for _, student := range students {
		gradeGroups[student.Grade] = append(gradeGroups[student.Grade], student)
	}

	for grade, group := range gradeGroups {
		fmt.Printf("Grade %s: %v\n", grade, group)
	}

	// 4. 有序遍历 Map
	fmt.Println("\n=== Ordered Map Traversal ===")

	m := map[string]int{
		"banana": 3,
		"apple":  5,
		"cherry": 2,
		"date":   4,
	}

	// 提取键并排序
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 按排序后的键遍历
	fmt.Println("Sorted by key:")
	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, m[key])
	}

	// 5. Map 的并发安全问题
	fmt.Println("\n=== Map Concurrency Issue ===")

	// 不安全的并发访问（会 panic）
	// unsafeMap := make(map[int]int)
	// var wg sync.WaitGroup
	// for i := 0; i < 100; i++ {
	//     wg.Add(1)
	//     go func(n int) {
	//         defer wg.Done()
	//         unsafeMap[n] = n // 并发写入会 panic
	//     }(i)
	// }
	// wg.Wait()

	// 方案 1：使用 sync.RWMutex
	fmt.Println("Solution 1: sync.RWMutex")
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			safeMap.Set(n, n*n)
		}(i)
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		if val, ok := safeMap.Get(i); ok {
			fmt.Printf("%d: %d\n", i, val)
		}
	}

	// 方案 2：使用 sync.Map
	fmt.Println("\nSolution 2: sync.Map")
	var syncMap sync.Map

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			syncMap.Store(n, n*n)
		}(i)
	}
	wg.Wait()

	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})

	// 6. Map 作为集合（Set）
	fmt.Println("\n=== Map as Set ===")

	// 使用 map[T]bool 实现集合
	set := make(map[int]bool)

	// 添加元素
	set[1] = true
	set[2] = true
	set[3] = true

	// 检查元素是否存在
	fmt.Println("Contains 2:", set[2])
	fmt.Println("Contains 4:", set[4])

	// 删除元素
	delete(set, 2)

	// 遍历集合
	fmt.Print("Set elements: ")
	for key := range set {
		fmt.Print(key, " ")
	}
	fmt.Println()

	// 使用 map[T]struct{} 更节省内存
	set2 := make(map[int]struct{})
	set2[1] = struct{}{}
	set2[2] = struct{}{}

	if _, ok := set2[1]; ok {
		fmt.Println("1 is in set2")
	}

	// 7. 缓存实现
	fmt.Println("\n=== Simple Cache ===")

	cache := NewCache()

	// 设置缓存
	cache.Set("user:1", "Alice")
	cache.Set("user:2", "Bob")

	// 获取缓存
	if val, ok := cache.Get("user:1"); ok {
		fmt.Println("Found in cache:", val)
	}

	// 删除缓存
	cache.Delete("user:2")

	// 清空缓存
	cache.Clear()
	fmt.Println("Cache size after clear:", cache.Size())
}

// SafeMap 线程安全的 Map
type SafeMap struct {
	mu sync.RWMutex
	m  map[int]int
}

// NewSafeMap 创建 SafeMap
func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[int]int),
	}
}

// Set 设置键值对
func (sm *SafeMap) Set(key, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

// Get 获取值
func (sm *SafeMap) Get(key int) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, ok := sm.m[key]
	return val, ok
}

// Delete 删除键
func (sm *SafeMap) Delete(key int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.m, key)
}

// Cache 简单的缓存实现
type Cache struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

// NewCache 创建缓存
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

// Set 设置缓存
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Get 获取缓存
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

// Delete 删除缓存
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

// Clear 清空缓存
func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]interface{})
}

// Size 获取缓存大小
func (c *Cache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.data)
}
