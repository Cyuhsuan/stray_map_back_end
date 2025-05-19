package main

import (
	"time"

	"github.com/Cyuhsuan/stray_map_back_end/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key") // 建議用環境變數管理

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
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	userStore := model.MockUserStore{}
	user, err := userStore.GetUserByEmail(loginData.Email)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	if user.Password != loginData.Password {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	userID := user.ID
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
	userID := claims.(jwt.MapClaims)["user_id"].(string)
	c.JSON(200, gin.H{"user_id": userID})
}
