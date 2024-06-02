package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"util"
)

var wg sync.WaitGroup

func main() {
	go util.Spinner()
	client() // 客户端
}

func client() {
	wg.Add(10)
	// 开启10个连接
	for i := 0; i < 10; i++ {
		// connToServer() // 顺序连接
		go connToServer() // 并发连接
	}
	wg.Wait()
	os.Stdout.Write([]byte(fmt.Sprintf("\rfinished!")))
}

func connToServer() {
	// 连接服务器
	conn, err := net.Dial("tcp", "localhost:8000")
	log.Printf("\r[%s] connected to localhost:8000...", conn.LocalAddr())
	if err != nil {
		log.Fatal(err)
		return
	}

	go handlerResult(conn)
}

func handlerResult(conn net.Conn) {
	// 创建缓冲区
	buffer := bytes.NewBuffer(nil)
	var bs [128]byte // 每次读取的字节数
	for {
		buffer.Write([]byte{'\r'})  // 读取前先写入一个回车字符，效果是去掉 util.Spinner 的这个字符
		n, err := conn.Read(bs[0:]) // 将连接中的数据读取到bytes
		buffer.Write(bs[0:n])       // 写入到缓冲区，仅写读取到的长度
		if err != nil {
			if err == io.EOF { // 如果err是EOF，说明全部读取完成，退出连接
				conn.Close() // 关闭连接
				wg.Done()    // 读取完成，连接中指
				return
			} else {
				log.Fatal(err)
			}
			return
		}
		_, err = os.Stdout.Write(buffer.Bytes()) // 将读取的内容打印到控制台
		// _, handleerr = io.Copy(os.Stdout, conn) // 如果不去掉 util.Spinner 的字符，直接是哟经 io.Copy 更简单
		if err != nil {
			log.Fatal(err)
		}
	}
}
