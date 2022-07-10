package util

import (
	"fmt"
	"time"
)

// SpinnerDelay 自旋，防止主程序退出，一个转动的动画效果
func SpinnerDelay(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r) // \r 回车符，每次覆盖上一次的内容
			time.Sleep(delay)
		}
	}
}

func Spinner() {
	SpinnerDelay(100 * time.Millisecond)
}
