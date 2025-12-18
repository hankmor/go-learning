package main

import (
	"blog-api/internal/handler"
	"blog-api/internal/logger"
	"blog-api/internal/middleware"
	"blog-api/internal/model"
	"blog-api/internal/repository"
	"blog-api/internal/service"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// ==================== 0. 初始化日志 ====================
	// 可以通过环境变量或配置文件来选择日志类型
	loggerType := os.Getenv("LOGGER_TYPE") // log, slog, zap
	if loggerType == "" {
		loggerType = "slog" // 默认使用 slog
	}

	log := logger.New(&logger.Config{
		Type:       logger.LoggerType(loggerType),
		Level:      "info",
		OutputPath: "logs/app.log", // 输出到文件
		MaxSize:    100,             // 100MB
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   true,
	})

	log.Info("Starting application", "logger_type", loggerType)

	// ==================== 1. 初始化数据库 ====================
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", "error", err)
	}

	// 自动迁移数据库表结构
	db.AutoMigrate(&model.User{}, &model.Post{})
	log.Info("Database migrated successfully")

	// 配置连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大生命周期

	// ==================== 2. 依赖注入 ====================
	// Repository 层
	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)

	// Service 层
	authService := service.NewAuthService(userRepo)
	postService := service.NewPostService(postRepo)

	// Handler 层
	authHandler := handler.NewAuthHandler(authService)
	postHandler := handler.NewPostHandler(postService)

	log.Info("Dependencies injected successfully")

	// ==================== 3. 创建 Gin 引擎 ====================
	// 禁用 Gin 默认的日志，使用我们自己的日志中间件
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 使用 Recovery 中间件（panic 恢复）
	r.Use(gin.Recovery())

	// ==================== 4. 注册中间件 ====================
	// 日志中间件（记录所有 HTTP 请求）
	r.Use(middleware.LoggingMiddleware(log))

	// CORS 跨域配置
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 生产环境应该指定具体域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ==================== 5. 注册路由 ====================
	// 公开路由（不需要认证）
	api := r.Group("/api")
	{
		// 认证相关
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register) // 注册
			auth.POST("/login", authHandler.Login)       // 登录
		}

		// 文章查询（公开）
		api.GET("/posts", postHandler.List)    // 文章列表
		api.GET("/posts/:id", postHandler.Get) // 文章详情
	}

	// 需要认证的路由
	authorized := api.Group("")
	authorized.Use(middleware.AuthMiddleware()) // 使用 JWT 认证中间件
	{
		// 文章管理
		authorized.POST("/posts", postHandler.Create)       // 创建文章
		authorized.PUT("/posts/:id", postHandler.Update)    // 更新文章
		authorized.DELETE("/posts/:id", postHandler.Delete) // 删除文章
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	log.Info("Routes registered successfully")

	// ==================== 6. 启动服务 + 优雅关闭 ====================
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 在 Goroutine 中启动服务
	go func() {
		log.Info("Server started", "address", "http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start", "error", err)
		}
	}()

	// 监听系统信号
	quit := make(chan os.Signal, 1)
	// SIGINT: Ctrl+C
	// SIGTERM: kill 命令（Docker/K8s 停止容器时发送的信号）
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞，直到收到信号

	log.Info("Shutting down server...")

	// 设置 5 秒的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭：等待所有请求处理完毕，或超时
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown", "error", err)
	}

	log.Info("Server exited")
}
