package main

import (
	"log"
	"net/http"
)

func main() {
	// 创建 hub
	hub := newHub()
	// 开启hub
	go hub.run()

	http.HandleFunc("/", serveHtml)
	// 注册 websokcket 端点
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
    log.Println("http server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("http server error:", err)
	}
}

func serveHtml(w http.ResponseWriter, r *http.Request) {
	log.Println("request url:", r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}
