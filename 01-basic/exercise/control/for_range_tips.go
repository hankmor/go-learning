package main

import (
	"fmt"
	"time"
)

// for range 的一些坑：
// 1）每次遍历的变量是重用的，只有一份
// 2）遍历数组时，需要拷贝一份数组，遍历的是拷贝后的数组，遍历过程中修改数组可能得不到预期结果

func forRangeReuseBad() {
	var m = []int{1, 2, 3, 4, 5}
	for i, v := range m { // i, v 在主和新 goroutine 中实现了共享，并且在整个循环过程中是重用的，仅有一份
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("i = %d, v = %d\n", i, v)
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("for range end")
	/* output:
	i = 4, v = 5
	i = 4, v = 5
	i = 4, v = 5
	i = 4, v = 5
	i = 4, v = 5
	for range end
	*/
}

func forRangeReuseCorrect() {
	var m = []int{1, 2, 3, 4, 5}
	for i, v := range m {
		go func(i, v int) {
			time.Sleep(1 * time.Second)
			fmt.Printf("i = %d, v = %d\n", i, v)
		}(i, v) // 通过 i, v 作为参数传递到新的 goroutine 中，实现与闭包绑定
	}
	time.Sleep(2 * time.Second)
	fmt.Println("for range end")
	/* output: (输出顺序无序)
	i = 3, v = 4
	i = 0, v = 1
	i = 4, v = 5
	i = 2, v = 3
	i = 1, v = 2
	for range end
	*/
}

func main() {
	fmt.Println("for range reused value:")
	forRangeReuseBad()
	forRangeReuseCorrect()

	fmt.Println("for range array copied value:")
	forRangeArrayCopiedValue()
	forRangeArrayPtrCopiedValue()

	fmt.Println("for range slice:")
	forRangeSlice()
	forRangeSliceChangeLen()
}

func forRangeArrayCopiedValue() {
	// 被遍历的是一个数组
	var m = [...]int{1, 2, 3, 4}
	var x [4]int

	fmt.Println("before modify: ")
	fmt.Println("  m: ", m)

	for i, v := range m { // for range 遍历数组，需要拷贝整个数组，性能有影响
		if i == 0 { // 第一次循环时想要修改数组的元素
			m[1] = 21
			m[2] = 31
		}
		x[i] = v // 修改了数组元素，然后赋值给新的数组
	}
	fmt.Println("after modify: ")
	fmt.Println("  m: ", m)
	fmt.Println("  x: ", x)
	/*output:
	before modify:
	  m:  [1 2 3 4]
	after modify:
	  m:  [1 21 31 4]
	  x:  [1 2 3 4]

	结果可以看出，修改了数组元素，但是循环过程中赋值给新的数组并不是新的元素，而是原元素的一个副本。
	*/
}

func forRangeArrayPtrCopiedValue() {
	// 被遍历的是一个数组
	var m = [...]int{1, 2, 3, 4}
	var x [4]int

	fmt.Println("before modify: ")
	fmt.Println("  m: ", m)

	for i, v := range &m { // for range 遍历数组指针，此时结果正确
		if i == 0 {
			m[1] = 21
			m[2] = 31
		}
		x[i] = v
	}
	fmt.Println("after modify: ")
	fmt.Println("  m: ", m)
	fmt.Println("  x: ", x)
	/*output:
	before modify:
	  m:  [1 2 3 4]
	after modify:
	  m:  [1 21 31 4]
	  x:  [1 21 31 4]
	*/
}

func forRangeSlice() {
	// 改为遍历切片
	var m = []int{1, 2, 3, 4}
	var x [4]int

	fmt.Println("before modify: ")
	fmt.Println("  m: ", m)

	for i, v := range m { // for range 遍历切片
		if i == 0 { // 第一次循环时想要修改切片的元素
			m[1] = 21
			m[2] = 31
		}
		x[i] = v // 修改了元素后，赋值给新的数组
	}
	fmt.Println("after modify: ")
	fmt.Println("  m: ", m)
	fmt.Println("  x: ", x)
	/*output:
	before modify:
	  m:  [1 2 3 4]
	after modify:
	  m:  [1 21 31 4]
	  x:  [1 21 31 4]

	从结果可以看出，遍历切片，修改切片元素后可以得到预期结果。
	*/
}

func forRangeSliceChangeLen() {
	var m = []int{1, 2, 3, 4}
	var x []int

	fmt.Println("before change length: ")
	fmt.Println("  m: ", m)

	for i, v := range m { // for range 遍历切片
		if i == 0 { // 第一次循环时想要修改切片的元素
			m = append(m, 5, 6)
		}
		x = append(x, v) // 试图将新添加的元素添加到另一个切片，无效
	}
	fmt.Println("after change length: ")
	fmt.Println("  m: ", m)
	fmt.Println("  x: ", x)
	/*output:
	before change length:
	  m:  [1 2 3 4]
	after change length:
	  m:  [1 2 3 4 5 6]
	  x:  [1 2 3 4]

	由于遍历时原 slice 底层的数组没有改变，所以即使遍历时修改了 slice 的长度，遍历还是原 slice。
	*/
}
