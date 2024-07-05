package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	wsurl := url.URL{Scheme: "ws", Host: "127.0.0.1:8080", Path: "/ws"}
	fmt.Printf("connecting to %s\n", wsurl.String())
	conn, r, err := websocket.DefaultDialer.Dial(wsurl.String(), nil)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer conn.Close()

	// 优雅的关闭 socket 连接
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	fmt.Println("conn response:", r)
	// read asynchronously
	go func() {
		for {
			mt, msg, err := conn.ReadMessage()
			if err != nil {
				log.Fatal("read message error:", err)
			}
			fmt.Printf("received message: %s, type: %d\n", msg, mt)
			// fmt.Printf("received message: %s, type: %d", msg, websocket.FormatMessageType(mt))
		}
	}()

	// send message
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Fatal("send message error:", err)
				break
			}
		case <-interrupt: // 接收到中断消息
			fmt.Println("interrupt")
			// 发送一个关闭消息然后等待服务端关闭连接， 服务端会收到 read: websocket: close 1000 (normal) 消息
			if err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
				log.Println("write close:", err)
			}
			// 如果操作超时，退出客户端链接
			select {
			case <-time.After(time.Second):
			}
			return
		}
	}
}
