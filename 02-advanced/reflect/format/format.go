package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// 将任意类型格式化为字符串

func Any(a any) string {
	return formatAtom(reflect.ValueOf(a))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() { // 类型的种类Kind，将类型归类，如：int、uint、string、bool、struct、ptr、interface等等
	case reflect.Invalid: // 零值
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10) // 按照十进制int格式化
	case reflect.Uint8, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', 10, 64)
	case reflect.Complex64, reflect.Complex128:
		return strconv.FormatComplex(v.Complex(), 'f', 10, 64)
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Chan, reflect.Map, reflect.Func, reflect.Pointer, reflect.Slice: // 引用类型和函数、指针
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16) // 地址
	default:
		return v.Type().String() + " value" // 默认输出 v的类型Type value
	}
}

// Display 显示任何类型的所有元素类型和值
func Display(name string, a any) {
	fmt.Printf("display %s (%T): \n", name, a)
	display(name, reflect.ValueOf(a))
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ { // 遍历切片和数组
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i)) // 递归显示其中的每一个元素
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ { // 遍历结构体的所有字段，递归显示
			display(fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name), v.Field(i)) // 取出字段名称，然后递归，格式为 e.FiledName
		}
	case reflect.Map: // Map
		for _, k := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(k)), v.MapIndex(k)) // map的key直接格式化为string，然后递归显示
		}
	case reflect.Pointer: // 指针
		if v.IsNil() {
			fmt.Printf("%s is nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", v.Type()), v.Elem()) // 通过 v.Elem() 获取指针指向的具体Value类型
		}
	case reflect.Interface: // 接口类型
		if v.IsNil() {
			fmt.Printf("%s is nil\n", path)
		} else {
			fmt.Printf(fmt.Sprintf("%s.type = %s\n", path, v.Elem().Type())) // 输出接口对应的具体类型
			display(path+".value", v.Elem())
		}
	default: // 基本类型、chan、func
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func main() {
	// invalid
	fmt.Println(reflect.TypeOf(nil), Any(nil))

	// int
	var i1 int8 = 1
	var i2 int16 = 1
	fmt.Println(reflect.TypeOf(i1), Any(i1))
	fmt.Println(reflect.TypeOf(i2), Any(i2))

	// float and complex
	var flt = 3.1415
	fmt.Println(reflect.TypeOf(flt), Any(flt))
	var cmp = complex(flt, 1) // 复数由实部+虚部组成
	fmt.Println(reflect.TypeOf(cmp), Any(cmp))
	var cmp1 complex64
	cmp1 = 3.1415 + 1i
	fmt.Println(reflect.TypeOf(cmp1), Any(cmp1))

	// ref type
	var c chan int
	fmt.Println(reflect.TypeOf(c), Any(c))
	c = make(chan int)
	fmt.Println(reflect.TypeOf(c), Any(c))

	var slice []int
	fmt.Println(reflect.TypeOf(slice), Any(slice))
	slice = make([]int, 1)
	fmt.Println(reflect.TypeOf(slice), Any(slice))

	var m map[string]string
	fmt.Println(reflect.TypeOf(m), Any(m))
	m = make(map[string]string)
	fmt.Println(reflect.TypeOf(m), Any(m))

	type st struct {
	}
	var sti = st{}
	fmt.Println(reflect.TypeOf(sti), Any(sti))

	var f = func() {}
	fmt.Println(reflect.TypeOf(f), Any(f))

	println("=========== display ============")

	type programmer struct {
		Name        string
		Age         int8
		Hobby       []string
		From        *string
		ProgramLang map[string]float32
	}

	from := "secret"
	p := programmer{
		Name:  "huzhou",
		Age:   18,
		Hobby: []string{"coding", "reading", "playing"},
		From:  &from,
		ProgramLang: map[string]float32{
			"java":   95,
			"go":     80.5,
			"python": 35.5,
		},
	}
	Display("programmer", p)

	println()

	Display("os.Stderr", os.Stderr)
}
