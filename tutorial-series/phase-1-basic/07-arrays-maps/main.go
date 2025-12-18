package main

import "fmt"

func main() {
    // 1. 数组 (Array)
    // 长度是类型的一部分，固定不可变
    var arr [3]int = [3]int{10, 20, 30}
    fmt.Println("Array:", arr)

    // 2. 切片 (Slice)
    // 动态数组，引用类型
    slice := []int{1, 2, 3, 4, 5}
    
    // 操作：切取子集 [start:end] (左闭右开)
    subSlice := slice[1:3] // 包含索引 1, 2 的元素 -> {2, 3}
    fmt.Println("SubSlice:", subSlice)

    // 操作：追加元素
    // 当容量不足时，append 会自动扩容
    slice = append(slice, 6)
    fmt.Println("Appended Slice:", slice)

    // 3. 映射 (Map)
    // 键值对集合，类似 Python 的 dict 或 Java 的 HashMap
    scores := make(map[string]int)
    scores["Alice"] = 95
    scores["Bob"] = 88

    // 检查键是否存在
    // val 是值，ok 是布尔值（存在为 true）
    if score, ok := scores["Alice"]; ok {
        fmt.Printf("Alice's score is %d\n", score)
    }

    // 删除键值对
    delete(scores, "Bob")

    // 遍历 Map (注意：遍历顺序是随机的)
    for name, score := range scores {
        fmt.Printf("%s: %d\n", name, score)
    }
}
