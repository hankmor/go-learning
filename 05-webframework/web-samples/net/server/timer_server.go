package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	server() // 启动服务端
}

// server 每秒向客户端写入当前时间，9秒后断开连接
// 客户端可以使用 telnet: telnet localhost 8000
// 也可以使用netcat(nc): nc localhost 8000
// 也可以使用 client 程序来观察并发连接执行情况
func server() {
	// 监听8000端口
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("service started on port: 8000")
	for {
		conn, err := listener.Accept() // 接收请求
		log.Println("client connected: ", conn.RemoteAddr())
		if err != nil {
			log.Fatal(err)
		}

		// handleConn(conn) // 每次只能处理一个连接
		go handleConn(conn) // 并发执行，同时处理多个连接
	}
}

func handleConn(conn net.Conn) {
	// 每秒打印当前时间9次，第10次断开连接
	for i := 0; i < 10; i++ {
		if i == 9 { // 断开连接
			log.Println("client quit: ", conn.RemoteAddr())
			err := conn.Close()
			if err != nil {
				log.Fatal(err)
			}
			break
		}
		_, err := io.WriteString(conn, fmt.Sprintf("%d - %s", i, time.Now().Format("2006-01-02 15:04:05")))
		_, err = conn.Write([]byte("\n"))
		if err != nil {
			log.Fatal(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
