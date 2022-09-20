package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 监听系统信号：即将系统信号抽象成os.Signal通道
	signalChan := make(chan os.Signal, 1) // 注意自1.17开始这里的chan必须是带缓冲的，见参考资料③
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	// 启动一个协程，协程内部读取（接收）系统信号转换的 chan os.Signal signalChan本身这个通道的自主关闭
	// 即当系统传递给进程一个信号时 signalChan 这个变量的 channel将会可读，否则一直阻塞
	go func() {
		<-signalChan // 此处没有系统信号时阻塞，后续代码不执行，有信号时后续代码执行

		signal.Stop(signalChan) // 显式停止监听系统信号
		close(signalChan)       // 显式关闭监听信号的通道
	}()

	// 初始化http-sever
	server := &http.Server{
		Addr:           ":9080",
		Handler:        nil,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	// 主进程（主协程）阻塞channel，以便控制http-server优雅退出后才退出主进程（主协程）
	// 就是一个简单的空结构体channel
	idleCloser := make(chan struct{})

	// 启动一个监听系统信号控制的channel
	go func() {
		<-signalChan

		// 超时context
		timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer timeoutCancel()

		// 系统包提供的优雅退出方法，见参考资料①
		if err := server.Shutdown(timeoutCtx); err != nil {
			// Error from closing listeners, or context timeout:
			fmt.Println("Http服务暴力停止，一般是达到超时context的时间当前还有尚未完结的http请求：" + err.Error())
		} else {
			fmt.Println("Http服务优雅停止")
		}

		// 关闭 主进程（主协程）阻塞channel，本子协程安全退出然后下方57行的阻塞终止主进程安全退出
		close(idleCloser)
	}()

	// 启动进入for循环的http-server，启动返回了不为 http.ErrServerClosed 的 error 时则表示启动有异常，例如端口号被占用
	// http.ErrServerClosed 错误则是当前server正在关闭中，当多个协程控制启动关闭通信不当时可能会出现这种情况
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("Http服务异常：" + err.Error())
		close(idleCloser)
	}

	// 通过空结构体channel阻塞主进程（主协程）达到持续运行的目的
	<-idleCloser
	fmt.Println("进程已退出：服务已关闭")
}
