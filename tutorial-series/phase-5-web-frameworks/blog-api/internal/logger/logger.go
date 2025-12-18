package logger

import (
	"io"
	"log"
	"log/slog"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// LoggerType 日志类型
type LoggerType string

const (
	LoggerTypeStd  LoggerType = "log"   // 标准库 log
	LoggerTypeSlog LoggerType = "slog"  // Go 1.21+ slog
	LoggerTypeZap  LoggerType = "zap"   // Uber zap
)

// Config 日志配置
type Config struct {
	Type       LoggerType // 日志类型
	Level      string     // 日志级别：debug, info, warn, error
	OutputPath string     // 输出路径，空则输出到 stdout
	MaxSize    int        // 日志文件最大大小（MB）
	MaxBackups int        // 保留的旧日志文件数量
	MaxAge     int        // 保留的旧日志文件天数
	Compress   bool       // 是否压缩旧日志文件
}

// Logger 统一的日志接口
type Logger interface {
	Debug(msg string, fields ...interface{})
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})
}

// New 创建日志实例
func New(cfg *Config) Logger {
	if cfg == nil {
		cfg = &Config{
			Type:  LoggerTypeSlog,
			Level: "info",
		}
	}

	switch cfg.Type {
	case LoggerTypeStd:
		return newStdLogger(cfg)
	case LoggerTypeSlog:
		return newSlogLogger(cfg)
	case LoggerTypeZap:
		return newZapLogger(cfg)
	default:
		return newSlogLogger(cfg)
	}
}

// ==================== 标准库 log ====================

type stdLogger struct {
	logger *log.Logger
}

func newStdLogger(cfg *Config) *stdLogger {
	var output io.Writer = os.Stdout

	if cfg.OutputPath != "" {
		output = &lumberjack.Logger{
			Filename:   cfg.OutputPath,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		}
	}

	return &stdLogger{
		logger: log.New(output, "[APP] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *stdLogger) Debug(msg string, fields ...interface{}) {
	l.logger.Printf("[DEBUG] %s %v", msg, fields)
}

func (l *stdLogger) Info(msg string, fields ...interface{}) {
	l.logger.Printf("[INFO] %s %v", msg, fields)
}

func (l *stdLogger) Warn(msg string, fields ...interface{}) {
	l.logger.Printf("[WARN] %s %v", msg, fields)
}

func (l *stdLogger) Error(msg string, fields ...interface{}) {
	l.logger.Printf("[ERROR] %s %v", msg, fields)
}

func (l *stdLogger) Fatal(msg string, fields ...interface{}) {
	l.logger.Fatalf("[FATAL] %s %v", msg, fields)
}

// ==================== slog ====================

type slogLogger struct {
	logger *slog.Logger
}

func newSlogLogger(cfg *Config) *slogLogger {
	var output io.Writer = os.Stdout

	if cfg.OutputPath != "" {
		output = &lumberjack.Logger{
			Filename:   cfg.OutputPath,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		}
	}

	level := parseLevel(cfg.Level)
	handler := slog.NewJSONHandler(output, &slog.HandlerOptions{
		Level: level,
	})

	return &slogLogger{
		logger: slog.New(handler),
	}
}

func (l *slogLogger) Debug(msg string, fields ...interface{}) {
	l.logger.Debug(msg, fields...)
}

func (l *slogLogger) Info(msg string, fields ...interface{}) {
	l.logger.Info(msg, fields...)
}

func (l *slogLogger) Warn(msg string, fields ...interface{}) {
	l.logger.Warn(msg, fields...)
}

func (l *slogLogger) Error(msg string, fields ...interface{}) {
	l.logger.Error(msg, fields...)
}

func (l *slogLogger) Fatal(msg string, fields ...interface{}) {
	l.logger.Error(msg, fields...)
	os.Exit(1)
}

func parseLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// ==================== zap ====================

type zapLogger struct {
	logger *zap.Logger
}

func newZapLogger(cfg *Config) *zapLogger {
	var core zapcore.Core

	if cfg.OutputPath != "" {
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.OutputPath,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		})

		level := parseZapLevel(cfg.Level)
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			w,
			level,
		)
	} else {
		config := zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(parseZapLevel(cfg.Level))
		logger, _ := config.Build()
		return &zapLogger{logger: logger}
	}

	return &zapLogger{
		logger: zap.New(core),
	}
}

func (l *zapLogger) Debug(msg string, fields ...interface{}) {
	l.logger.Debug(msg, convertToZapFields(fields...)...)
}

func (l *zapLogger) Info(msg string, fields ...interface{}) {
	l.logger.Info(msg, convertToZapFields(fields...)...)
}

func (l *zapLogger) Warn(msg string, fields ...interface{}) {
	l.logger.Warn(msg, convertToZapFields(fields...)...)
}

func (l *zapLogger) Error(msg string, fields ...interface{}) {
	l.logger.Error(msg, convertToZapFields(fields...)...)
}

func (l *zapLogger) Fatal(msg string, fields ...interface{}) {
	l.logger.Fatal(msg, convertToZapFields(fields...)...)
}

func parseZapLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func convertToZapFields(fields ...interface{}) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields)/2)
	for i := 0; i < len(fields)-1; i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			continue
		}
		zapFields = append(zapFields, zap.Any(key, fields[i+1]))
	}
	return zapFields
}
