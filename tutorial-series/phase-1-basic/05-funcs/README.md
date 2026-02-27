# Go 函数与方法示例

本目录包含 Go 语言函数和方法的完整示例代码，对应博客文章《GoLang教程——函数与方法》。

## 目录结构

### 基础示例
- `01-basic.go` - 函数和方法的基础用法
  - 普通函数定义和调用
  - 多返回值
  - 值接收者和指针接收者方法

### 函数特性
- `02-named-return.go` - 命名返回值
  - 命名返回值的使用
  - 裸 return 语句
  - 命名返回值的最佳实践

- `03-variadic.go` - 可变参数
  - 可变参数函数定义
  - 切片展开
  - 可变参数的实际应用

- `04-closure.go` - 匿名函数和闭包
  - 匿名函数
  - 闭包捕获外部变量
  - 闭包的实际应用（计数器、生成器）
  - 高阶函数（filter）

- `05-first-class.go` - 函数作为一等公民
  - 函数类型定义
  - 函数作为参数和返回值
  - 函数组合
  - map、reduce 等高阶函数
  - 函数切片

- `06-defer-advanced.go` - defer 语句详解
  - defer 的执行顺序（LIFO）
  - defer 参数立即求值
  - defer 修改命名返回值
  - defer 与闭包
  - defer 测量执行时间
  - defer 在循环中的陷阱
  - defer 与 panic/recover

### 方法
- `07-method.go` - 方法详解
  - 值接收者 vs 指针接收者
  - 方法链式调用
  - 方法值（Method Value）
  - 方法表达式（Method Expression）

- `08-method-set.go` - 方法集规则
  - 方法集的定义
  - 值类型和指针类型的方法集差异
  - 接口实现的差异
  - 嵌入类型的方法提升
  - 方法集最佳实践

### 错误处理
- `09-error-handling.go` - 错误处理最佳实践
  - 基本错误处理
  - 错误包装（%w）
  - 自定义错误类型
  - errors.Is 和 errors.As
  - 多个错误处理
  - panic 和 recover

## 运行示例

每个文件都是独立的可执行程序，可以直接运行：

```bash
# 运行基础示例
go run 01-basic.go

# 运行命名返回值示例
go run 02-named-return.go

# 运行可变参数示例
go run 03-variadic.go

# 运行闭包示例
go run 04-closure.go

# 运行函数作为值示例
go run 05-first-class.go

# 运行 defer 示例
go run 06-defer-advanced.go

# 运行方法示例
go run 07-method.go

# 运行方法集示例
go run 08-method-set.go

# 运行错误处理示例
go run 09-error-handling.go
```

## 学习路径

建议按照以下顺序学习：

1. **基础** (01-basic.go)
   - 理解函数和方法的基本概念
   - 掌握值接收者和指针接收者的区别

2. **函数特性** (02-05)
   - 命名返回值
   - 可变参数
   - 闭包和匿名函数
   - 函数作为一等公民

3. **defer** (06-defer-advanced.go)
   - defer 的三个重要特性
   - defer 的实际应用场景

4. **方法深入** (07-08)
   - 方法的高级用法
   - 方法集规则和接口实现

5. **错误处理** (09-error-handling.go)
   - Go 的错误处理哲学
   - 错误处理最佳实践

## 关键概念

### 函数
- 函数是一等公民，可以作为值传递
- 支持多返回值，常用于错误处理
- 支持命名返回值和裸 return
- 支持可变参数
- 支持闭包

### 方法
- 方法是带接收者的函数
- 值接收者：复制数据，不能修改原对象
- 指针接收者：传递引用，可以修改原对象
- 方法集规则影响接口实现

### defer
- 延迟执行，直到函数返回
- LIFO 顺序执行
- 参数立即求值
- 可以修改命名返回值
- 常用于资源清理

### 错误处理
- 使用显式错误返回，不使用异常
- 使用 %w 包装错误保留错误链
- 使用 errors.Is 和 errors.As 判断错误类型
- panic 只用于不可恢复的错误

## 最佳实践

1. **优先使用指针接收者**
   - 避免大对象拷贝
   - 能够修改对象状态
   - 保持方法接收者类型一致

2. **合理使用 defer**
   - 确保资源释放
   - 注意性能开销
   - 小心循环中的 defer

3. **错误处理要明确**
   - 不要忽略错误
   - 使用 %w 包装错误
   - 提供有用的错误上下文

4. **闭包要小心**
   - 注意变量捕获
   - 特别是在循环中使用闭包

5. **命名返回值要谨慎**
   - 简单函数中使用
   - 复杂函数中可能降低可读性

## 相关资源

- [Go 官方文档 - Functions](https://go.dev/tour/basics/4)
- [Go 官方文档 - Methods](https://go.dev/tour/methods/1)
- [Effective Go - Functions](https://go.dev/doc/effective_go#functions)
- [Effective Go - Methods](https://go.dev/doc/effective_go#methods)
- [Go Blog - Error handling](https://go.dev/blog/error-handling-and-go)
- [Go Blog - Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)

## 练习题

完成以下练习以巩固所学知识：

1. 编写一个函数 `div(a, b int) (result int, err error)`，实现除法运算，使用命名返回值
2. 为 `User` 结构体添加一个 `SetAge(age int)` 方法，确保能修改用户的年龄
3. 编写一个 `filter` 函数，接收一个整数切片和一个过滤函数
4. 使用闭包实现一个计数器生成器
5. 编写一个函数，使用 defer 来测量函数的执行时间
6. 实现一个自定义错误类型 `RangeError`，包含最小值、最大值和实际值信息

## 贡献

欢迎提交 Issue 和 Pull Request 来改进这些示例！
