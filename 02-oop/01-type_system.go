package main

import "fmt"

// go语言的类型系统
// 空接口 interface {} 是go中的Any 类型，类似 Java 中的 Object
// 可以给任何类型（除了指针）添加方法
//
// 与java不同，在Go语言中没有隐藏的this指针，即：
// 1、方法的目标显示传递，通过 func 后边跟 (类型) 来附加，这也是与函数不同的地方
// 2、方法施加的目标（也就是“对象”）不需要非得是指针，也不用非得叫this
func main() {
	// interface具体类型断言
	var i interface{}
	i = 10
	// r := i.(int) // 断言
	// println(r)
	i = "abc"
	r := i.(string)
	println(r)

	fmt.Println("===== 获取类型 =====")
	printType(1)
	printType("abc")
	printType(3.1415)
	printType([10]int{})
	printType([]int{})
	var m map[string]int
	printType(m)
	type t struct {
	}
	var t1 t
	printType(t1)
	var t2 = func() {}
	printType(t2)

	fmt.Println("===== 自定义类型 =====")
	var a Integer = 1
	var b Integer = 10
	fmt.Println(a.less(b))  // true
	fmt.Println(less(a, b)) // true

	fmt.Println("值传递")
	fmt.Println("before add: ", a) // 1
	a.add(b)                       // 值传递，不会改变原来a的值，实际上是方法将a的值重新拷贝了一份
	fmt.Println("after add: ", a)  // 1

	fmt.Println("引用(指针)传递")
	fmt.Println("before add: ", a) // 1
	a.addRef(b)                    // 指针传递，改变了原来a的值，这里go将a解引用为 *a，相当于 (&a).addRef(b)
	fmt.Println("after add: ", a)  // 11

	fmt.Println("===== 数组的值传递 =====")
	arrayVal()
	fmt.Println("===== 数组的引用传递 =====")
	arrayRef()
}

func printType(t any) {
	switch v := t.(type) {
	case int:
		fmt.Printf("type: %T, value: %v\n", v, v)
	case string:
		fmt.Printf("type: %T, value: %s\n", v, v)
	case float64:
		fmt.Printf("type: %T, value: %v \n", v, v)
	default:
		fmt.Printf("type: %T, value: %v \n", v, v)
	}
}

func arrayVal() {
	var a = [3]int{1, 2, 3}
	var b = a         // 将 a 赋值给b，此时会完全拷贝一个数组给b
	b[1]++            // 将b数组的第2个元素+1
	fmt.Println(a, b) // [1 2 3] [1 3 3]
}

func arrayRef() {
	var a = [3]int{1, 2, 3}
	var b = &a         // 将 a 的指针地址赋值给 b，不会拷贝数组
	b[1]++             // go自动解引用为 (*b)[1]++
	fmt.Println(a, *b) // [1 3 3] [1 3 3]
}

// Integer 定义一个 Integer 类型，基础类型是 int 类型
type Integer int

// 面向对象的方法，给 Integer 类增加一个 less 方法
func (intA Integer) less(intB Integer) bool {
	return intA < intB
}

// 面向过程的 less 方法
func less(intA Integer, intB Integer) bool {
	return intA < intB
}

// 为类型 Integer 再添加一个 add 方法，申明的是 值调用 add
func (intA Integer) add(intB Integer) {
	intA += intB
}

// 为类型 Integer 再添加一个 add 方法，此时申明的为 指针引用
func (intA *Integer) addRef(intB Integer) {
	*intA += intB
}
