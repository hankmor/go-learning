package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DemoCtrl struct {
}

func (*DemoCtrl) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "portal/index.html", gin.H{
		"title": "你好，hello！",
	})
}
