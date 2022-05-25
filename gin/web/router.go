package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	var demoCtrl DemoCtrl
	var userCtrl UserCtrl

	fmt.Printf("%T\n", userCtrl)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})
	r.GET("/index", demoCtrl.Index)

	r.GET("/user/home", userCtrl.Home)
	r.POST("/user", userCtrl.Add)
	r.PUT("/user", userCtrl.Update)
	r.DELETE("/user", userCtrl.Delete)
	r.GET("/user", userCtrl.GetOne)
	r.GET("/users", userCtrl.GetAll)
	return r
}
