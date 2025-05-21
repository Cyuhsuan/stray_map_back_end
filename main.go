package main

import (
	"log"
	"os"

	"github.com/Cyuhsuan/stray_map_back_end/api/handlers"
	"github.com/Cyuhsuan/stray_map_back_end/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var port = ":9000"

func init() {
	_ = godotenv.Load()
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = ":" + envPort
	}
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
		// 這裡添加需要認證的路由
		authorized.GET("/user/profile", handlers.GetProfile) // TODO: 實現這個處理器
	}

	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
