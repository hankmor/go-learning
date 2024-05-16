# Data Race Detector

数据竞态(Data Race, 也叫竞态条件, Race Condition)是并发系统中最常见和最难调试的错误类型之一，当两个 goroutine 同时访问同一个变量并且至少其中一次访问是写入时，就会发生数据竞争。

go在1.1版本时支持[数据竞态检测](https://go.dev/doc/articles/race_detector)。

使用 `-race` 参数可以用来检测这种竞态：

```shell
$ go test -race mypkg    // to test the package
$ go run -race mysrc.go  // to run the source file
$ go build -race mycmd   // to build the command
$ go install -race mypkg // to install the package
```

示例程序中，两个 `goroutine` 同时读写 `map` 存在竞态，可能造成数据不正确的情况，但是错误难以发现，通过执行时添加 `-race` 选择可以检测：

```shell
 ➜  datarace git:(main) ✗ go run -race race_demo.go 
==================
WARNING: DATA RACE
Write at 0x00c00007e180 by goroutine 6:
  runtime.mapassign_faststr()
      /Users/hank/software/go1.20.5/src/runtime/map_faststr.go:203 +0x0
  main.main.func1()
      /Users/hank/workspace/mine/go-projects/go-learning/02-advanced/goroutine/datarace/race_demo.go:10 +0
x50

Previous write at 0x00c00007e180 by main goroutine:
  runtime.mapassign_faststr()
      /Users/hank/software/go1.20.5/src/runtime/map_faststr.go:203 +0x0
  main.main()
      /Users/hank/workspace/mine/go-projects/go-learning/02-advanced/goroutine/datarace/race_demo.go:13 +0
x13a

Goroutine 6 (running) created at:
  main.main()
      /Users/hank/workspace/mine/go-projects/go-learning/02-advanced/goroutine/datarace/race_demo.go:9 +0x
11d
==================
==================
WARNING: DATA RACE
Write at 0x00c00010e7d8 by goroutine 6:
  main.main.func1()
      /Users/hank/workspace/mine/go-projects/go-learning/02-advanced/goroutine/datarace/race_demo.go:10 +0
x5c

Previous write at 0x00c00010e7d8 by main goroutine:
  main.main()
      /Users/hank/workspace/mine/go-projects/go-learning/02-advanced/goroutine/datarace/race_demo.go:13 +0
x146

Goroutine 6 (running) created at:
  main.main()
      /Users/hank/workspace/mine/go-projects/go-learning/02-advanced/goroutine/datarace/race_demo.go:9 +0x
11d
==================
key = a, val = 1
Found 2 data race(s)
exit status 66
```

运行后输出如上的报告，可以看到 10、13 行两个 `goroutine` 中存在竞态。

# Options

可以通过 `GORACE` 环境变量配置选项，包括：
- log_path: 默认为 `stderr`, Race 检测器将其报告写入名为 log_path.pid 的文件。特殊值 `stdout` 和 `stderr` 会将报告分别写入标准输出和标准错误。
- exitcode （默认 66 ）：检测到竞态后的退出码。
- strip_path_prefix （默认 "" ）：从所有报告的文件路径中删除此前缀，以使报告更加简洁。
- history_size （默认 1 ）：每个 `goroutine` 的内存访问历史记录是 `32K * 2**history_size elements` 。增加此值可以避免报告中出现“failed to restore the stack” 错误，但代价是增加内存使用量。
- halt_on_error （默认 0 ）：控制程序在报告第一次数据竞态后是否退出。
- atexit_sleep_ms （默认 1000 ）：退出之前在主 goroutine 中休眠的毫秒数。

例如，运行时添加参数设置：
```
GORACE="log_path=./log strip_path_prefix=Goroutine" go run -race race_demo.go
```
运行后会在当前目录创建 `log.xxx` 的日志文件，并记录 race 报告，过滤掉 Goroutine 开头的行。

常见的数据竞态存在于循环变量中，如示例代码中的 counter.go:
```go
func race() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Print(i)
			wg.Done()
		}()
	}
	wg.Wait()
}
```
在 go1.20 及之前的版本中，goroutine中读取的i值通常都是5，所以程序输出 `55555`。go1.20之后的版本已经解决了这个问题。

go1.21之前一般通过给 goroutine 传递参数来复制循环变量解决该问题:
```go
func fixRace() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			fmt.Print(j) // Good. Read local copy of the loop counter.
			wg.Done()
		}(i) // copy var i to j
	}
	wg.Wait()
}
```

执行测试代码难以全面的检测竞态代码，而且只能在运行时检测到，可以在编译时添加 `-race` 选项，然后运行可执行程序来检测：

```
go build -race .
./count
```

竞争检测器需要启用 cgo，并且在非 Darwin 系统上需要安装 C 编译器。竞争检测器支持 linux/amd64 、 linux/ppc64le 、 linux/arm64 、 linux/s390x 、 freebsd/amd64 、 netbsd/amd64 、 darwin/amd64 、 darwin/arm64 和 windows/amd64 。

在 Windows 上，竞争检测器运行时对安装的 C 编译器版本敏感；从 Go 1.21 开始，使用 -race 构建程序需要一个包含版本 8 或更高版本的 mingw-w64 运行时库的 C 编译器。您可以通过使用参数 --print-file-name libsynchronization.a 调用 C 编译器来测试它。较新的兼容 C 编译器将打印该库的完整路径，而较旧的 C 编译器只会回显该参数。

竞争检测的成本因程序而异，但对于典型的程序，内存使用量可能会增加 5-10 倍，执行时间可能会增加 2-20 倍。

目前，竞争检测器为每个 defer 和 recover 语句分配额外的 8 个字节。这些额外的分配在 goroutine 退出之前不会被恢复。这意味着，如果您有一个长时间运行的 goroutine 定期发出 defer 和 recover 调用，则程序内存使用量可能会无限增长。这些内存分配不会显示在 runtime.ReadMemStats 或 runtime/pprof 的输出中。

