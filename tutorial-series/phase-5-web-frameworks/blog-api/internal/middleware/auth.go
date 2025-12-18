package middleware

import (
	"blog-api/utils/jwt"
	"blog-api/utils/response"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT 认证中间件
// 验证请求头中的 Authorization Token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从 Header 中获取 Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, 401, "缺少 Authorization Header")
			c.Abort() // 终止后续处理
			return
		}

		// 2. 去掉 "Bearer " 前缀
		// 标准格式：Authorization: Bearer <token>
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			// 说明没有 "Bearer " 前缀
			response.Error(c, 401, "Token 格式错误")
			c.Abort()
			return
		}

		// 3. 解析 Token
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			response.Error(c, 401, "无效的 Token")
			c.Abort()
			return
		}

		// 4. 将用户信息存入上下文，供后续 Handler 使用
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		// 5. 继续处理请求
		c.Next()
	}
}
