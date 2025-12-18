# 日志管理集成

本项目集成了三种日志库：标准库 log、Go 1.21+ slog 和 Uber zap。

## 使用方法

### 1. 通过环境变量选择日志类型

```bash
# 使用标准库 log
LOGGER_TYPE=log go run cmd/server/main.go

# 使用 slog（默认）
LOGGER_TYPE=slog go run cmd/server/main.go

# 使用 zap
LOGGER_TYPE=zap go run cmd/server/main.go
```

### 2. 日志配置

在 `main.go` 中可以配置日志参数：

```go
log := logger.New(&logger.Config{
    Type:       logger.LoggerType(loggerType),
    Level:      "info",          // debug, info, warn, error
    OutputPath: "logs/app.log",  // 输出路径
    MaxSize:    100,             // 日志文件最大大小（MB）
    MaxBackups: 3,               // 保留的旧日志文件数量
    MaxAge:     7,               // 保留的旧日志文件天数
    Compress:   true,            // 是否压缩旧日志文件
})
```

### 3. 使用日志

```go
// 在代码中使用
log.Debug("Debug message", "key", "value")
log.Info("Info message", "user_id", 123)
log.Warn("Warning message")
log.Error("Error message", "error", err)
log.Fatal("Fatal error", "error", err) // 会退出程序
```

## 功能特性

### 1. 统一接口

所有三种日志库都实现了相同的接口：

```go
type Logger interface {
    Debug(msg string, fields ...interface{})
    Info(msg string, fields ...interface{})
    Warn(msg string, fields ...interface{})
    Error(msg string, fields ...interface{})
    Fatal(msg string, fields ...interface{})
}
```

### 2. HTTP 请求日志

通过 `LoggingMiddleware` 自动记录所有 HTTP 请求：

```json
{
  "level": "INFO",
  "msg": "HTTP Request",
  "method": "GET",
  "path": "/api/posts",
  "status": 200,
  "duration_ms": 15,
  "client_ip": "127.0.0.1",
  "user_agent": "Mozilla/5.0..."
}
```

### 3. 日志轮转

使用 lumberjack 实现日志文件自动轮转：
- 当日志文件达到指定大小时自动创建新文件
- 自动删除过期的日志文件
- 可选择压缩旧日志文件

## 日志类型对比

| 特性     | log      | slog     | zap        |
| -------- | -------- | -------- | ---------- |
| 性能     | 中       | 中       | 高         |
| 结构化   | ❌       | ✅       | ✅         |
| 学习成本 | 低       | 低       | 中         |
| 适用场景 | 简单应用 | 一般应用 | 高性能服务 |

## 示例输出

### slog (JSON 格式)

```json
{
  "time": "2025-12-18T14:00:00Z",
  "level": "INFO",
  "msg": "Server started",
  "address": "http://localhost:8080"
}
```

### zap (JSON 格式)

```json
{
  "level": "info",
  "ts": 1702900800.123,
  "msg": "Server started",
  "address": "http://localhost:8080"
}
```

## 测试

运行日志包的单元测试：

```bash
cd internal/logger
go test -v
```

## 相关文章

- [日志管理：从 log 到 zap](https://hankmo.com/posts/golang/20250718-工程实践-日志管理/)
