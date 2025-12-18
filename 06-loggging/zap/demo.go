package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// 1. 开发环境 Logger
	devLogger, _ := zap.NewDevelopment()
	defer devLogger.Sync()

	devLogger.Info("User logged in",
		zap.String("username", "admin"),
		zap.Int("user_id", 123))

	// 2. 生产环境 Logger
	prodLogger, _ := zap.NewProduction()
	defer prodLogger.Sync()

	prodLogger.Info("Server started",
		zap.String("port", "8080"))

	// 3. 自定义配置
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	config.OutputPaths = []string{
		"stdout",
		"app.log",
	}
	config.ErrorOutputPaths = []string{
		"stderr",
		"error.log",
	}

	customLogger, _ := config.Build()
	defer customLogger.Sync()

	customLogger.Info("Application started")

	// 4. 日志轮转
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    100, // MB
		MaxBackups: 3,
		MaxAge:     28, // 天
		Compress:   true,
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)

	rotateLogger := zap.New(core)
	defer rotateLogger.Sync()

	rotateLogger.Info("Log with rotation")

	// 5. 动态调整日志级别
	atom := zap.NewAtomicLevel()

	dynamicConfig := zap.NewProductionConfig()
	dynamicConfig.Level = atom

	dynamicLogger, _ := dynamicConfig.Build()

	dynamicLogger.Info("This will be logged")
	dynamicLogger.Debug("This won't be logged")

	// 动态调整为 Debug 级别
	atom.SetLevel(zapcore.DebugLevel)

	dynamicLogger.Debug("Now this will be logged")
}
