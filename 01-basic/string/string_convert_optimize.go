package main

import (
	"fmt"
	"testing"
)

func convert() {
	s := "中国欢迎您！"
	s1 := []byte(s) // 先转换再使用
	for _, v := range s1 {
		_ = v
	}
}

func convertWithOptimize() {
	s := "中国欢迎您！"
	for _, v := range []byte(s) { // range 循环时转换
		_ = v
	}
}

func main() {
	fmt.Println(testing.AllocsPerRun(1, convert)) // 运行 1 次的内存分配次数
	fmt.Println(testing.AllocsPerRun(1, convertWithOptimize))
}

/*output:
0
0

可以看到，转为 []byte 后其实使用的是 string 的底层数据，所以不会进行内存分配。
*/
