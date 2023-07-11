package main

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestBigNumber(t *testing.T) {
	mp := make(map[int]int, 0)
	mp[1] = 1
	mp[2] = 2
	fmt.Printf("%v\n", mp)

	var sl []int
	sl = append(sl, 1)
	fmt.Printf("%v\n", sl)

	a := float64(0.2)
	b := float64(0.1)
	println(strconv.FormatFloat(a-b, 'e', 22, 64))
	println(strconv.FormatFloat(a-b, 'g', 22, 64))

	numstr := "1515631351536151161464461151511561"
	// 加
	fmt.Println(BigIntAdd(numstr, 99))
	// 减
	fmt.Println(BigIntReduce(numstr, 99))
	// 乘
	fmt.Println(BigIntMul(numstr, 99))
	// 除
	fmt.Println(BigIntDiv(numstr, 99))
	println(math.MaxInt8)
	println(math.MaxUint8)
	println(math.MaxInt16)
	println(math.MaxInt32)
	println(math.MaxInt64)
	// fmt.Printf("%v", math.MaxUint64)
	println(math.MaxFloat32)
	println(math.MaxFloat64)
}
