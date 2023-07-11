package main

import (
	"sync"
	"time"
)

// 使用无缓冲的channel来模拟 4 x 100 接力赛

var rwg sync.WaitGroup

func Runner() {
	rwg.Add(1) // 只管理正在跑的那一位运动员状态

	var baton = make(chan int8) // 无缓冲通道，内部为运动员编号，接力棒通过运动员传递

	go Run(baton) // 第一棒
	baton <- 1    // 第一棒交给第一位运动员

	rwg.Wait()
}

func Run(baton chan int8) {
	runner := <-baton // 等待接力棒

	// 赛跑
	println("runner ", runner, " is running")
	time.Sleep(200 * time.Millisecond)

	// 是否需要创建下一个运动员
	var nextRunner int8
	if runner != 4 {
		nextRunner = runner + 1
		println("runner ", nextRunner, " get ready")
		go Run(baton) // 创建运动员后立刻执行，准备接收接力棒
	} else { // 比赛结束
		println("race over")
		rwg.Done()
		return
	}

	println("runner ", runner, " exchange the baton to next runner ", nextRunner)
	baton <- nextRunner // 交接
}

func main() {
	Runner()
}
