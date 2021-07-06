package main

import (
	"fmt"
	"log"
	"net/http"
)

// ==========================
// http服务器示例程序
// ==========================

// 请求处理器，http请求会交给该处理器处理
func handler(writer http.ResponseWriter, req *http.Request) {
	// 格式化输出并使用writer写入到http client
	// 最后一个参数req.URL.Path[1:]，意思是将请求的url从第1个字符开始截取（也就是去掉第一个字符"/"，URL都是"/"开头）
	_, err := fmt.Fprintf(writer, "Hi there, I love %s!", req.URL.Path[1:])
	if err != nil {
		log.Println(err)
	}
}

func main() {
	// 将匹配"/"开头的请求，交给handler方法来处理
	http.HandleFunc("/", handler)

	fmt.Printf("Server starting on port %s.", "8080")
	// 开启服务，IP不写默认监听IP为本机ip(127.0.0.1, 或者实际物理IP都可以访问)，端口为8080
	// 当服务出现异常才会返回，并且仅返回一个异常信息
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	/*
		启动成功后，浏览器输入：http://localhost:8080/cat，可以看到浏览器上打印：Hi there, I love cat!
	*/
}
