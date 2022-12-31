package main

import (
	"fmt"
	"strings"
	"time"
)

// 换行符: \n，换一行, code: 10
// 回车符：\r, 回到行的开头, code: 13

func main() {
	fmt.Printf("%x\n", "\b") // 08，退格符
	fmt.Printf("%x\n", "\t") // 09，水平指标符
	fmt.Printf("%x\n", "\n") // 0a，换行符
	fmt.Printf("%x\n", "\v") // 0b，垂直制表符
	fmt.Printf("%x\n", "\f") // 0c，换页符
	fmt.Printf("%x\n", "\r") // 0d，回车符
	const col = 30
	// Clear the screen by printing \x0c.
	// Note: \x0c not worked in macOs, use \x0d instead, e.g: \r.
	bar := fmt.Sprintf("\x0d[%%-%vs]", col)
	for i := 0; i < col; i++ {
		fmt.Printf(bar, strings.Repeat("=", i)+">")
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf(bar+" Done!", strings.Repeat("=", col))
}
