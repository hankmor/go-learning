package main

import "fmt"

// defer 执行的时机

func func1() {
	var s = []int{1, 2, 3}
	// defer 时传入 slice
	defer func(s []int) {
		fmt.Println("invoking defer with []int")
		for _, v := range s {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}(s)
	fmt.Println("changing slice")
	// 修改 slice
	s = []int{3, 2, 1}
	fmt.Println("slice changed")

	/*output:
	changing slice
	slice changed
	invoking defer with []int
	1 2 3

	结论：defer 在入栈时的函数为 func([]int{1,2,3})，后来更改 slice 并没有影响栈内这个函数，所以输出仍然是 1 2 3
	*/
}

func func2() {
	var s = []int{1, 2, 3}
	// defer 时传入 slice 的指针
	defer func(s *[]int) {
		fmt.Println("invoking defer with *[]int")
		for _, v := range *s {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}(&s)
	fmt.Println("changing slice")
	// 修改 slice
	s = []int{3, 2, 1}
	fmt.Println("slice changed")

	/*output:
	changing slice
	slice changed
	invoking defer with *[]int
	3 2 1

	结论：defer 入栈时的函数为 func(&s)，参数为 *[]int，这是一个指针，更改 slice 则 defer 匿名函数的参数变化，所以输出 3 2 1
	*/
}

// func main() {
// 	func1()
// 	func2()
// }
