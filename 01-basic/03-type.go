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

		复合类型：
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

	boolType()

	// ==============
	// 整型
	// ==============

	intType()

	// ==============
	// 浮点型
	// ==============

	floatType()

	// ==============
	// 复数类型
	// ==============

	complexType()

	// ==============
	// 字符串
	// ==============

	stringType()

	// ==============
	// 字符类型
	// ==============

	noCharType()

	// ==============
	// 数组
	// ==============

	arrayType()

	// ==============
	// 数组切片
	// ==============

	sliceType()

	// ==============
	// map
	// ==============

	mapType()
}

func mapType() {
	/*
		Go支持map类型，map是一堆键值对的未排序集合。
		map通过make函数创建，通过delete函数删除元素
	*/

	fmt.Println("===== map ====")

	// 创建人的结构体
	type Person struct {
		Id   string // 身份证
		Name string // 姓名
	}

	// > 申明map
	var persons map[string]Person
	// > 创建map，第二个可选参数是map的存储能力
	persons = make(map[string]Person)
	persons = make(map[string]Person, 5) // 创建容量为5的map
	// > 插入数据
	persons["张三"] = Person{Id: "1", Name: "张三"}
	persons["李四"] = Person{Id: "2", Name: "李四"}
	persons["王五"] = Person{Id: "3", Name: "王五"}
	// > 查找，第二个参数为一个bool，表示是否查找到
	person1, ok := persons["张三"]
	fmt.Println(ok, person1) // ~: true {1 张三}
	person2, ok := persons["哈哈"]
	fmt.Println(ok, person2) // ~: false { }
	// > 遍历map
	fmt.Println(len(persons)) // len获取长度为3
	// 执行结果不同，可以看出map的无序性
	for key, person := range persons {
		fmt.Println(key, person)
	}
	// ~:
	// 张三 {1 张三}
	// 李四 {2 李四}
	// 王五 {3 王五}
	// > 删除元素
	delete(persons, "张三")
	for key, person := range persons {
		fmt.Println(key, person)
	}
	// ~:
	// 王五 {3 王五}
	// 李四 {2 李四}
}

