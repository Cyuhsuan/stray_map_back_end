package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key") // 建議用環境變數管理

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

// JWT 產生
func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// JWT middleware
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Missing Authorization header"})
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}
		c.Set("user", token.Claims)
		c.Next()
	}
}

// login 範例
func loginHandler(c *gin.Context) {
	// 實際應驗證帳密，這裡僅示範
	userID := "123"
	token, err := GenerateJWT(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(200, gin.H{"token": token})
}

// 需 JWT 驗證
func profileHandler(c *gin.Context) {
	claims, _ := c.Get("user")
	c.JSON(200, gin.H{"user": claims})
}
