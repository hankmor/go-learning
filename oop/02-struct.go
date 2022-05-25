package oop

import (
	"fmt"
	"log"
	"os"
)

func ObjMain() {
	fmt.Println("结构体")

	// ===== 创建对象 =====

	// 未进行显式初始化的变量都会被初始化为该类型的零值
	rect := new(Rect) // 通过new创建，此时rect是指向Rect实例的指针
	fmt.Println(rect)
	r := Rect{} // 通过指针赋值
	fmt.Println(r)
	rect = &Rect{} // 通过指针赋值
	fmt.Println(rect)
	rect = &Rect{} // 通过指针赋值
	fmt.Println(rect)
	// 按照字段顺序赋值
	rect = &Rect{0, 0, 5, 10}
	fmt.Println(rect)
	// 指定字段名称赋值，可以无序
	rect = &Rect{width: 10, height: 100}
	fmt.Println(rect)

	calcArea()

	/*
		&{0 0 0 0}
		{0 0 0 0}
		&{0 0 0 0}
		&{0 0 0 0}
		&{0 0 5 10}
		&{0 0 10 100}
		area:  50
	*/

	// ===== 匿名组合 =====
	// go没有继承，但是提供了组合，称为匿名组合
	foo := Foo{
		Base: Base{"base"},
		Name: "foo",
	}
	foo.Bar()
	/*
		base Bar
		foo Bar
	*/

	// 匿名组合的一个用法
	job := Job{"copy", log.New(os.Stdout, "", log.LstdFlags)}
	job.start()
	/*
		2022/03/28 18:40:27 copy starting...
		2022/03/28 18:40:27 copy executing...
		2022/03/28 18:40:27 copy finished
	*/
}

// Rect 定义结构体
// Go中放弃了大量面向对象的特性，比如继承、多态、封装，只保留了组合
type Rect struct {
	x, y          float64
	width, height float64
}

// Area 给 Rect 结构体增加计算面积的方法
func (r *Rect) Area() float64 {
	return r.width * r.height
}

func calcArea() {
	rect := Rect{
		width:  5,
		height: 10,
	}
	area := rect.Area()
	fmt.Println("area: ", area) // 50
}

// 组合

type Base struct {
	Name string
}

func (base *Base) Foo() {
	fmt.Println(base.Name + " Foo")
}

func (base *Base) Bar() {
	fmt.Println(base.Name + " Bar")
}

type Foo struct {
	Base // 匿名组合一个 Base 对象
	Name string
}

func (foo *Foo) Bar() {
	foo.Base.Bar()
	fmt.Println(foo.Name + " Bar")
}

type Job struct {
	Command     string
	*log.Logger // 匿名组合了一个log.Logger的指针
}

func (job *Job) start() {
	// 可以直接调用job的 *log.Logger中的方法，甚至不需要知道它的存在
	job.Print(job.Command + " starting...")
	job.Println(job.Command + " executing...")
	job.Println(job.Command + " finished")
}
