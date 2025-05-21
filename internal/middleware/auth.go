package middleware

import (
	"net/http"
	"strings"

	"github.com/Cyuhsuan/stray_map_back_end/internal/auth"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 驗證 JWT token 的中間件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			c.Abort()
			return
		}

		// 檢查 Bearer token 格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}

		// 驗證 token
		claims, err := auth.ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// 將用戶 ID 存儲在上下文中
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
