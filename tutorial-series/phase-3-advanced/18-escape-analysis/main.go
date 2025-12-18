package main

type User struct {
    Name string
}

// 这是一个工厂函数
func NewUser() *User {
    u := User{Name: "Hank"} // 本地变量 u
    return &u               // 返回了 u 的地址！
}

func main() {
    _ = NewUser()
}
