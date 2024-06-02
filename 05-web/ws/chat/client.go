package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// 允许写入到 peer 的最大时间
	writeWait = 10 * time.Second

	// 允许读取下次 pong 消息的最大时间
	pongWait = 60 * time.Second

	// 发送 ping 消息到 peer 的间隔时间，必须小于 pongWait
	pingPeriod = (pongWait * 9) / 10

	// 最大的消息大小
	maxMessageSize = 512
)

// Client 代表一个客户端连接
type Client struct {
	// 客户端的连接
	conn *websocket.Conn

	// 缓存的消息通道
	send chan []byte

	// 客户端需要与 hub 交互
	hub *Hub
}

// 从 conn 中读取消息到 hub
// 每一个 conn 都会创一个 goroutine 来调用该方法，需要确保 conn 上最多一个 reader
// 来并发调用该方法
func (c *Client) read() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	// 消息大小限制设置我饿0，则默认为 4096
	c.conn.SetReadLimit(maxMessageSize)
	// 可读取时间设置为 60s
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	// 设置 pong handler
	c.conn.SetPongHandler(func(appData string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		msgType, msg, err := c.conn.ReadMessage()
		fmt.Printf("msg type: %v", msgType)
		if err != nil {
			// 如果 conn 关闭
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				fmt.Printf("error: %v", err)
			}
			break
		}
		// 格式化消息
		msg = bytes.TrimSpace(bytes.Replace(msg, []byte{'\n'}, []byte{' '}, -1))
		// 向 hub 写入消息
		c.hub.broadcast <- msg
	}
}

// 从 hub 写入消息给 conn
// 每一个 conn 都会创建一个 goroutine 来调用该方法，需要确保 conn 上最多一个 writer
// 来并发调用该方法
func (c *Client) write() {
	// TODO why need ticker?
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case msg, ok := <-c.send: // 处理接收到的广播消息
			// 设置写入的最大时间
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			// hub 关闭了发送通道
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			}

			// 获取 writer
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(msg)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)

			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

var upgrader = websocket.Upgrader{}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// 升级 http 为 websocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error", err)
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 10)}
	client.hub.register <- client

	// 异步读写
	go client.read()
	go client.write()
}
