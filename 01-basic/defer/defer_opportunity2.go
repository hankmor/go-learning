package main

import "fmt"

// defer 的执行时机

func f1() {
	for i := 0; i <= 3; i++ {
		fmt.Printf("%d ", i)
	}
	// 0 1 2 3
}

func f2() {
	for i := 0; i <= 3; i++ {
		defer func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}
	// 3 2 1 0
	// 因为 defer 出栈与入栈顺序相反（LIFO），所以最后入栈的先执行
}

func f3() {
	for i := 0; i <= 3; i++ {
		defer func() {
			fmt.Printf("%d ", i)
		}()
	}
	// 4 4 4 4
	// defer 执行时 i 的值为 4
}

// func main() {
// 	f1()
// 	fmt.Println()
// 	f2()
// 	fmt.Println()
// 	f3()
// }
