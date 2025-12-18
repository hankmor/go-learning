package response

import "github.com/gin-gonic/gin"

// Response 统一响应结构
// 所有 API 都应该返回这个格式，便于前端统一处理
type Response struct {
	Code    int         `json:"code"`    // 业务状态码：0 表示成功，其他表示各种错误
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 业务数据，可以是任意类型
}

// Success 成功响应
// 使用示例：response.Success(c, user)
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// Error 失败响应
// 使用示例：response.Error(c, 404, "用户不存在")
func Error(c *gin.Context, code int, message string) {
	c.JSON(200, Response{ // 注意：HTTP 状态码仍然是 200，业务错误通过 code 字段区分
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// SuccessWithMessage 成功响应（自定义消息）
func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(200, Response{
		Code:    0,
		Message: message,
		Data:    data,
	})
}
