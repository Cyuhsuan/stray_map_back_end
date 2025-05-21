package mock

import (
	"github.com/Cyuhsuan/stray_map_back_end/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	// Users 模擬的用戶數據
	Users = []models.User{
		{
			ID:       1,
			Username: "test_user1",
			Password: generateHashPassword("password123"),
			Email:    "test1@example.com",
		},
		{
			ID:       2,
			Username: "test_user2",
			Password: generateHashPassword("password456"),
			Email:    "test2@example.com",
		},
	}
)

// FindUserByUsername 根據用戶名查找用戶
func FindUserByUsername(username string) *models.User {
	for _, user := range Users {
		if user.Username == username {
			return &user
		}
	}
	return nil
}

// FindUserByEmail 根據郵箱查找用戶
func FindUserByEmail(email string) *models.User {
	for _, user := range Users {
		if user.Email == email {
			return &user
		}
	}
	return nil
}

// generateHashPassword 生成密碼哈希
func generateHashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
