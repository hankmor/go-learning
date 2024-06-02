package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

var inErr bool

func main() {
	port := os.Args[1]
	node := os.Args[2]
	r := gin.Default()
	go func() {
		for {
			if node == "node1" {
				break
			}
			inErr = !inErr
			time.Sleep(3 * time.Second)
		}
	}()
	r.GET("/", func(ctx *gin.Context) {
		if inErr {
			ctx.String(http.StatusInternalServerError, "error: "+node)
		} else {
			ctx.JSON(http.StatusOK, gin.H{"msg": "ok, " + node})
		}
	})
	r.GET("/err", func(ctx *gin.Context) {
		ctx.String(http.StatusInternalServerError, "error: "+node)
		ctx.Abort()
	})
	r.GET("/timeout", func(ctx *gin.Context) {
		time.Sleep(time.Second * 10)
		ctx.String(http.StatusOK, "after 10s responsed")
	})
	_ = r.Run(":" + port) // 参数0为执行文件本身信息，真正的参数下标为1
}
