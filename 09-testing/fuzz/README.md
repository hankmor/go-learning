# Go 模糊测试示例

这个目录包含了 Go 模糊测试的实战示例，配合极客老墨的文章《Go模糊测试实战：让AI帮你找Bug》。

## 示例代码

### 1. 字符串反转（reverse.go）
经典的字符串反转函数，展示了 UTF-8 编码的边界情况。

### 2. JSON 解析器（json_parser.go）
一个带深度限制的 JSON 解析器，展示了如何防止深度嵌套导致的 OOM 问题。

## 运行测试

### 单元测试
```bash
# 运行所有单元测试
go test -v

# 运行特定测试
go test -v -run TestParseJSON
```

### 模糊测试
```bash
# 运行模糊测试（默认运行直到发现问题或手动停止）
go test -fuzz=FuzzParseJSON

# 限制运行时间（10秒）
go test -fuzz=FuzzParseJSON -fuzztime=10s

# 限制运行次数（1000次）
go test -fuzz=FuzzParseJSON -fuzztime=1000x

# 运行所有模糊测试（不推荐，会一直运行）
go test -fuzz=.
```

### 查看测试覆盖率
```bash
# 生成覆盖率报告
go test -cover

# 生成详细的覆盖率报告
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### 基准测试
```bash
# 运行所有基准测试
go test -bench=.

# 运行特定基准测试
go test -bench=BenchmarkParseJSON

# 显示内存分配信息
go test -bench=. -benchmem
```

## 模糊测试语料库

模糊测试会自动生成语料库文件，存储在 `testdata/fuzz/` 目录下：

```
testdata/
└── fuzz/
    ├── FuzzParseJSON/
    │   ├── seed1
    │   └── seed2
    └── FuzzValidateJSON/
        └── seed1
```

这些文件包含了导致测试失败的输入，可以用于回归测试。

## 常见问题

### 1. 模糊测试一直运行怎么办？
按 `Ctrl+C` 停止，或者使用 `-fuzztime` 参数限制运行时间。

### 2. 如何查看模糊测试发现的问题？
模糊测试发现问题后会自动停止，并显示导致问题的输入。这个输入会被保存到语料库中。

### 3. 如何重现模糊测试发现的问题？
```bash
# 模糊测试会显示类似这样的输出：
# failing input: testdata/fuzz/FuzzParseJSON/abc123

# 可以直接运行单元测试来重现
go test -v
```

## 实战技巧

1. **从简单的种子开始**：在 `f.Add()` 中添加典型的输入样本
2. **设置合理的限制**：使用 `-fuzztime` 避免无限运行
3. **结合 AI 工具**：让 AI 帮你生成更多的种子语料
4. **定期运行**：在 CI/CD 中集成模糊测试

## 参考资料

- [Go 官方模糊测试文档](https://go.dev/doc/fuzz/)
- [极客老墨的文章：Go模糊测试实战](https://hankmo.com)
