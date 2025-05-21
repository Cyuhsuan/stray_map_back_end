package main

import (
	"log"

	"github.com/Cyuhsuan/stray_map_back_end/api/handlers"
	"github.com/Cyuhsuan/stray_map_back_end/internal/config"
	"github.com/Cyuhsuan/stray_map_back_end/internal/middleware"
	"github.com/Cyuhsuan/stray_map_back_end/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// 加載 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// 加載配置
	config.LoadConfig()

	// 初始化服務
	service.InitUserService(config.AppConfig.UseMock)
}

func main() {
	router := gin.Default()

	// 公開路由
	router.POST("/api/auth/register", handlers.Register)
	router.POST("/api/auth/login", handlers.Login)

	// 需要認證的路由組
	authorized := router.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/user/profile", handlers.GetProfile)
	}

	port := ":" + config.AppConfig.Port
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
