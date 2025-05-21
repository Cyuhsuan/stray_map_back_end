package handlers

import (
	"net/http"

	"github.com/Cyuhsuan/stray_map_back_end/internal/auth"
	"github.com/Cyuhsuan/stray_map_back_end/internal/mock"
	"github.com/Cyuhsuan/stray_map_back_end/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// Register 處理用戶註冊
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 檢查用戶名是否已存在
	if existingUser := mock.FindUserByUsername(req.Username); existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
		return
	}

	// 檢查郵箱是否已存在
	if existingUser := mock.FindUserByEmail(req.Email); existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		return
	}

	// 加密密碼
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	// 創建新用戶
	newUser := models.User{
		ID:       uint(len(mock.Users) + 1), // 簡單的 ID 生成
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
	}

	// 在實際應用中，這裡會保存到數據庫
	mock.Users = append(mock.Users, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"user": gin.H{
			"id":       newUser.ID,
			"username": newUser.Username,
			"email":    newUser.Email,
		},
	})
}

// Login 處理用戶登入
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查找用戶
	user := mock.FindUserByUsername(req.Username)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// 驗證密碼
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// 生成 JWT token
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

// GetProfile 獲取用戶資料
func GetProfile(c *gin.Context) {
	userID, _ := c.Get("userID")

	// 在實際應用中，這裡會從數據庫查詢用戶資料
	var foundUser *models.User
	for _, user := range mock.Users {
		if user.ID == userID.(uint) {
			foundUser = &user
			break
		}
	}

	if foundUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":       foundUser.ID,
			"username": foundUser.Username,
			"email":    foundUser.Email,
		},
	})
}
