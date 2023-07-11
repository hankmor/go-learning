package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	// 基本类型: int, uint, float, complex, string
	bvals := []any{1, 1.0, 'a', "", "haha", errors.New("an error")}
	for _, v := range bvals {
		ref(v)
	}

	// 引用类型：map, struct, chan, pointer, slice, array, interface
	c := make(chan int)
	type Interface interface {
		Print()
	}
	var i *Interface // kind: invalid
	var i1 *It
	i1 = new(It) // kind: struct
	rvals := []any{struct{}{}, map[string]any{}, []string{}, [2]int{1, 2}, c, &c, i, i1}
	for _, v := range rvals {
		ref(v)
	}
}

type It struct {
}

func (it *It) Print() {}

func ref(value any) {
	reflectValue := reflect.Indirect(reflect.ValueOf(value))
	// fmt.Printf("reflectValue: %v\n", reflectValue)
	for reflectValue.Kind() == reflect.Ptr || reflectValue.Kind() == reflect.Interface {
		reflectValue = reflect.Indirect(reflectValue)
	}
	fmt.Printf("reflectValue: %v\n", reflectValue)
	fmt.Printf("reflectValue.Kind(): %v\n", reflectValue.Kind())

	switch reflectValue.Kind() {
	case reflect.Slice, reflect.Array:
	case reflect.Struct:
	case reflect.Interface:
	case reflect.Map:
	case reflect.Pointer:
	case reflect.Chan:
	case reflect.Bool:
	case reflect.Int | reflect.Int8 | reflect.Int16 | reflect.Int32 | reflect.Int64:
	// case reflect.Uint | reflect.Uint8 | reflect.Uint16 | reflect.Uint32 | reflect.Uint64:
	case reflect.Float32 | reflect.Float64:
	case reflect.Complex64 | reflect.Complex128:
	default:
	}
}
