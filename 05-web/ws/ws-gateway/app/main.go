// main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	port      = flag.Int("p", 8080, "The server port, default: -p 8080")
	serviceId = flag.String("i", "ws-gateway", "The server id, default: ws-gateway")
)

func main() {
	flag.Parse()

	// 注册Consul服务
	registerService(*port)
	defer deregisterService()

	// 启动服务状态监控
	stopChan := make(chan struct{})
	go watchServiceChanges(stopChan)

	r := gin.Default()

	r.Static("/static", "./client")
	// 启动HTTP服务器
	r.GET("/ws", wsHandler)
	r.GET("/health", healthHandler)

	go func() {
		if err := r.Run(fmt.Sprintf(":%d", *port)); err != nil {
			panic(err)
		}
	}()

	// 优雅关闭处理
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	close(stopChan)

	setUnhealthy()              // 标记为不健康状态
	time.Sleep(5 * time.Second) // 等待Consul检测
}
