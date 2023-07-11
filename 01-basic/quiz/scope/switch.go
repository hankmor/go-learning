package main

import "fmt"

func main() {
	switch a, b := 1, 2; a + b {
	case 3:
		x := 1
		fmt.Println("case1, x = ", x)
		fallthrough
	case 5:
		x := 2
		fmt.Println("case2, x = ", x)
		fallthrough
	default:
		x := 3
		fmt.Println("default, x = ", x)
	}
	/*output:
	case1, x =  1
	case2, x =  2
	default, x =  3
	*/
	fmt.Println()
	// 	上述代码等价于：
	switchUnfold()
}

func switchUnfold() {
	a, b := 1, 2
	switch a + b {
	case 3:
		{ // 隐式代码块, x 变量包含于其中
			x := 1
			fmt.Println("case1, x = ", x)
		}
		fallthrough
	case 10:
		{ // 隐式代码块, x 变量包含于其中
			x := 2
			fmt.Println("case2, x = ", x)
		}
		fallthrough
	default:
		{ // 隐式代码块, x 变量包含于其中
			x := 3
			fmt.Println("default, x = ", x)
		}
	}
}
