# 快速开始

## 运行所有测试
```bash
go test -v
```

## 运行模糊测试（3秒）
```bash
# JSON 解析器模糊测试
go test -fuzz=FuzzParseJSON -fuzztime=3s

# 字符串反转模糊测试
go test -fuzz=FuzzReverse -fuzztime=3s

# 深度计算模糊测试
go test -fuzz=FuzzCalculateDepth -fuzztime=3s
```

## 查看测试覆盖率
```bash
go test -cover
# 输出: coverage: 96.7% of statements
```

## 运行基准测试
```bash
go test -bench=. -benchmem
```

## 测试结果示例

### 单元测试
```
PASS
coverage: 96.7% of statements
ok      github.com/hankmor/gotesting/fuzz       0.351s
```

### 模糊测试
```
fuzz: elapsed: 3s, execs: 1360687 (453562/sec), new interesting: 274 (total: 284)
PASS
```
3 秒内执行了 136 万次测试，每秒 45 万次！

### 基准测试
```
BenchmarkParseJSON-12            1732294               653.5 ns/op
BenchmarkParseJSONDeep-12         224982              5093 ns/op
```

## 文件说明

- `reverse.go` - 字符串反转函数（处理 UTF-8）
- `reverse_test.go` - 字符串反转测试
- `json_parser.go` - JSON 解析器（带深度限制）
- `json_parser_test.go` - JSON 解析器测试（包含 3 个模糊测试）
- `example_test.go` - 示例代码
- `README.md` - 详细文档
- `QUICKSTART.md` - 本文件
