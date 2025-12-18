package main

import (
	"log/slog"
	"os"
)

func main() {
	// 创建 JSON 格式的 Logger，输出到文件
	file, err := os.OpenFile("app.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	logger := slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// 记录不同级别的日志
	logger.Info("Application started", "version", "1.0.0")
	logger.Warn("Configuration missing", "key", "database.host", "using_default", true)
	logger.Error("Failed to connect to database",
		"error", "connection timeout",
		"retry_count", 3)

	// 使用 With 创建带上下文的 logger
	userLogger := logger.With(
		"service", "user-service",
		"environment", "production",
	)

	userLogger.Info("User created", "user_id", 12345, "username", "john_doe")
	userLogger.Info("User updated", "user_id", 12345, "fields", []string{"email", "phone"})
}
