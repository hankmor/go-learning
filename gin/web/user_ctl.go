package web

import (
	"gin/db"
	"gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserCtrl struct {
	userService service.UserService
}

func (uc *UserCtrl) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "users/user.html", gin.H{
		"Age": 20, "Name": "Sam",
	})
}

func (uc *UserCtrl) Add(c *gin.Context) {
	user := &db.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if uc.userService.Add(user) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "SERVER ERROR"})
	}
}

func (uc *UserCtrl) Update(c *gin.Context) {
	user := &db.User{}
	if err := c.ShouldBindJSON(user); err != nil || user.Id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if uc.userService.Update(user) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "SERVER ERROR"})
	}
}

func (uc *UserCtrl) Delete(c *gin.Context) {
	id := uc.getQueryId(c)
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter"})
		return
	}
	if uc.userService.Delete(id) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "SERVER ERROR"})
	}
}

func (uc *UserCtrl) GetOne(c *gin.Context) {
	id := uc.getQueryId(c)
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter"})
		return
	}
	user := uc.userService.Get(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user, // 渲染为json
	})
}

func (uc *UserCtrl) getQueryId(c *gin.Context) int64 {
	idstr := c.Query("id")
	if idstr == "" {
		return 0
	}
	id, _ := strconv.ParseInt(idstr, 10, 64)
	return id
}

func (uc *UserCtrl) GetAll(c *gin.Context) {
	users := uc.userService.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    users, // 渲染为json
	})
}
