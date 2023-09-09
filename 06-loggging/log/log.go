package main

import (
	"io"
	"log"
	"os"
)

func main() {
	defLog()
	writeToFile()
	writeToMultiWriter()
}

func writeToMultiWriter() {
	file, err := os.OpenFile("./multi.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, file)) // 多个writer
	log.Println("this log will write to console and multi.log")
}

func defLog() {
	log.Println("this is a go standard log message")
}

func writeToFile() {
	file, err := os.OpenFile("./demo.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	log.Println("this log msg will write to demo.log")
}
