package server

import (
	"github.com/gin-gonic/gin"
	"github.com/huzhouv/go-learning/gin/web"
)

func Start() {
	gin.DisableConsoleColor()
	gin.SetMode("debug")

	r := web.InitRouter()
	// 按目录加载模板
	r.LoadHTMLGlob("gin/templates/**/*")
	// 按文件加载模板
	// r.LoadHTMLFiles("templates/index.html")

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
