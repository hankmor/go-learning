package main

import (
	"bytes"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	//fmt.Println("normalLog:")
	//normalLog()

	//fmt.Println("printCaller:")
	//printCaller()

	//fmt.Println("addField:")
	//addField()

	//fmt.Println("redirectOutput:")
	//redirectOutput()

	//fmt.Println("printJson:")
	//printJson()

	//fmt.Println("customFormatter:")
	//customFormatter()

	//fmt.Println("thirdPartyFormatter:")
	//thirdPartyFormatter()

	fmt.Println("write no output")
	writeNoOutput()
}

func normalLog() {
	logrus.SetLevel(logrus.TraceLevel)

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")
}

// 打印调用的方法信息
func printCaller() {
	logrus.SetReportCaller(true)
	logrus.Info("info msg")
}

// 通过调用logrus.WithField和logrus.WithFields添加一些公用字段
func addField() {
	logrus.WithFields(logrus.Fields{
		"name": "dj",
		"age":  18,
	}).Info("info msg")

	requestLogger := logrus.WithFields(logrus.Fields{
		"user_id": 10010,
		"ip":      "192.168.32.15",
	})

	requestLogger.Info("info msg")
	requestLogger.Error("error msg")
}

// 默认情况下，日志输出到io.Stderr。可以调用logrus.SetOutput传入一个io.Writer参数。后续调用相关方法日志将写到io.Writer中
func redirectOutput() {
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	// 输出到多个writer中
	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	logrus.Info("info msg")
}

// 将日志输出为json格式
func printJson() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")
}

func customFormatter() {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 03:04:05",
		FullTimestamp:   true,
	})
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Trace("trace msg...")
	logrus.Debug("debug msg...")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")
}

func thirdPartyFormatter() {
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})

	logrus.Info("info msg")

	logrus.SetFormatter(&nested.Formatter{
		// HideKeys:        true,
		// TimestampFormat: time.RFC3339,
		TimestampFormat: "2006-01-02 03:04:05",
		FieldsOrder:     []string{"name", "age"},
	})

	logrus.WithFields(logrus.Fields{
		"name": "dj",
		"age":  18,
	}).Info("info msg")
}

func writeNoOutput() {
	logger := logrus.New()
	logger.SetOutput(io.Discard)
	logger.WithFields(logrus.Fields{
		"url":     "http://foo.com",
		"attempt": 3,
		"backoff": time.Second,
	}).Info("failed to fetch URL")
}
