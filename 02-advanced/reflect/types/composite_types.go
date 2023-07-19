package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 反射操作符合类型

	// 数组
	var a = [5]int{1, 2, 3, 4, 5}
	vaa := reflect.ValueOf(a) // reflect Value of Address of arr
	va0 := vaa.Index(0)
	fmt.Printf("a0 = [%d], va0 = [%d]\n", a[0], va0.Int()) // a0 = [1], va0 = [1]
	//va0.SetInt(100 + 1) // panic: reflect: reflect.Value.SetInt using unaddressable value
	//fmt.Printf("after set, a0 = [%d]\n", a[0]) // after set, a0 = [101]

	// 数组指针
	vaa = reflect.ValueOf(&a) // reflect Value of Address of arr
	va := vaa.Elem()
	va1 := va.Index(0)                                     // 根据获取下标获取value
	fmt.Printf("a0 = [%d], va1 = [%d]\n", a[0], va1.Int()) // a0 = [1], va0 = [1]
	va1.SetInt(100 + 1)                                    // 正常修改
	fmt.Printf("after set, a0 = [%d]\n", a[0])             // after set, a0 = [101]

	// 切片
	var s = []int{1, 1, 3}
	//vs := reflect.ValueOf(s)
	//vs0 := vs.Index(0)
	vss := reflect.ValueOf(&s)
	vs := vss.Elem()
	vs0 := vs.Index(0)                                     // 不能用在指针类型的value，需要先 Elem 获取到具体value
	fmt.Printf("s0 = [%d], vs0 = [%d]\n", s[0], vs0.Int()) // s0 = [1], vs0 = [1]
	vs0.SetInt(vs0.Int() + 100)
	fmt.Printf("after set, s0 = [%d]\n", s[0]) // after set, s0 = [101]

	// map
	var m = map[int]string{
		1: "tom",
		2: "jerry",
		3: "lucy",
	}
	vm := reflect.ValueOf(m)
	vm1V := vm.MapIndex(reflect.ValueOf(1))                      // MapIndex可以根据map的key获取value
	fmt.Printf("m_1 = [%s], vm_1 = [%s]\n", m[1], vm1V.String()) // m_1 = [tom], vm_1 = [tom]
	vm.SetMapIndex(reflect.ValueOf(1), reflect.ValueOf("tony"))  // 修改value值
	fmt.Printf("after set, m_1 = [%s]\n", m[1])                  // after set, m_1 = [tony]
	// 为map m新增一组key-value，如果key存在，直接修改
	vm.SetMapIndex(reflect.ValueOf(1), reflect.ValueOf("hank"))
	vm.SetMapIndex(reflect.ValueOf(4), reflect.ValueOf("amy"))
	fmt.Printf("after set, m = [%#v]\n", m) // after set, m = [map[int]string{1:"hank", 2:"jerry", 3:"lucy", 4:"amy"}]

	// 结构体
	var f = Foo{
		Name: "lily",
		age:  16,
	}
	vaf := reflect.ValueOf(&f)
	vf := vaf.Elem()
	field1 := vf.FieldByName("Name")
	fmt.Printf("the Name of f = [%s]\n", field1.String()) // the Name of f = [lily]
	field2 := vf.FieldByName("age")
	fmt.Printf("the age of f = [%d]\n", field2.Int()) // the age of f = [16]

	field1.SetString("ally")
	// field2.SetInt(8) // panic: reflect: reflect.Value.SetInt using value obtained using unexported field
	// NewAt创建一个新的 int 实例，并指向filed2属性地址，field2必须可以寻址，如果 reflect.ValueOf(f)获取的field 则会 panic
	nAge := reflect.NewAt(field2.Type(), unsafe.Pointer(field2.UnsafeAddr())).Elem()
	// 此时可以给 nAge 赋值，由于指向的指针地址，field2属性会跟着变化
	nAge.SetInt(8)
	fmt.Printf("after set, f is [%#v]\n", f) // after set, f is [main.Foo{Name:"ally", age:8}]

	// 接口
	var g = Foo{
		Name: "Jordan",
		age:  40,
	}
	// 接口底层动态类型为复合类型变量
	var i interface{} = &g // 接口类型变量接收
	vi := reflect.ValueOf(i)
	vg := vi.Elem()
	field1 = vg.FieldByName("Name")
	fmt.Printf("the Name of g = [%s]\n", field1.String()) // the Name of g = [Jordan]
	field2 = vg.FieldByName("age")
	fmt.Printf("the age of g = [%d]\n", field2.Int()) // the age of g = [40]
	// field2必须是指针下的可寻址value
	nAge = reflect.NewAt(field2.Type(), unsafe.Pointer(field2.UnsafeAddr())).Elem()
	nAge.SetInt(50)
	fmt.Printf("after set, g is [%#v]\n", g) // after set, g is [main.Foo{Name:"Jordan", age:50}]
	// 接口底层动态类型为基本类型变量
	var n = 5
	i = &n
	vi = reflect.ValueOf(i).Elem()
	fmt.Printf("i = [%d], vi = [%d]\n", n, vi.Int()) // i = [5], vi = [5]
	vi.SetInt(10)
	fmt.Printf("after set, n is [%d]\n", n) // after set, n is [10]

	// channel
	var ch = make(chan int, 100)
	vch := reflect.ValueOf(ch)    // 获取 ch 反射环境下的 value
	vch.Send(reflect.ValueOf(22)) // 等价于非反射环境下的 ch <- 22
	j := <-ch
	fmt.Printf("recv [%d] from channel\n", j) // recv [22] from channel
	ch <- 33
	vj, ok := vch.Recv()                           // 等价与非反射环境下的 <- ch
	fmt.Printf("recv [%d] ok[%t]\n", vj.Int(), ok) // recv [33] ok[true]
}

type Foo struct {
	Name string
	age  int // 未导出
}
