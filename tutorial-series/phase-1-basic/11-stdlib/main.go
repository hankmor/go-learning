package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    // json:"-" 表示在 JSON 中忽略该字段
    Email string `json:"-"`
}

func main() {
    // 1. fmt - 格式化
    age := 18
    // Sprintf 返回字符串，Printf 直接打印
    msg := fmt.Sprintf("I am %d years old", age)
    fmt.Println(msg)

    // 2. strings - 字符串处理
    str := " hello world  "
    fmt.Println(strings.TrimSpace(str))           // 去除首尾空格
    fmt.Println(strings.Contains(str, "world"))   // true
    fmt.Println(strings.Join([]string{"a", "b"}, "-")) // a-b

    // 3. time - 时间处理
    now := time.Now()
    // Go 的独特日期格式化模板：2006-01-02 15:04:05
    fmt.Println(now.Format("2006-01-02 15:04:05"))
    
    // 时间计算
    later := now.Add(time.Hour)
    fmt.Println(later.Sub(now)) // 1h0m0s

    // 4. encoding/json - JSON 编解码
    user := User{ID: 1, Name: "Hank", Email: "admin@hankmo.com"}
    
    // 序列化 Struct -> JSON
    jsonData, _ := json.Marshal(user)
    fmt.Println(string(jsonData)) // {"id":1,"name":"Hank"}

    // 反序列化 JSON -> Struct
    var u2 User
    json.Unmarshal(jsonData, &u2)
    fmt.Println(u2.Name)
}