func sliceType() {
	/*
		Go中数组一旦定义，长度不可改变，不能完全满足开发需求，因此提供数组切片类型(Slice)，它基于数组创建，并且长度可以改变
		切片底层依赖一个数组，因此切片有长度和容量的概念，切片的容量就是底层数组的长度，而切片的长度是其存储元素的个数
	*/

	fmt.Println("===== slice =====")
	// > 基于数组创建切片
	// 使用array[开始下标:结束下标]的格式创建切片，注意不会包含结束下标
	// 如果是基于整个数组创建切片，可以写为[:]，开始下标0可以省略，写为[:结束下标]，如果结束下标省略，则为开始下标到之后的所有数组元素
	var array = [5]int{1, 2, 3, 4, 5} // 定义数组
	var slice = array[0:3]            // 以数组第1个元素到第3个元素(不包括下标为3的元素)创建切片
	fmt.Println(slice)                // ~: [1 2 3]
	slice1 := array[:]                // 基于所有数组元素创建切片
	fmt.Println(slice1)               // ~: [1 2 3 4 5]
	slice2 := array[:3]               // 省略开始下标，则相当于[0:3]
	fmt.Println(slice2)               // ~: [1 2 3]
	slice3 := array[3:]               // 省略结束下标，则按后续所有数组元素创建切片
	fmt.Println(slice3)               // ~: [4 5]

	// > 遍历切片
	for i := 0; i < len(slice1); i++ { // 使用len()获取切片长度
		fmt.Print(slice1[i], " ")
	}
	// ~: 1 2 3 4 5
	fmt.Println()
	for i, s := range slice1 { // 使用range
		fmt.Print("[", i, "]=", s, " ")
	}
	// ~: [0]=1 [1]=2 [2]=3 [3]=4 [4]=5
	fmt.Println()

	// > 直接创建切片，不依赖数组
	slice4 := make([]int, 5)     // 使用make方法创建slice，第二个参数为长度
	fmt.Println(slice4)          // ~: [0 0 0 0 0]
	slice5 := make([]int, 5, 10) // 创建一个int类型的切片，初始长度为5，容量为10（预留10个元素的存储空间）
	// 使用len方法获取切片长度，cap方法获取切片的容量
	fmt.Println("len(slice5)=", len(slice5), "cap(slice5)=", cap(slice5), "slice5: ", slice5)
	// ~: len(slice5)= 5 cap(slice5)= 10 slice5:  [0 0 0 0 0]

	// > 动态添加元素
	// 切片应该按照业务需求设置合理的容量，以避免底层数组的重新内存分配和移动，影响程序性能
	// 使用append方法可以给切片动态添加元素, 返回一个新切片类型，而不影响原切片
	slice6 := []int{0, 0}
	slice7 := append(slice6, 2)                   // 添加单个元素
	fmt.Println(slice6)                           // ~: [0 0]
	fmt.Println(slice7)                           // ~: [0 0 2]
	slice8 := append(slice7, 3, 5, 7)             // 添加多个元素
	fmt.Println(slice8)                           // ~: [0 0 2 3 5 7]
	slice9 := append(slice8, slice6...)           // 向切片中添加切片
	fmt.Println(slice9)                           // ~: [0 0 2 3 5 7 0 0]
	slice10 := append(slice9[3:6], slice8[2:]...) // 添加动态切片
	fmt.Println(slice10)                          // ~: [3 5 7 2 3 5 7]
	// 任意类型的切片
	var slice11 []interface{}
	slice11 = append(slice11, 42, 3.1415, "foo") // 将数组切片添加不同的类型
	fmt.Println(slice11)                         // ~: [42 3.1415 foo]
	var slice12 []byte
	slice12 = append(slice12, "bar"...)
	fmt.Println(slice12) // ~: [98 97 114]
	// > 尽量保持切片的容量与长度相同，这样在第一次调用append方法时会拷贝底层数组到一个新数组，
	// 避免对原底层数组的修改而影响其他基于该数组的切片
	fmt.Println("切片的append修改原始数组，造成依赖原始数组的切片数据都变化")
	var srcArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	intSlice := srcArray[2:4]
	intSlice1 := srcArray[3:5]
	fmt.Println("Before Append:")
	fmt.Printf("  src array: %v\n", srcArray) // src array: [1 2 3 4 5 6 7 8 9 10]
	fmt.Printf("  slice: %v, len: %d, cap: %d \n", intSlice,
		len(intSlice), cap(intSlice)) // slice: [3 4], len: 2, cap: 8
	fmt.Printf("  another slice: %v \n", intSlice1) // [4 5]
	intSlice = append(intSlice, 11)
	fmt.Println("After Append:")
	fmt.Printf("  src array: %v\n", srcArray) // src array: [1 2 3 4 11 6 7 8 9 10]
	fmt.Printf("  slice: %v, len: %d, cap: %d \n", intSlice,
		len(intSlice), cap(intSlice)) //   slice: [3 4 11], len: 3, cap: 8
	fmt.Printf("  another slice: %v \n", intSlice1) // 由于intSlice append后修改了底层数组，导致intSlice1数据变化： [4 11]
	// 可以看到append修改了原数组的数据，所有基于原数组的切片都可能变化
	// 现在，创建容量和长度相同的切片
	fmt.Println("指定长度和容量相同的优势：")
	srcArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 通过第三个索引值限制容量，长度 = 4 - 2 = 2，容量 = 4 - 2 = 2，长度和容量相同
	intSlice = srcArray[2:4:4]
	intSlice1 = srcArray[3:5]
	fmt.Println("Before Append:")
	fmt.Printf("  src array: %v\n", srcArray) // src array: [1 2 3 4 5 6 7 8 9 10]
	fmt.Printf("  slice: %v, len: %d, cap: %d \n", intSlice,
		len(intSlice), cap(intSlice)) // slice: [3 4], len: 2, cap: 2
	fmt.Printf("  another slice: %v \n", intSlice1) // [4 5]
	intSlice = append(intSlice, 11)
	fmt.Println("After Append:")
	fmt.Printf("  src array: %v\n", srcArray) //重新创建了一个数组，原数组不变： src array: [1 2 3 4 5 6 7 8 9 10]
	fmt.Printf("  slice: %v, len: %d, cap: %d \n", intSlice,
		len(intSlice), cap(intSlice)) // 切片元素少于1000，容量增长两倍，大于1000时，1.25被增长：slice: [3 4 11], len: 3, cap: 4
	fmt.Printf("  another slice: %v \n", intSlice1) // 现在intSlice1数据无变化： [4 5]

	// > 切片拷贝
	// 使用copy方法将原切片src的元素拷贝到目标切片dst中，返回拷贝的元素个数。拷贝的数组切片必须属于同一类型，并且如果两个切片长度不一致，则按照
	// 两者最小长度来拷贝
	// copy方法的定义如下：
	// copy(dst, src []T) int
	// copy(dst []byte, src string) int
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7} // 数组
	var s = make([]int, 6)
	n1 := copy(s, a[0:])     // 将a[0:1]切片拷贝到s中
	fmt.Println(n1, s)       // ~: 6 [0 1 2 3 4 5]
	n2 := copy(s, s[2:])     // s[2:] 为 [2 3 4 5]，按照最短的拷贝，所以占了前边4个位置，后边两个元素不变
	fmt.Println(n2, s)       // ~: 4 [2 3 4 5 4 5]
	var b = make([]byte, 5)  // 创建5个元素长度的byte类型切片
	n3 := copy(b, "abcdefg") // 将string拷贝到b中，string会传为uin8的字符型
	fmt.Println(n3, b)       // ~: 5 [97 98 99 100 101]

	//
}

