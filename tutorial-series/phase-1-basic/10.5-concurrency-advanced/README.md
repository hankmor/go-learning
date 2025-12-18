# 并发进阶示例

本目录包含 Go 并发进阶的各种示例代码。

## 文件说明

- `main.go` - 基础并发工具示例（WaitGroup、Mutex、RWMutex、Once、Cond）
- `semaphore.go` - 信号量示例
- `errgroup.go` - errgroup 示例

## 运行示例

```bash
# 运行基础示例
go run main.go

# 运行信号量示例
go run semaphore.go

# 运行 errgroup 示例
go run errgroup.go
```

## 依赖安装

```bash
# 安装 golang.org/x/sync 包
go get golang.org/x/sync/semaphore
go get golang.org/x/sync/errgroup
```

## 学习要点

1. **WaitGroup**：等待一组 goroutine 完成
2. **Mutex**：互斥锁，保护共享资源
3. **RWMutex**：读写锁，读多写少场景优化
4. **Once**：确保某个操作只执行一次
5. **Cond**：条件变量，协调 goroutine
6. **Semaphore**：信号量，限制并发数
7. **errgroup**：带错误处理的 WaitGroup
