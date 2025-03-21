package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
		cnt := 0
		tk := time.NewTicker(time.Second)
		defer tk.Stop()
		for {
			// _, msg, err := conn.ReadMessage()
			// if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
			// 	fmt.Println("客户端连接异常断开")
			// 	break
			// } else if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
			// 	fmt.Println("客户端主动断开连接")
			// 	break
			//
			// } else {
			// 	if err != nil {
			// 		fmt.Printf("read message error: %v\n", err)
			// 		break
			// 	}
			// }
			// 等待4秒测试当客户端推送结束并断开后，服务端是否还能继续接收数据，模拟处理数据延迟的情景
			// websocket是可靠协议，可以接收客户端发过来的所有数据，即使客户端已经关闭
			//          fmt.Println("wait 4s")
			// time.Sleep(time.Second * 4)
			// fmt.Printf("received message: %v\n", string(msg))
			select {
			case c := <-tk.C:
				if err := conn.WriteMessage(websocket.TextMessage, []byte(c.String())); err != nil {
					fmt.Printf("write message error: %v", err)
					break
				}
				cnt++
				if cnt > 5 {
                    // 关闭服务端，测试客户端是否可以继续收到数据
                    //  可以, 推送的数据会消费完
					os.Exit(-1)
				}
			}
		}
	})
	fmt.Println("服务器已经启动，等待客户端连接...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