func arrayType() {
	fmt.Println("===== array =====")
	// 数组的申明
	const N = 2
	var a1 [3]byte                     // 长度为32的数组，每个元素为一个字节
	a1 = [3]byte{2, 2, 2}              // 给数组赋值
	var a2 [2 * N]struct{ x, y int32 } // 复杂类型数组
	var a3 [2]*float64                 // 指针数组
	var a4 [3][2]int                   // 二维数组
	var a5 [2][2][2]float64            // 多维数组，等同于[2]([2]([2]float64))

	fmt.Println(a1) // ~: [2 2 2]
	fmt.Println(a2) // ~: [{0 0} {0 0} {0 0} {0 0}]
	fmt.Println(a3) // ~: [<nil> <nil>]
	fmt.Println(a4) // ~: [[0 0] [0 0] [0 0]]
	// 可使用len方法获取数组长度
	fmt.Println(a5, len(a5)) // ~: [[[0 0] [0 0]] [[0 0] [0 0]]] 2

	// 数组一旦定义，长度不可更改
	a6 := [2]int{2, 1}
	fmt.Println(a6)

	// 元素访问
	for i := 0; i < len(a6); i++ { // 循环遍历
		fmt.Print(a6[i], " ") // 使用下标访问
	}
	// ~: 2 1
	fmt.Println()
	for i, a := range a6 { // 使用range遍历
		fmt.Println("i=", i, ", v=", a)
	}
	// ~:
	// i=0 , v=2
	// i=1 , v=1

	// go中数组的传递是值类型，作为参数传递时，函数操作的仅仅是数组的复制副本
	a7 := [2]int{1, 2}
	fmt.Println("before modify: ", a7) // ~; array in main:  [1 2]
	modifyArray(a7)                    // ~: array in modifyArray:  [1 10]
	fmt.Println("after modify: ", a7)  // ~; array in main:  [1 2]

	// 用 ... 可用替代数组长度，go根据元素来自动确定长度
	array1 := [...]int{10, 20, 30, 40}
	fmt.Println(array1)
	// 通过 下标:元素 的形式可以指定数组内某些下标的初始元素，其他元素为默认值
	array2 := [...]int{1: 10, 2: 20}
	fmt.Println(array2) // [0 10 20]
	array3 := [5]int{1: 10, 2: 20}
	fmt.Println(array3) // [0 10 20 0 0]
}

func noCharType() {
	fmt.Println("===== 没有字符类型，用 rune 和 byte =====")
	// Go的字符类型有两种：
	// byte：uint8的别名，代表UTF-8字符串的单个字节的值
	// rune：代表单个unicode字符
	// 见前边字符串的例子
}

func stringType() {
	fmt.Println("===== string =====")
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
}

func complexType() {
	fmt.Println("===== complex =====")
	var cv1 complex64 // 由2个float32构成的复数类型
	cv1 = 3.2 + 12i
	cv2 := 3.2 + 12i           // value2是complex128类型
	cv3 := complex(3.2, 12)    // value3结果同 value2
	fmt.Println(cv1, cv2, cv3) // ~: (3.2+12i) (3.2+12i) (3.2+12i)
	cv4 := real(cv3)           // 获取实数的实部
	cv5 := imag(cv3)           // 获取实数的虚部
	fmt.Println(cv4, cv5)      // ~: 3.2 12
}

func floatType() {
	fmt.Println("===== float =====")
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
}

func intType() {
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

	fmt.Println("===== int =====")
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
}

func boolType() {
	fmt.Println("===== bool =====")
	// bool类型值为true和false，不接受其他值，不支持类型转换
	var vb1 bool
	vb1 = true
	vb2 := false
	var vb3 bool // 默认值为false
	// vb3 = 1 // 编译错误：cannot use 1 (type untyped int) as type bool in assignment
	// vb4 := bool(1) // 编译错误：cannot convert 1 (type untyped int) to type bool
	vb5 := 1 == 0
	fmt.Println(vb1, vb2, vb3, vb5) // ~: true false false false
}

func IsEqual(f1, f2, p float64) bool {
	return math.Dim(f1, f2) < p
}

func printSlice() {

}
func modifyArray(array [2]int) {
	array[1] = 10
	fmt.Println("array in modifyArray: ", array)
}
