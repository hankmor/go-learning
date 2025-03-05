// websocket_server.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
)

func wsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	// 生成唯一客户端ID
	clientID := generateClientID()

	// 存储会话信息到Redis
	err = redisClient.HSet(c,
		"ws:session:"+clientID,
		"gateway_node", serviceId,
		"last_active", time.Now().Unix(),
	).Err()
	if err != nil {
		log.Println("Redis store failed:", err)
		return
	}

	// html页面无法发送ping帧
	conn.SetPingHandler(func(appData string) error {
		fmt.Println("收到ping消息")
		return conn.WriteMessage(websocket.PongMessage, []byte("pong"))
	})

	conn.SetPongHandler(func(appData string) error {
		fmt.Println("收到pong消息")
		return nil
	})

	conn.SetCloseHandler(func(code int, text string) error {
		log.Printf("🔌 WebSocket 关闭, 代码: %d, 原因: %s\n", code, text)
		return nil
	})

	// 消息处理循环
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("读取消息失败: %v ", err)
			removeSession(clientID)
			break
		}

		switch messageType {
		case websocket.TextMessage:
			if string(p) == "ping" {
				fmt.Println("收到心跳消息: 回复pong")
				if err := conn.WriteMessage(websocket.PongMessage, nil); err != nil {
					fmt.Println("Write error:", err)
					return
				}
			} else {
				// 更新最后活跃时间
				redisClient.HSet(c,
					"ws:session:"+clientID,
					"last_active", time.Now().Unix(),
				)

				// 处理消息（示例：回声服务）
				if err := conn.WriteMessage(messageType, p); err != nil {
					fmt.Println("Write error:", err)
					return
				}
			}
		}
	}
}

func removeSession(clientID string) {
	redisClient.Del(context.Background(), "ws:session:"+clientID)
}

func generateClientID() string {
	// 实现基于UUID的生成逻辑
	return "client-" + uuid.New().String()
}
