package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 捕获系统信号示例

func main() {
	// Go signal notification works by sending `os.Signal`
	// values on a channel. We'll create a channel to
	// receive these notifications (we'll also make one to
	// notify us when the program can exit).
	sigs := make(chan os.Signal, 1) // 创建 chan 来接收信号通知
	done := make(chan bool, 1)      // 创建 chan 来检测知否接收到信号

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	// Notify 方法注册参数列表指定的系统信号到通知列表，产生这些系统信号后 sigs chan 会接收到，实现了将系统信号转为 chan.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGHUP)

	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	// 在 goroutine 中接收信号，未收到则阻塞，收到则打印这个信号，并向 done chan 写入 true，程序结束
	go func() {
		sig := <-sigs // 从 chan 中读取信号，未读取到则阻塞
		fmt.Println()
		fmt.Printf("sign: %v \n", sig) // 打印信号
		done <- true                   // 向 done chan 写入 true，退出程序
	}()

	// The program will wait here until it gets the
	// expected signal (as indicated by the goroutine
	// above sending a value on `done`) and then exit.
	fmt.Println("awaiting signal")
	<-done // done chan 未接收到值则阻塞，等到收到信号时会接收到 true，此时程序可以继续运行
	fmt.Println("exiting")
}
