# Go 数组、切片与映射示例

本目录包含 Go 语言数组、切片和映射的完整示例代码，对应博客文章《GoLang教程——数组、切片与映射》。

## 目录结构

### 数组
- `01-array.go` - 数组的完整示例
  - 数组的多种初始化方式
  - 数组的基本操作
  - 数组比较
  - 多维数组
  - 数组作为函数参数
  - 实际应用场景

### 切片基础
- `02-slice-basics.go` - 切片的基础操作
  - 切片的创建方式
  - nil 切片 vs 空切片
  - 切片的切片操作
  - append 和 copy 操作
  - 遍历切片
  - 长度和容量

### 切片高级
- `03-slice-advanced.go` - 切片的高级操作
  - 切片扩容机制
  - 删除元素
  - 插入元素
  - 过滤元素
  - 反转切片
  - 去重
  - 合并切片
  - 切片作为栈和队列
  - 二维切片

- `04-slice-pitfalls.go` - 切片的常见陷阱
  - 共享底层数组
  - append 可能改变底层数组
  - 循环中的切片扩容
  - 切片作为函数参数
  - nil 切片 vs 空切片
  - 切片的比较
  - 切片的零值

### 映射基础
- `05-map-basics.go` - Map 的基础操作
  - Map 的创建方式
  - Map 的基本操作
  - 遍历 Map
  - Map 的键类型
  - Map 的值类型
  - 清空 Map

### 映射高级
- `06-map-advanced.go` - Map 的高级应用
  - 统计词频
  - 字符频率统计
  - 分组
  - 有序遍历 Map
  - Map 的并发安全
  - Map 作为集合（Set）
  - 缓存实现

## 运行示例

每个文件都是独立的可执行程序，可以直接运行：

```bash
# 运行数组示例
go run 01-array.go

# 运行切片基础示例
go run 02-slice-basics.go

# 运行切片高级示例
go run 03-slice-advanced.go

# 运行切片陷阱示例
go run 04-slice-pitfalls.go

# 运行 Map 基础示例
go run 05-map-basics.go

# 运行 Map 高级示例
go run 06-map-advanced.go
```

## 学习路径

建议按照以下顺序学习：

1. **数组基础** (01-array.go)
   - 理解数组的特性和限制
   - 掌握数组的基本操作

2. **切片基础** (02-slice-basics.go)
   - 理解切片的内部结构
   - 掌握切片的创建和基本操作

3. **切片高级** (03-slice-advanced.go)
   - 掌握切片的高级操作
   - 学习常见的切片算法

4. **切片陷阱** (04-slice-pitfalls.go)
   - 理解切片的常见陷阱
   - 学习如何避免这些问题

5. **Map 基础** (05-map-basics.go)
   - 理解 Map 的特性
   - 掌握 Map 的基本操作

6. **Map 高级** (06-map-advanced.go)
   - 掌握 Map 的高级应用
   - 学习并发安全的 Map 使用

## 关键概念

### 数组
- 长度是类型的一部分
- 值类型，赋值和传参会复制
- 长度固定，不能改变
- 内存连续存储

### 切片
- 动态数组，长度可变
- 引用类型，包含指针、长度和容量
- 底层是数组
- append 可能导致扩容
- 共享底层数组要小心

### Map
- 无序的键值对集合
- 引用类型
- 键必须可比较
- 不是线程安全的
- nil map 不能写入

## 性能优化

### 切片优化
1. 预分配容量：`make([]int, 0, 100)`
2. 避免频繁扩容
3. 使用 copy 而不是 append 合并切片

### Map 优化
1. 预分配容量：`make(map[string]int, 1000)`
2. 避免频繁创建和销毁
3. 使用 sync.Map 处理并发

## 最佳实践

1. **优先使用切片而不是数组**
   - 除非有特殊需求（如固定大小的缓冲区）

2. **预分配容量**
   - 如果知道大致大小，使用 make 预分配

3. **使用 copy 创建独立副本**
   - 避免共享底层数组的问题

4. **检查 map 键是否存在**
   - 使用 `value, ok := m[key]` 模式

5. **并发访问 map 要加锁**
   - 或使用 sync.Map

6. **不要依赖 map 的遍历顺序**
   - 如果需要有序，使用切片存储键

7. **nil 切片可以直接 append**
   - 但 nil map 不能写入

## 常见陷阱

### 切片共享底层数组
```go
// 错误
s1 := []int{1, 2, 3}
s2 := s1
s2[0] = 100 // s1 也被修改

// 正确
s2 := make([]int, len(s1))
copy(s2, s1)
```

### append 可能改变底层数组
```go
// 错误
s1 := []int{1, 2, 3}
s2 := s1[:2]
s2 = append(s2, 100) // 可能修改 s1

// 正确
s2 := s1[:2:2] // 限制容量
s2 = append(s2, 100)
```

### 切片作为函数参数
```go
// 错误
func appendSlice(s []int) {
    s = append(s, 4) // 不会影响原切片
}

// 正确
func appendSlice(s []int) []int {
    return append(s, 4)
}
```

### Map 并发访问
```go
// 错误
m := make(map[int]int)
go func() { m[1] = 1 }()
go func() { m[2] = 2 }() // panic

// 正确
var mu sync.RWMutex
mu.Lock()
m[1] = 1
mu.Unlock()
```

## 练习题

完成以下练习以巩固所学知识：

1. 创建一个包含 10 个元素的整型切片，使用 range 遍历并打印所有偶数
2. 统计一段英文文本中每个单词出现的次数，使用 map[string]int 存储并打印结果
3. 实现一个函数，删除切片中的重复元素
4. 实现一个函数，反转切片中的元素
5. 使用 map 实现一个简单的缓存系统，支持 Get、Set 和 Delete 操作
6. 编写代码演示切片共享底层数组的问题，并给出解决方案
7. 实现一个线程安全的计数器，使用 map 存储不同键的计数
8. 实现一个函数，合并两个有序切片为一个有序切片

## 相关资源

- [Go 官方文档 - Slices](https://go.dev/blog/slices-intro)
- [Go 官方文档 - Maps](https://go.dev/blog/maps)
- [Go 官方文档 - Arrays, slices (and strings)](https://go.dev/blog/slices)
- [Effective Go - Slices](https://go.dev/doc/effective_go#slices)
- [Effective Go - Maps](https://go.dev/doc/effective_go#maps)
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)

## 贡献

欢迎提交 Issue 和 Pull Request 来改进这些示例！
