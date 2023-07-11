package main

import "fmt"

// 包级别的变量初始化顺序:
// 1) 按照变量申明的先后顺序依次初始化
// 2) 如果变量 a 依赖变量 b，则 b 初始化先于 a
// 3) 如果一个未初始化的变量 a 不依赖于任何其他变量，则表明这个变量"ready for initialization"
// 4) 初始化时按照变量申明的先后顺序依次寻找 "ready for initialization" 的变量，不断重复，直到所有变量初始化完成
// 5) 位于同一包内但不同文件中的变量初始化依赖于编译器处理文件的顺序，先编译的文件中的变量申明顺序优先且先初始化

var (
	a = c + b // a 依赖 c, b，故先初始化 c, b，由于 b 先申明，所以 b 初始化先于 c
	b = fn()  // b 先初始化，得到 2
	c = fn()  // c 再初始化，得到 3，现在可以初始化 a 了，得到 2 + 3 = 5
	d = 1     // 最终 d 变为 3
)

func fn() int {
	d++
	return d
}

func main() {
	fmt.Println(a, b, c, d)
	// output: 5 2 3 3
}
