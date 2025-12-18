package main

import (
	"log"
	"os"
)

func main() {
	// 1. 基础使用
	log.Println("This is a log message")
	log.Printf("User %s logged in", "admin")

	// 2. 自定义 Logger
	logger := log.New(
		os.Stdout,
		"[MyApp] ",                              // 前缀
		log.Ldate|log.Ltime|log.Lshortfile, // 标志
	)

	logger.Println("Custom logger message")
	// 输出：[MyApp] 2025/12/18 10:00:00 main.go:15: Custom logger message

	// 3. 输出到文件
	file, err := os.OpenFile("app.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println("This goes to file")

	// 4. 不同的日志级别（标准库 log 没有级别概念，这里演示不同的输出方式）
	log.SetOutput(os.Stdout) // 恢复输出到控制台

	log.Print("Print: regular message")
	log.Println("Println: with newline")
	log.Printf("Printf: formatted %s", "message")

	// Fatal 和 Panic 会终止程序，这里注释掉
	// log.Fatal("Fatal: exits with status 1")
	// log.Panic("Panic: triggers panic")
}
