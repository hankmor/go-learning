package basic

import (
	"fmt"
)

// 变量
func VarMain() {
	// ==============
	// 变量声明
	// ==============

	varDeclaration()

	// ==============
	// 变量初始化
	// ==============

	varInit()

	// ==============
	// 变量赋值
	// ==============

	assignValue()
}

func varDeclaration() {
	// 变量声明过后，必须被使用，否则编译错误

	// 变量声明：var 变量名 变量类型
	var v0 rune   // 字符
	var v1 int    // 整型
	var v2 string // 字符串
	// 变量赋值
	v0 = 'A'
	fmt.Println("v0=", v0, "v1=", v1, "v2=", v2) // ~: v0= 65 v1= 0 v2=

	// 可以将多个变量合并声明
	var (
		v3 [2]int          // 数组
		v4 []int           // 数组切片
		v5 struct{ i int } // 结构体
		v6 *int            // 指针
		v7 map[string]int  // map
		// v8 func(a int, b int) int // 匿名函数
	)
	fmt.Println("v3=", v3, "v4=", v4, "v5=", v5, "v6=", v6, "v7=", v7, "v8=")
	// ~: v3=[0 0] v4=[] v5={0} v6=<nil> v7=map[] v8=<nil>
}

func varInit() {
	var v9 int = 10 // 声明变量时初始化赋值
	var v10 = 10    // 自动推导类型
	v11 := 10       // 简写方式，省略var关键字，且可推导类型
	v12 := "abc"    // 声明一个string变量
	// v9 := 100       // v9变量已经声明过了，编译出错：no new variables on left side of :=
	fmt.Println("v9=", v9, "v10=", v10, "v11=", v11, "v12=", v12)
	// ~: v9= 10 v10= 10 v11= 10 v12= abc
}

func assignValue() {
	var v13 int         // 声明变量
	v13 = 10            // 变量赋值
	v14 := 11           // 声明并赋值
	v13, v14 = v14, v13 // 使用多重赋值交换两个变量的值，不需要引入中间变量
	fmt.Println("v13=", v13, "v14=", v14)
	// ~: v13= 11 v14= 10

	// ==============
	// 匿名变量
	// ==============

	f, l, n := getName()                                         // 调用函数，返回多个值
	fmt.Println("firstName=", f, "lastName=", l, "nickName=", n) // ~; firstName= sam lastName= sune nickName= belonk
	// 如果不需要某些值，可以使用匿名变量将其忽略
	_, _, nickName := getName()
	fmt.Println("nickName=", nickName) // ~: nickName= belonk
}

func getName() (firstName, lastName, nickName string) {
	return "sam", "sune", "belonk"
}
