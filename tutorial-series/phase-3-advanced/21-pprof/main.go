package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // 关键：副作用导入，自动注册路由
	"time"
)

func datas() {
    for {
        log.Println("working...")
        time.Sleep(1 * time.Second)
    }
}

func main() {
    go datas()

    // 启动一个专门用于监控的 HTTP 服务
    // 访问 http://localhost:6060/debug/pprof/ 即可看到仪表盘
    fmt.Println("PPROF server started at :6060")
    if err := http.ListenAndServe(":6060", nil); err != nil {
        log.Fatal(err)
    }
}
