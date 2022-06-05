package goroutine

import (
	"math/rand"
	"sync"
)

// 更好的同步方式是使用通道 channel，通过通道交换数据是线程安全的
// 无缓冲的通道：每次需要一个消费者从通道获取值后生产者才能再放一个数据到通道(否则阻塞)
// 带缓冲的通道：缓存数量满后生产者才阻塞

var ppsg sync.WaitGroup

// PingPong 乒乓球游戏，两位玩家只能等到对方击球后才可以接球，可以使用无缓冲通道
func PingPong() {
	ppsg.Add(2)                // 两个玩家同步
	var ppc = make(chan int16) // 无缓冲 channel
	go Player("huzhou", ppc)   // 玩家1
	go Player("belonk", ppc)   // 玩家2

	ppc <- 0    // 先发求
	ppsg.Wait() // 等待游戏结束
}

func Player(name string, ppc chan int16) {
	println(name, " joined the game")
	defer ppsg.Done()
	for {
		// 击球次数，从通道获取
		ball, ok := <-ppc
		if !ok { // 通道关闭，则表示当前玩家胜利
			println(name, " win the game")
			return
		}
		// 随机数，模拟失败
		n := rand.Intn(100)
		if n%13 == 0 {
			println(name, " lost the game") // 失败
			close(ppc)                      // 失败了关闭通道
			return
		}

		ball++                       // 击球数加1
		println(name, " hit ", ball) // 打印击球信息
		ppc <- ball                  // 将球打给对方
	}
}
