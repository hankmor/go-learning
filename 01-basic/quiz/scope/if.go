package main

import "fmt"

func main() {
	if a := 1; false {
	} else if b := 2; false {
	} else if c := 3; false {
	} else {
		fmt.Println(a, b, c)
	}
	/*out:
	1 2 3
	*/
	fmt.Println()
	ifUnfold()
}

// 上边的代码，展开后相当于：
func ifUnfold() {
	a := 1 // 变量提前
	if false {
	} else { // 隐式代码块
		b := 2
		if false { // 隐式代码块
		} else {
			c := 3
			if false { // 隐式代码块
			} else {
				fmt.Println(a, b, c)
			}
		}
	}
}
