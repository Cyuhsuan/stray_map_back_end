package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", rootHandler)
	router.GET("/users", getUsers)
	router.POST("/users", createUser)
	router.GET("/users/:id", getUserByID)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)
	router.POST("/login", loginHandler)

	auth := router.Group("/auth")
	auth.Use(JWTAuthMiddleware())
	auth.GET("/profile", profileHandler)
}

func rootHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GET all users"})
}

func createUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "POST create user"})
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "GET user", "id": id})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "PUT update user", "id": id})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "DELETE user", "id": id})
}
