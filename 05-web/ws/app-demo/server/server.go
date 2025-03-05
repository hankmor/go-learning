package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var upgrader = websocket.Upgrader{
	HandshakeTimeout: time.Second * 10,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Trade 数据模型
type Trade struct {
	gorm.Model
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price" gorm:"type:decimal(10,2)"`
	Quantity  float64 `json:"quantity" gorm:"type:decimal(10,4)"`
	Timestamp string  `json:"timestamp"`
}

func main() {
	// 初始化数据库
	db, err := gorm.Open(sqlite.Open("trades.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Trade{})

	r := gin.Default()

	// 静态文件服务
	r.Static("/static", "./client")

	// WebSocket 路由
	r.GET("/ws", func(c *gin.Context) {
		handleWebSocket(c, db)
	})
	// 查询所有交易
	r.GET("/trades", func(c *gin.Context) {
		var trades []Trade
		db.Order("created_at desc").Limit(50).Find(&trades)
		c.JSON(200, trades)
	})
	// 启动服务
	fmt.Println("server started at: ", 8080)
	r.Run(":8080")
}

func handleWebSocket(c *gin.Context, db *gorm.DB) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("websocket upgrade error:", err)
		return
	}
	defer conn.Close()

	// 生成实时交易数据
	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()
	for range ticker.C {
		trade := generateTrade()
		db.Create(&trade)
		msg, _ := json.Marshal(trade)
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			fmt.Println("write error: ", err)
			return
		}
	}
}

func generateTrade() Trade {
	return Trade{
		Symbol:    "RMB/USD",
		Price:     7.00 + rand.Float64()*1,
		Quantity:  rand.Float64() * 500,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
