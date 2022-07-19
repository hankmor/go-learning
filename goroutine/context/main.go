package main

import (
	"context-demo/server"
	"log"
	"net/http"
)

// 请求示例
// http://localhost:8080/search?q=golang&timeout=1s  展示请求结果
// http://localhost:8080/search?q=golang&timeout=1ms 请求timeout设置1ms，直接超时：context deadline exceeded

func main() {
	http.HandleFunc("/search", server.HandleSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
