package main

import "fmt"

// SumInts 将map的所有value相加
func SumInts(m map[string]int64) int64 {
	var ret int64
	for _, v := range m {
		ret += v
	}
	return ret
}

// SumFloats 将map的所有value相加
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Number 声明类型约束， Number 可以表示 int64 和 float64 两种类型
type Number interface {
	int64 | float64
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

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
	fmt.Printf("Non_Generic Sums: %v and % v\n", SumInts(ints), SumFloats(floats))

	// 调用泛型方法
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	// 省略泛型参数，自动推断
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	// 使用类型约束的 Number 类型
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}
