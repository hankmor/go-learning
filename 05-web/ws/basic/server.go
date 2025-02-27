package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	var upgrader websocket.Upgrader
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, w.Header())
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		defer conn.Close()
		fmt.Println("客户端已经连接: ", conn.RemoteAddr().String())
		for {
			mt, msg, err := conn.ReadMessage()
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				fmt.Println("客户端连接异常断开")
				break
			} else if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				fmt.Println("客户端主动断开连接")
				break

			} else {
				if err != nil {
					fmt.Printf("read message error: %v\n", err)
					break
				}
			}
			fmt.Printf("received message: %v\n", string(msg))
			if err := conn.WriteMessage(mt, msg); err != nil {
				fmt.Printf("write message error: %v", err)
				break
			}
		}
	})
	fmt.Println("服务器已经启动，等待客户端连接...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
