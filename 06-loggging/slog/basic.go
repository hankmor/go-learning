package main

import (
	"log/slog"
	"os"
)

func main() {
	// 1. 基础使用 - JSON 格式
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonLogger.Info("User logged in",
		"user_id", 123,
		"username", "admin",
		"ip", "192.168.1.1")

	// 2. 文本格式
	textLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	textLogger.Info("Server started", "port", 8080)

	// 3. 设置日志级别
	levelLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // 只输出 Info 及以上级别
	}))

	levelLogger.Debug("Debug message")  // 不会输出
	levelLogger.Info("Info message")    // 会输出
	levelLogger.Warn("Warning message") // 会输出
	levelLogger.Error("Error message")  // 会输出

	// 4. 使用 With 创建子 Logger
	logger := slog.Default()

	// 创建带有固定字段的子 Logger
	requestLogger := logger.With(
		"request_id", "abc123",
		"user_id", 456,
	)

	requestLogger.Info("Processing request")
	requestLogger.Info("Request completed")
	// 两条日志都会自动包含 request_id 和 user_id

	// 5. 分组字段
	groupLogger := slog.Default()
	groupLogger.Info("User action",
		slog.Group("user",
			slog.String("id", "123"),
			slog.String("name", "admin"),
		),
		slog.Group("action",
			slog.String("type", "login"),
			slog.String("ip", "192.168.1.1"),
		),
	)
}
