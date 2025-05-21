package main

import (
	"log"

	"stray_map_back_end/api/handlers"
	"stray_map_back_end/internal/config"
	"stray_map_back_end/internal/middleware"
	"stray_map_back_end/internal/service"

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
	service.InitStrayMapService(config.AppConfig.UseMock)
}

func main() {
	router := gin.Default()

	// 公開路由
	router.POST("/api/auth/register", handlers.Register)
	router.POST("/api/auth/login", handlers.Login)
	router.GET("/api/stray_map", handlers.GetStrayMapList)

	// 需要認證的路由組
	authorized := router.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/user/profile", handlers.GetProfile)
		// stray map
		authorized.POST("/stray_map", handlers.CreateStrayMap)
		authorized.PUT("/stray_map/:id", handlers.UpdateStrayMap)
		authorized.DELETE("/stray_map/:id", handlers.DeleteStrayMap)
	}

	port := ":" + config.AppConfig.Port
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
