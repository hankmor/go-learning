package logger

import (
	"os"
	"testing"
)

func TestStdLogger(t *testing.T) {
	logger := New(&Config{
		Type:  LoggerTypeStd,
		Level: "debug",
	})

	logger.Debug("Debug message", "key", "value")
	logger.Info("Info message", "user_id", 123)
	logger.Warn("Warning message")
	logger.Error("Error message", "error", "something went wrong")
}

func TestSlogLogger(t *testing.T) {
	logger := New(&Config{
		Type:  LoggerTypeSlog,
		Level: "debug",
	})

	logger.Debug("Debug message", "key", "value")
	logger.Info("Info message", "user_id", 123)
	logger.Warn("Warning message")
	logger.Error("Error message", "error", "something went wrong")
}

func TestZapLogger(t *testing.T) {
	logger := New(&Config{
		Type:  LoggerTypeZap,
		Level: "debug",
	})

	logger.Debug("Debug message", "key", "value")
	logger.Info("Info message", "user_id", 123)
	logger.Warn("Warning message")
	logger.Error("Error message", "error", "something went wrong")
}

func TestFileOutput(t *testing.T) {
	tmpFile := "test.log"
	defer os.Remove(tmpFile)

	logger := New(&Config{
		Type:       LoggerTypeSlog,
		Level:      "info",
		OutputPath: tmpFile,
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   false,
	})

	logger.Info("Test message", "test", true)

	// 检查文件是否创建
	if _, err := os.Stat(tmpFile); os.IsNotExist(err) {
		t.Error("Log file was not created")
	}
}
