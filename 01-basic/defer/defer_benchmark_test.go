package main

import "testing"

func sum(max int) int {
	var t int
	for i := 0; i < max; i++ {
		t += i
	}
	return t
}

func sumWithDefer() {
	defer func() {
		sum(100)
	}()
}

func sumWithoutDefer() {
	sum(100)
}

func BenchmarkWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumWithDefer()
	}
}

func BenchmarkWithoutDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumWithoutDefer()
	}
}

/*
➜  defer git:(main) ✗ go test -bench . defer_benchmark_test.go
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkWithDefer-8            31565316                37.02 ns/op
BenchmarkWithoutDefer-8         32567470                35.01 ns/op
PASS
ok      command-line-arguments  3.655s

结论：defer 的性能与不使用 defer 比稍微差一点，但是已经优化的很好了，性能差别不大
*/
