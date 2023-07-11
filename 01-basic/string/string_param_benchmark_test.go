package main

import "testing"

var s = `Go started in September 2007 when Robert Griesemer, Ken Thompson, and I began discussing a new language to 
address the engineering challenges we and our colleagues at Google were facing in our daily work. When we first released 
Go to the public in November 2009, we didn’t know if the language would be widely adopted or if it might influence future 
languages. Looking back from 2020, Go has succeeded in both ways: it is widely used both inside and outside Google, 
and its approaches to network concurrency and software engineering have had a noticeable effect on other languages and 
their tools. Go has turned out to have a much broader reach than we had ever expected. Its growth in the industry has 
been phenomenal, and it has powered many projects at Google. — Rob Pike`

func handleString(s string) string {
	return s + " hello, go!"
}

func handleStringPtr(s *string) string {
	return *s + " hello, go!"
}

func BenchmarkHandleString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		handleString(s)
	}
}

func BenchmarkHandleStringPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		handleStringPtr(&s)
	}
}

/*
➜  string git:(main) ✗ go test -bench . -benchmem string_param_benchmark_test.go
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkHandleString-8          8070846               129.1 ns/op           896 B/op          1 allocs/op
BenchmarkHandleStringPtr-8       9164533               133.2 ns/op           896 B/op          1 allocs/op
PASS
ok      command-line-arguments  3.382s

可以看出：传递字符串参数与传递字符串指针参数的性能差别不大，其实字符串底层也是直接指向了存储空间上。
*/
