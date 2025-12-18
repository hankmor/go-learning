# 日志管理示例

本目录包含 Go 语言日志管理的完整示例，对应教程文章：[日志管理：从 log 到 zap](https://hankmo.com/posts/golang/20250718-工程实践-日志管理/)

## 目录结构

```
06-loggging/
├── log/              # 标准库 log 示例
│   └── log.go       # 基础用法、自定义 logger、文件输出
├── slog/            # Go 1.21+ slog 示例
│   ├── basic.go     # 基础用法、日志级别、子 logger
│   └── file_logging.go  # 文件输出、结构化日志
├── zap/             # Uber zap 高性能日志
│   ├── demo.go      # 完整示例：开发/生产环境、日志轮转、动态级别
│   ├── go.mod
│   └── go.sum
├── logrus/          # logrus 示例（已有）
└── compare/         # 性能对比（已有）
```

## 运行示例

### 1. 标准库 log

```bash
cd log
go run log.go
```

### 2. slog (Go 1.21+)

```bash
cd slog
go run basic.go
go run file_logging.go
```

### 3. zap

```bash
cd zap
go run demo.go
```

## 主要特性对比

| 特性     | log      | slog     | zap        |
| -------- | -------- | -------- | ---------- |
| 性能     | 中       | 中       | 高         |
| 结构化   | ❌       | ✅       | ✅         |
| 学习成本 | 低       | 低       | 中         |
| 适用场景 | 简单脚本 | 一般应用 | 高性能服务 |

## 推荐使用

- **简单脚本/工具**：使用标准库 `log`
- **一般 Web 应用**：使用 `slog`（Go 1.21+）
- **高性能服务**：使用 `zap`

## 相关文章

- [日志管理：从 log 到 zap](https://hankmo.com/posts/golang/20250718-工程实践-日志管理/)
