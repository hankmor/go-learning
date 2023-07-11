package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

// 使用 gob 序列化对象和反序列化

type user struct {
	Name   string
	Age    int
	Height float32
}

func (u *user) Info() string {
	return fmt.Sprintf("name: %s, age: %d, height: %.2f", u.Name, u.Age, u.Height)
}

func main() {
	// 序列化到文件
	f, err := os.Create("gob/myuser")
	if err != nil {
		fmt.Println(err)
	}
	u1 := user{Name: "huzhou", Age: 18, Height: 65.5}
	u2 := user{Name: "张三", Age: 28, Height: 60.5}
	us := []user{u1, u2}
	encoder := gob.NewEncoder(f)
	err = encoder.Encode(&us)
	if err != nil {
		fmt.Println(err)
	}

	// 从文件反序列化
	f, err = os.Open("gob/myuser")
	if err != nil {
		fmt.Println(err)
	}
	decoder := gob.NewDecoder(f)
	dus := make([]user, 2)
	err = decoder.Decode(&dus)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dus)
}
