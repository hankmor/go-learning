package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var log *zap.Logger

func main() {
	fmt.Println("basic demo:")
	basic()

	fmt.Println("write file")
	writeFile()

	fmt.Println("write file and console")
	writeFileAndConsole()

	fmt.Println("write file by log level")
	writeFileByLogLvl()
}

func basic() {
	log = zap.NewExample()
	log.Debug("This is a DEBUG message")
	log.Info("This is an INFO message")
	// {"level":"debug","msg":"This is a DEBUG message"}
	// {"level":"info","msg":"This is an INFO message"}

	fmt.Println()
	log, _ = zap.NewDevelopment()
	log.Debug("This is a DEBUG message")
	log.Info("This is an INFO message")
	// 2023-02-10T10:25:10.819+0800    DEBUG   zap/zap.go:16   This is a DEBUG message
	// 2023-02-10T10:25:10.819+0800    INFO    zap/zap.go:17   This is an INFO message

	fmt.Println()
	log, _ = zap.NewProduction()
	log.Debug("This is a DEBUG message")
	log.Info("This is an INFO message")
	// {"level":"info","ts":1675995910.819356,"caller":"zap/zap.go:22","msg":"This is an INFO message"}
	// 没有输出debug日志

	fmt.Println()
	var sugarLogger *zap.SugaredLogger
	logger, _ := zap.NewProduction()
	sugarLogger = logger.Sugar() // 更友好但稍慢的api
	defer sugarLogger.Sync()
	var hello = "hello, zap"
	sugarLogger.Info("info message: ", hello)
	sugarLogger.Infow("info message: ", "say-hello", hello) // info with, key value形式打印
	sugarLogger.Infof("info message: %s", hello)
	sugarLogger.Infoln("info message: ", hello)
	// {"level":"info","ts":1675996578.8245518,"caller":"zap/zap.go:36","msg":"info message: hello, zap"}
	// {"level":"info","ts":1675996578.824576,"caller":"zap/zap.go:37","msg":"info message: ","say-hello":"hello, zap"}
	// {"level":"info","ts":1675996578.824599,"caller":"zap/zap.go:38","msg":"info message: hello, zap"}
	// {"level":"info","ts":1675996578.824607,"caller":"zap/zap.go:39","msg":"info message:  hello, zap"}
}

var timeEncoder = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")

func writeFile() {
	writeSyncer, _ := os.Create("./info.log")         // 日志文件存放目录
	encoderConfig := zap.NewProductionEncoderConfig() // 指定时间格式
	encoderConfig.EncodeTime = timeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)               // 获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel) // 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	log = zap.New(core, zap.AddCaller())                              // AddCaller()为显示文件名和行号
	log.Info("hello world")
	log.Error("hello world")
}

func writeFileAndConsole() {
	encoderConfig := zap.NewProductionEncoderConfig() // 指定时间格式
	encoderConfig.EncodeTime = timeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig) // 获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式

	// 文件writeSyncer
	fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./info1.log", // 日志文件
		MaxSize:    1,             // 文件大小限制,单位MB
		MaxBackups: 5,             // 最大保留日志文件数量
		MaxAge:     30,            // 日志文件保留天数
		Compress:   false,         // 是否压缩处理
	})

	mws := zapcore.NewMultiWriteSyncer(fileWriteSyncer, zapcore.AddSync(os.Stdout))
	fileCore := zapcore.NewCore(encoder, mws, zapcore.DebugLevel) // 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	log = zap.New(fileCore, zap.AddCaller())                      // AddCaller()为显示文件名和行号
	log.Info("hello world")
	log.Error("hello world")
}

func writeFileByLogLvl() {
	var coreArr []zapcore.Core

	// 获取编码器
	encoderConfig := zap.NewProductionEncoderConfig() // NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.EncodeTime = timeEncoder            // 指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	// encoderConfig.EncodeCaller = zapcore.FullCallerEncoder       // 显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // error级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // info和debug级别,debug级别是最低的
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	// info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/info.log", // 日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    1,                // 文件大小限制,单位MB
		MaxBackups: 5,                // 最大保留日志文件数量
		MaxAge:     30,               // 日志文件保留天数
		Compress:   false,            // 是否压缩处理
	})
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority) // 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	// error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/error.log", // 日志文件存放目录
		MaxSize:    1,                 // 文件大小限制,单位MB
		MaxBackups: 5,                 // 最大保留日志文件数量
		MaxAge:     30,                // 日志文件保留天数
		Compress:   false,             // 是否压缩处理
	})
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority) // 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)
	log = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()) // zap.AddCaller()为显示文件名和行号，可省略

	log.Info("hello info")
	log.Debug("hello debug")
	log.Error("hello error")
}
