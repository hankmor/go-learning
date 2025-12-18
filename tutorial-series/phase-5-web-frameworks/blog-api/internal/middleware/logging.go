package middleware

import (
	"blog-api/internal/logger"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware 日志中间件
func LoggingMiddleware(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		duration := time.Since(start)

		// 记录请求日志
		log.Info("HTTP Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration_ms", duration.Milliseconds(),
			"client_ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
		)

		// 如果有错误，记录错误日志
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				log.Error("Request error",
					"error", err.Error(),
					"path", c.Request.URL.Path,
				)
			}
		}
	}
}
