package server

import (
	"gin/web"
	"github.com/gin-gonic/gin"
)

func Start() {
	gin.DisableConsoleColor()
	gin.SetMode("debug")

	r := web.InitRouter()
	// 按目录加载模板
	r.LoadHTMLGlob("templates/**/*")
	// 按文件加载模板
	// r.LoadHTMLFiles("templates/index.html")

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
