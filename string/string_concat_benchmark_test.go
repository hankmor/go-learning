package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// 构造字符串的几种方式：
// 1) 直接通过 "+" 连接
// 2) 使用 fmt.Sprintf
// 3) 使用 strings.Join
// 4) 使用 strings.Builder
// 5) 使用 bytes.Buffer

var ss = strings.Split(`Go started in September 2007 when Robert Griesemer, Ken Thompson, and Rob Pike began discussing a new language to
address the engineering challenges we and our colleagues at Google were facing in our daily work`, " ")

func concatStringUsePlus(p []string) string {
	var s string
	for _, v := range p {
		s += v + " "
	}
	return s
}

func concatStringUseFmt(p []string) string {
	var s string
	for _, v := range p {
		s = fmt.Sprintf("%s %s", s, v)
	}
	return s
}

func concatStringUseStringsJoin(p []string) string {
	return strings.Join(p, " ")
}

func concatStringUseStringsBuilder(p []string) string {
	var b strings.Builder // 未设置初始容量，则默认为 0
	for _, v := range p {
		b.WriteString(v)
		b.WriteString(" ")
	}
	return b.String()
}

func concatStringUseStringsBuilderWithInitSize(p []string) string {
	var b strings.Builder
	b.Grow(len(p)) // 给 builder 设置初始容量
	for _, v := range p {
		b.WriteString(v)
		b.WriteString(" ")
	}
	return b.String()
}

func concatStringUseByteBuffer(p []string) string {
	var b bytes.Buffer
	for _, v := range p {
		b.WriteString(v)
		b.WriteString(" ")
	}
	return b.String()
}

func concatStringUseByteBufferWithInitSize(p []string) string {
	var b = make([]byte, 0, 128)
	var bf = bytes.NewBuffer(b)
	for _, v := range p {
		bf.WriteString(v)
		bf.WriteString(" ")
	}
	return bf.String()
}

func main() {
	fmt.Println(concatStringUsePlus(ss))
	fmt.Println(concatStringUseFmt(ss))
	fmt.Println(concatStringUseStringsJoin(ss))
	fmt.Println(concatStringUseStringsBuilder(ss))
	fmt.Println(concatStringUseStringsBuilderWithInitSize(ss))
	fmt.Println(concatStringUseByteBuffer(ss))
	fmt.Println(concatStringUseByteBufferWithInitSize(ss))
}

func BenchmarkConcatStringUsePlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringUsePlus(ss)
	}
}
func BenchmarkConcatStringUseFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringUseFmt(ss)
	}
}
func BenchmarkConcatStringUseStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringUseStringsJoin(ss)
	}
}
func BenchmarkConcatStringUseStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringUseStringsBuilder(ss)
	}
}
func BenchmarkConcatStringUseStringsBuilderWithInitSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringUseStringsBuilderWithInitSize(ss)
	}
}
func BenchmarkConcatStringUseByteBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringUseByteBuffer(ss)
	}
}
func BenchmarkConcatStringUseByteBufferWithInitSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringUseByteBufferWithInitSize(ss)
	}
}

/*output:
➜  string git:(main) ✗ go test -bench . -benchmem string_concat_benchmark_test.go
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkConcatStringUsePlus-8                            540818              1900 ns/op            3992 B/op         34 allocs/op
BenchmarkConcatStringUseFmt-8                             186873              5915 ns/op            5065 B/op        101 allocs/op
BenchmarkConcatStringUseStringsJoin-8                    3633477               317.6 ns/op           224 B/op          1 allocs/op
BenchmarkConcatStringUseStringsBuilder-8                 2946358               391.4 ns/op           504 B/op          6 allocs/op
BenchmarkConcatStringUseStringsBuilderWithInitSize-8     3448684               341.8 ns/op           608 B/op          4 allocs/op
BenchmarkConcatStringUseByteBuffer-8                     2076072               583.3 ns/op           720 B/op          4 allocs/op
BenchmarkConcatStringUseByteBufferWithInitSize-8         2350407               487.7 ns/op           512 B/op          2 allocs/op
PASS
ok      command-line-arguments  11.017s
*/
