package main

import (
	"fmt"
	"time"

	"github.com/huzhouv/go-learning/util"
)

func main() {
	// 测试\r，看到world覆盖了长度与其相等的hello: worldabcd
	// fmt.Printf("helloabcd\rworld\n")
	// 主程序未退出，则一直展示旋转动画
	go util.Spinner()
	n := 6
	i := fib(n)
	fmt.Printf("fib %d = %d\n", n, i)
}

// 斐波那契数列: 1 1 2 3 5
func fib(n int) int {
	if n > 2 {
		return fib(n-1) + fib(n-2)
	}
	return 1
}

// 自旋，防止主程序退出，一个转动的动画效果
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r) // \r 回车符，每次覆盖上一次的内容
			time.Sleep(delay)
		}
	}
}
