package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.String("port", "8001", "port, default is 8001")

func main() {
	flag.Parse()

	// 创建 hub
	hub := newHub()
	// 开启hub
	go hub.run()

	http.HandleFunc("/", serveHtml)
	// 注册 websokcket 端点
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	addr := fmt.Sprintf(":%s", *port)
	log.Println("http server started on:", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("http server error:", err)
	}
}

func serveHtml(w http.ResponseWriter, r *http.Request) {
	log.Println("request url:", r.URL)
	if r.URL.Path != "/p" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}
