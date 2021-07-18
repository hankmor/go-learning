package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		基础类型：
		1、布尔类型：bool
		2、整型：int8，byte(uint8的别名)，int16，int32，int64，uint8，uint16，uint32，uint64
				与操作系统有关的32位或64位：int，uint
				uintptr：一个足够大的无符号整数，可以存储指针值的未解释位
		3、浮点类型：float32，float64
		4、复数类型：complex64，complex128
		5、字符串：string
		6、字符类型：rune(int32的别名)
		7、错误类型：error

		符合类型：
		1、数组：array
		2、切片：slice
		3、结构体：struct
		4、指针：pointer
		5、接口：interface
		6、字典：map
		7、通道：channel
	*/

	// ==============
	// 布尔
	// ==============

	// bool类型值为true和false，不接受其他值，不支持类型转换
	var vb1 bool
	vb1 = true
	vb2 := false
	var vb3 bool // 默认值为false
	// vb3 = 1 // 编译错误：cannot use 1 (type untyped int) as type bool in assignment
	// vb4 := bool(1) // 编译错误：cannot convert 1 (type untyped int) to type bool
	vb5 := 1 == 0
	fmt.Println(vb1, vb2, vb3, vb5) // ~: true false false false

	// ==============
	// 整型
	// ==============

	// 整型长度和值范围
	/*
		类型							长度（字节）			值范围
		int8						1					-128~127
		uint8(即byte)				1					0~255
		int16						2					-32768~32767
		uint16						2					0~65535
		int32						4					-2147483648~2147483647
		uint32						4					0~4294967295
		int64						8					-9223372036854775808~9223372036854775807
		uint64						8					0~18446744073709551615
		int							平台相关				平台相关
		uint						平台相关				平台相关
		uintptr						同指针				在32位平台下为4字节，64位平台下为8字节
	*/

	// int int32是两种类型，不能直接赋值
	var i1 int32
	i2 := 20
	// i1 = i2 // 编译错误：cannot use i2 (type int) as type int32 in assignment
	i1 = int32(i2)      // 可以通过强制类型转换赋值
	fmt.Println(i1, i2) // ~: 20 20

	// 整形比较，支持 >、<、==、>=、<=和!=
	var i3 int32
	var i4 int
	i3, i4 = 1, 2
	// compare := i3 == i4 // 不同整型类型之间不能直接比较，编译错误：invalid operation: i3 == i4 (mismatched types int32 and int)
	compare := i3 == 1 && i4 == 2
	fmt.Println(compare) // ~: true

	// 整型运算，支持 + - * / %
	i5 := 10
	i6 := 3
	fmt.Println(i5 / i6) // ~: 3
	fmt.Println(i5 % i6) // ~: 1

	// 整型位运算，支持 与(&) 或(|) 异或(^) 左移(<<) 右移(>>) 取反(^，大多数语言取反是~)
	// 注意：go语言没有无符号移动(>>> <<<)
	i7 := 2
	i8 := 3
	fmt.Println(i7&i8, i7|i8, i7^i8, i7<<i8, i7>>i8, ^i7) // ~: 2 3 1 16 0 -3

	// ==============
	// 浮点型
	// ==============

	var fv1 float32
	fv1 = 12
	fv2 := 12.0 // 被自动推导为float64类型，如果不加小数点，fvalue2会被推导为整型而不是浮点型
	// fv1 = fv2 // 编译错误：cannot use fv2 (type float64) as type float32 in assignment
	fv1 = float32(fv2)    // 强制类型转换可以赋值
	fmt.Println(fv1, fv2) // ~: 12 12

	feq := fv1 == float32(fv2)
	fmt.Println(feq) // ~: true
	fv3 := 12.156000000000001
	fv4 := 12.1560000000000001
	fmt.Println(fv3 == fv4) // 这里的结果为true，是不准确的
	// fmt.Println(math.Dim(fv3, fv4))
	// feq = IsEqual(fv3, fv4, 0.00000000000001)
	// fmt.Println(feq)

	// ==============
	// 复数类型
	// ==============

	var cv1 complex64 // 由2个float32构成的复数类型
	cv1 = 3.2 + 12i
	cv2 := 3.2 + 12i           // value2是complex128类型
	cv3 := complex(3.2, 12)    // value3结果同 value2
	fmt.Println(cv1, cv2, cv3) // ~: (3.2+12i) (3.2+12i) (3.2+12i)
	cv4 := real(cv3)           // 获取实数的实部
	cv5 := imag(cv3)           // 获取实数的虚部
	fmt.Println(cv4, cv5)      // ~: 3.2 12

	// ==============
	// 字符串
	// ==============

	str := "Hello go"
	c := str[0]
	fmt.Printf("The length of %s is %d.\n", str, len(str)) // 使用len函数取字符串长度
	fmt.Printf("The first character of %s is %c.\n", str, c)
	// ~: The length of Hello go is 8.
	// ~: The first character of Hello go is H.

	// 字符串声明后，其中的内容不能修改
	// str[0] = "X" // 编译错误：cannot assign to str[0] (strings are immutable)

	// 字符串拼接
	fmt.Printf("Hello" + " " + "World!\n") // ~: Hello World!
	// 其他类型不能直接与字符串相加，需要转为字符串
	// fmt.Printf(100 + "Hello" + " " + "World!\n") // 编译错误: invalid operation: 100 + "Hello World!\n" (mismatched types untyped int and untyped string)
	fmt.Printf(string(100) + "Hello" + " " + "World!\n") // ~: dHello World!
	fmt.Printf("100" + "Hello" + " " + "World!\n")       // ~: 100Hello World!

	// 字符串遍历
	// go语言默认编码为UTF-8，每个中文占3个字节
	hw := "Hello,世界"
	for i := 0; i < len(hw); i++ {
		c := hw[i] // c的类型为: byte(uint8)
		fmt.Println(i, c)
	}
	// 上边存在中文字符，UTF8中每个中文占3个字节，所以一共输出12个字符，每个字符的类型为byte(uint8的别名)

	// 要正确打印中文等unicode字符，有两种方法
	// 1、使用range遍历
	for i, c := range hw { // c的类型为：rune(int32)
		fmt.Printf("i=%d, c=%c.\n", i, c) // 正确输出中文
	}
	// 上边会正确输出10个字符，遍历c类型为rune(int32的别名)，而不是byte
	// 2、使用rune slice，这种方法的下标i会跳跃，因为中文占3个字符
	hws := []rune(hw)
	fmt.Println(len(hws))
	for i := 0; i < len(hws); i++ {
		c := hws[i]                       // c的类型为：rune(int32)
		fmt.Printf("i=%d, c=%c.\n", i, c) // 正确输出中文
	}

	// ==============
	// 字符类型
	// ==============

	// ==============
	// 数组
	// ==============

	// ==============
	// 数组切片
	// ==============

	// ==============
	// map
	// ==============
}

func IsEqual(f1, f2, p float64) bool {
	return math.Dim(f1, f2) < p
}
