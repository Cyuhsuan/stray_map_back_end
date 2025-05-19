package main

import (
	"os"

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
	RegisterRoutes(router)
	router.Run(port)
}
