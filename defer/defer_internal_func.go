package main

// defer 后执行内置函数测试
// go 内置函数：append make new len cap panic recover delete copy close print println real complex imag

func f4() (int, int) {
	return 1, 2
}

func f5() {
	var s = []int{1, 2, 3}
	var m = make(map[string]int, 10)
	var c = make(chan int)
	var n = complex(3.14, 2.09)
	_ = n

	m["a"] = 1
	m["b"] = 2
	var o []int

	defer f4()
	defer close(c)
	defer println("println")
	defer print("print")
	defer copy(o, s)
	defer delete(m, "a")
	defer recover()
	defer panic(123)
	// 以下内置函数在 defer 后使用编译失败
	// defer cap(s)
	// defer append(s, 4)
	// defer complex(3.14, 2.09)
	// defer new(*int)
	// defer make([]int, 10)
	// defer real(n)
	// defer imag(c) // 返回复数的虚数部分
}

func main() {
	f5()
}
