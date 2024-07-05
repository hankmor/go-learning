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

		for {
			mt, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Printf("read message error: %v\n", err)
				break
			}
			fmt.Printf("received message: %v\n", string(msg))
			if err := conn.WriteMessage(mt, msg); err != nil {
				fmt.Printf("write message error: %v", err)
				break
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
