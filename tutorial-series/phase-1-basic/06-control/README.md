# Go 控制结构示例

本目录包含 Go 语言控制结构的完整示例代码，对应博客文章《GoLang教程——控制结构》。

## 目录结构

### 条件语句
- `01-if-else.go` - if-else 条件语句
  - 基本 if 和 if-else
  - if-else if-else 链
  - if 带初始化语句
  - 常用于错误处理、map 查找、类型断言

### 分支语句
- `02-switch.go` - switch 分支语句
  - 基本 switch
  - 多值 case
  - 带初始化语句的 switch
  - 无表达式的 switch
  - fallthrough 关键字
  - 类型 switch

### 循环语句
- `03-for-basic.go` - for 循环基础
  - 标准 for 循环
  - 类似 while 的形式
  - 无限循环
  - continue 和 break
  - 嵌套循环

- `04-for-range.go` - range 遍历
  - 遍历切片和数组
  - 遍历字符串
  - 遍历 map
  - 遍历 channel
  - 只要索引或只要值

- `05-range-pitfalls.go` - range 常见陷阱
  - 值是拷贝的问题
  - 循环变量地址问题
  - 修改 map 值的问题
  - range 表达式只求值一次

### 流程控制
- `06-labels.go` - 标签和多层循环控制
  - 使用标签跳出外层循环
  - 使用标签 continue 到外层循环
  - 在二维切片中查找值
  - 多层嵌套循环的复杂控制
  - 实际应用示例

- `07-goto.go` - goto 语句
  - 基本 goto 用法
  - 错误处理中的 goto
  - 跳出多层嵌套
  - goto 的限制和注意事项

## 运行示例

每个文件都是独立的可执行程序，可以直接运行：

```bash
# 运行 if-else 示例
go run 01-if-else.go

# 运行 switch 示例
go run 02-switch.go

# 运行 for 基础示例
go run 03-for-basic.go

# 运行 range 示例
go run 04-for-range.go

# 运行 range 陷阱示例
go run 05-range-pitfalls.go

# 运行标签示例
go run 06-labels.go

# 运行 goto 示例
go run 07-goto.go
```

## 学习路径

建议按照以下顺序学习：

1. **条件语句** (01-if-else.go)
   - 理解 if-else 的基本用法
   - 掌握 if 初始化语句的使用场景

2. **分支语句** (02-switch.go)
   - 理解 switch 的灵活性
   - 掌握无表达式 switch 和类型 switch

3. **循环基础** (03-for-basic.go)
   - 掌握 for 的三种形式
   - 理解 break 和 continue

4. **range 遍历** (04-for-range.go)
   - 掌握 range 遍历各种数据结构
   - 理解索引和值的使用

5. **range 陷阱** (05-range-pitfalls.go)
   - 理解值拷贝问题
   - 避免常见的 range 陷阱

6. **高级控制** (06-labels.go, 07-goto.go)
   - 掌握标签的使用
   - 了解 goto 的适用场景

## 关键概念

### if 语句
- 条件表达式不需要小括号
- 大括号是必须的
- 支持初始化语句，变量作用域限定在 if-else 块内
- 常用于错误处理、map 查找、类型断言

### switch 语句
- 默认不需要 break，自动终止
- 支持多值 case
- 支持无表达式形式，替代复杂的 if-else 链
- 支持类型 switch，判断接口变量的实际类型
- fallthrough 强制执行下一个 case

### for 循环
- Go 唯一的循环结构
- 三种形式：标准、类似 while、无限循环
- range 用于遍历集合
- break 跳出循环，continue 跳过本次迭代

### range 遍历
- 遍历数组、切片、字符串、map、channel
- 返回索引和值（或键和值）
- 值是拷贝，修改不影响原数据
- 循环变量地址问题（Go 1.22+ 已修复）

### 标签和 goto
- 标签用于跳出多层循环
- goto 用于错误处理和跳出深层嵌套
- goto 有限制：不能跳过变量声明，不能跳入其他作用域
- 现代 Go 代码中，defer 和结构化控制流通常是更好的选择

## 最佳实践

1. **优先使用 range**
   - 遍历集合时使用 range 而不是索引循环
   - 代码更简洁，意图更明确

2. **注意 range 的值拷贝**
   - 修改元素时使用索引访问
   - 或者使用指针切片

3. **使用无表达式 switch**
   - 替代复杂的 if-else 链
   - 代码更清晰，更易维护

4. **合理使用标签**
   - 跳出多层循环时使用标签
   - 避免复杂的标志变量

5. **避免滥用 goto**
   - 只在必要时使用（错误处理、跳出深层嵌套）
   - 优先考虑 defer 和结构化控制流

6. **if 初始化语句**
   - 限制变量作用域
   - 让代码更简洁

## 常见陷阱

### range 值拷贝
```go
// 错误
for _, p := range people {
    p.Age++ // 无效
}

// 正确
for i := range people {
    people[i].Age++
}
```

### 循环变量地址
```go
// 错误（Go 1.22 之前）
for _, v := range nums {
    ptrs = append(ptrs, &v) // 危险
}

// 正确
for _, v := range nums {
    v := v // 创建新变量
    ptrs = append(ptrs, &v)
}
```

### switch fallthrough
```go
// fallthrough 不判断下一个 case 的条件
switch num {
case 1:
    fmt.Println("One")
    fallthrough // 会执行 case 2
case 2:
    fmt.Println("Two")
}
```

## 练习题

完成以下练习以巩固所学知识：

1. 编写一个函数，使用 switch 判断分数的等级（90以上A，80以上B，70以上C，60以上D，否则F）
2. 使用 for 和 range 遍历一个字符串切片，打印每个字符串及其长度
3. 编写一个函数，使用标签和 break 在二维切片中查找特定值，找到后立即返回其位置
4. 使用类型 switch 编写一个函数，根据不同类型打印不同的信息
5. 编写一个函数，演示 range 的值拷贝陷阱，并给出正确的修改方式
6. 使用 goto 实现一个简单的状态机

## 相关资源

- [Go 官方文档 - For statements](https://go.dev/ref/spec#For_statements)
- [Go 官方文档 - Switch statements](https://go.dev/ref/spec#Switch_statements)
- [Go 官方文档 - If statements](https://go.dev/ref/spec#If_statements)
- [Effective Go - Control structures](https://go.dev/doc/effective_go#control-structures)
- [Go Blog - Go's Declaration Syntax](https://go.dev/blog/declaration-syntax)

## 贡献

欢迎提交 Issue 和 Pull Request 来改进这些示例！
