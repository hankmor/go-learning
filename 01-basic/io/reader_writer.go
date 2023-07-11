package main

import (
	"bytes"
	"fmt"
	"os"
)

func ReadAndWrite() {
	// 创建buffer
	var b bytes.Buffer
	// 写入字节
	b.Write([]byte("hello, "))

	// 将另一个字符串连接到 buffer 中
	_, err := fmt.Fprintf(&b, "Go!\n")
	if err != nil {
		panic(err)
	}

	// 输出到控制台
	_, err = b.WriteTo(os.Stdout)
	if err != nil {
		panic(err)
	}
}

func main() {
	ReadAndWrite()
}
