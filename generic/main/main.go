package main

import (
	"fmt"
	"github.com/huzhouv/go-learning/generic"
)

func main() {
	ints := map[string]int64{
		"one": 34,
		"two": 10,
	}

	floats := map[string]float64{
		"one": 3.1415,
		"two": 2.1719,
	}
	// 调用非泛型方法
	fmt.Printf("Non_Generic Sums: %v and % v\n", generic.SumInts(ints), generic.SumFloats(floats))

	// 调用泛型方法
	fmt.Printf("Generic Sums: %v and %v\n",
		generic.SumIntsOrFloats[string, int64](ints),
		generic.SumIntsOrFloats[string, float64](floats))

	// 省略泛型参数，自动推断
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		generic.SumIntsOrFloats(ints),
		generic.SumIntsOrFloats(floats))

	// 使用类型约束的 Number 类型
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		generic.SumNumbers(ints),
		generic.SumNumbers(floats))
}
