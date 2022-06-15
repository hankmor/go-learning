package main

import (
	"fmt"
	"reflect"
)

func main() {
	// getType()
	// typeAndKind()
	// elem()
	structIntrospection()
}

func getType() {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind()) // 类型名称和类型
}

type Enum int

func typeAndKind() {
	// Type: 指的是系统原生数据类型，如 int、string、bool、float32 等类型，以及使用 type 关键字定义的类型，这些类型的名称就是其类型本身的名称。
	// Kind: Type代表的底层类型

	// Map、Slice、Chan 属于引用类型，使用起来类似于指针，但是在种类常量定义中仍然属于独立的种类，不属于 Ptr。type A struct{} 定义的结构体属于 Struct 种类，*A 属于 Ptr。

	// reflect.TypeOf() 方法返回 Type 接口，包含 Name() 方法返回Type名称，Kind() 方法返回 reflect.Kind 类型的常量

	const Zero Enum = 0

	// 声明一个空结构体
	type cat struct {
	}

	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(cat{})
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind()) // cat struct
	// 获取Zero常量的反射类型对象
	typeOfA := reflect.TypeOf(Zero)
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind()) // Enum int
}

func elem() {

	// reflect.Elem() 获取Array, Chan, Map, Pointer 和 Slice类型中元素的类型，其他类型会触发panic
	// 指针获取反射对象时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型，这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作

	// 声明一个空结构体
	type cat struct {
	}

	println("==== int")
	var i = 10
	typeOfInt := reflect.TypeOf(i)
	fmt.Println(typeOfInt.Name(), typeOfInt.Kind())
	// fmt.Println(typeOfInt.Elem()) // panic: reflect: Elem of invalid type int

	println("==== slice")
	var arrs = []int{10}
	typeOfArr := reflect.TypeOf(arrs)
	fmt.Println(typeOfInt.Name(), typeOfInt.Kind()) // int int
	elemType := typeOfArr.Elem()
	fmt.Println(elemType.Name(), elemType.Kind()) // int int

	println("==== map")
	var mp = make(map[string]any, 4)
	mp["a"] = 1
	mp["b"] = 1.0
	mp["c"] = []int{1, 2}
	mp["d"] = &cat{}
	typeOfMap := reflect.TypeOf(mp)
	// 显示反射类型对象的名称和种类
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfMap.Name(), typeOfMap.Kind()) // name:'' kind:'map'
	// 取类型的元素
	elemType = typeOfMap.Elem()
	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", elemType.Name(), elemType.Kind()) // element name: '', element kind: 'interface'

	println("==== pointer")

	// 创建cat的实例，这是一个指针
	ins := &cat{}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 显示反射类型对象的名称和种类
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind()) // name:'' kind:'ptr'
	// 取类型的元素
	typeOfCat = typeOfCat.Elem()
	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind()) // element name: 'cat', element kind: 'struct'
}

type User struct {
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
}

func (u *User) Info() string {
	return fmt.Sprintf("name: %v, age: %d", u.Name, u.Age)
}

type Student struct {
	User   `json:"user"`
	school string
	class  string
}

func (s *Student) Intro() string {
	info := s.Info()
	return fmt.Sprintf("%v, school: %s, class: %s", info, s.school, s.class)
}

func structIntrospection() {
	// 通过 reflect.Type 的 NumField() 和 Field() 方法获得结构体成员的详细信息
	user := Student{User: User{Name: "张三", Age: 20}, school: "xxx学校", class: "1年级1班"}
	typeOfUser := reflect.TypeOf(user)
	valueOfUser := reflect.ValueOf(user)
	if typeOfUser.Kind() == reflect.Struct {
		// 属性
		fi := typeOfUser.NumField()
		for i := 0; i < fi; i++ {
			sf := typeOfUser.Field(i)
			fmt.Printf("index: %d, filed name: %v, type:%v, value: %v, tag: %v\n", sf.Index, sf.Name, sf.Type, valueOfUser.Field(i), sf.Tag.Get("json"))
		}
		/*
			index: [0], filed name: User, type:main.User, value: {张三 20}, tag: user
			index: [1], filed name: school, type:string, value: xxx学校, tag:
			index: [2], filed name: class, type:string, value: 1年级1班, tag:
		*/
		// 获取tag
		if catType, ok := typeOfUser.FieldByName("Name"); ok {
			// 从tag中取出需要的tag
			fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("form")) // name name
		}
		// 方法
		typeOfUser.NumMethod()
	} else {
		println("not a struct")
	}

	println("invalid tag")
	type cat struct {
		Name string
		Type int `json:"type" id:"100"`
		// Type int `json: "type" id:"100"` // 标签多了一个空格，不能正常解析，导致输出空字符串
	}
	typeOfCat := reflect.TypeOf(cat{})
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		fmt.Println(catType.Tag.Get("json"))
	}
}
