package main

import (
	"fmt"
)

func main() {
	// ==============
	// 常量声明
	// ==============

	// 不同于变量，常量声明过后，不被使用也不会编译错误

	const Pi float64 = 3.14159265358979323846 // float64类型的敞亮
	const zero = 0.0                          // 无类型浮点常量，可推导类型
	fmt.Println("Pi=", Pi, "zero=", zero)     // ~: Pi= 3.141592653589793 zero= 0

	// 同变量一样，支持同时定义多个常量
	const (
		size int64 = 1024 // 有类型整型常量
		eof        = -1   // 无类型整型常量，可推导类型
	)
	fmt.Println("size=", size, "eof=", eof) // ~: size= 1024 eof= -1

	const u, v float32 = 0, 3                                // 常量的多重赋值
	const a, b, c = 3, 4, "foo"                              // 无类型整型和字符串常量
	fmt.Println("u=", u, "v=", v, "a=", a, "b=", b, "c=", c) // ~: u= 0 v= 3 a= 3 b= 4 c= foo

	// 常量也可以定义为编译器可运算的表达式，但是需要运行时才能确定的表达式不能定义为常量
	const d = 1 << 3
	fmt.Println("d=", d)
	// const e = os.Getenv("GOROOT") // os.Getenv运行时才能确定其值，编译错误：const initializer os.Getenv("GOROOT") is not a constant

	// ==============
	// 预定义常量
	// ==============

	// go预定义了几个常量：true、false、iota
	// iota用在使用括号()同时定义多个常量时
	// iota表示可被编译器修改的常量
	// iota在const关键字出现时置为0，在下一次const出现之前没出现一次自动增加1
	const ( // iota被重置为0
		c0 = iota // iota为0
		c1 = iota // 1, iota自动增加1
		c2 = iota // 2, iota自动增加1
	)
	const i = iota                                        // iota被重置为0
	fmt.Println("c0=", c0, "c1=", c1, "c2=", c2, "i=", i) // ~: c0= 0 c1= 1 c2= 2 i= 0
	// iota用于运算
	const ( // iota被重置为0
		c3 = 1 << iota // iota为0
		c4 = 1 << iota // iota为1
		c5 = 1 << iota // iota为2
	)
	fmt.Println("c3=", c3, "c4=", c4, "c5=", c5) // ~: c3= 1 c4= 2 c5= 4

	// 这种一次定义两个常量，相同于使用了两次const关键字，所以iota都为0
	const c6, c7 = 2 * iota, 10.0 * iota
	fmt.Println("c6=", c6, "c7=", c7) // ~; c6= 0 c7= 0
	const c8 = iota                   // iota又被重置为0
	fmt.Println("c8=", c8)            // ~; c8= 0

	// 如果多个const赋值语句表达式是一样的，可以省略后一个iota关键字
	const (
		c9  = iota
		c10 // 省略赋值语句
		c11
	)
	fmt.Println("c9=", c9, "c10=", c10, "c11=", c11) // ~: c9= 0 c10= 1 c11= 2
	const (                                          // iota被重置为0
		c12 = 1 << iota
		c13 // 省略赋值语句
		c14
	)
	fmt.Println("c12=", c12, "c13=", c13, "c14=", c14) // ~: c12= 1 c13= 2 c14= 4

	// ==============
	// 枚举
	// ==============

	// go语言没有enum关键字，可以通过const圆括号的方式定义枚举值
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays = 7 // 小写开头的变量没有导出，其他包看不到
	)
	fmt.Println("Sunday=", Sunday, "Monday=", Monday, "Tuesday=", Tuesday, "Wednesday=", Wednesday, "Thursday=", Thursday, "Friday=", Friday, "Saturday=", Saturday, "numberOfDays=", numberOfDays)
	// ~: Sunday= 0 Monday= 1 Tuesday= 2 Wednesday= 3 Thursday= 4 Friday= 5 Saturday= 6 numberOfDays= 7
}
