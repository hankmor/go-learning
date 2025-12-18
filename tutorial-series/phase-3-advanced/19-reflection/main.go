package main

import (
	"fmt"
	"reflect"
	"strings"
)

func Marshal(v interface{}) (string, error) {
    t := reflect.TypeOf(v)
    val := reflect.ValueOf(v)

    // 只能处理结构体
    if t.Kind() != reflect.Struct {
        return "", fmt.Errorf("only support struct")
    }

    var sb strings.Builder
    sb.WriteString("{")

    // 遍历所有字段
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)     // 获取字段定义（如 Name string）
        value := val.Field(i)   // 获取字段值（如 "Hank"）
        
        // 获取 Tag (json:"name")
        tag := field.Tag.Get("json")
        if tag == "" {
            tag = field.Name // 没 tag 就用字段名
        }

        if i > 0 {
            sb.WriteString(",")
        }
        
        // 拼接 "key":"value"（简化版，仅支持 string 和 int）
        fmt.Fprintf(&sb, `"%s":`, tag)
        
        switch value.Kind() {
        case reflect.String:
            fmt.Fprintf(&sb, `"%s"`, value.String())
        case reflect.Int:
            fmt.Fprintf(&sb, `%d`, value.Int())
        default:
            fmt.Fprintf(&sb, `"unsupported"`)
        }
    }

    sb.WriteString("}")
    return sb.String(), nil
}

type User struct {
    Name string `json:"user_name"`
    Age  int    `json:"user_age"`
}

func main() {
    u := User{"Hank", 18}
    jsonStr, _ := Marshal(u)
    fmt.Println(jsonStr) 
    // Output: {"user_name":"Hank","user_age":18}
}
