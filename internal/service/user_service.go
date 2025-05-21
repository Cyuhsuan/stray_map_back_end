package service

import "stray_map_back_end/internal/models"

// UserService 定義用戶服務接口
type UserService interface {
	// 用戶相關操作
	CreateUser(username, password, email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	ValidatePassword(user *models.User, password string) error
}

// UserServiceProvider 用於獲取 UserService 實例
var UserServiceProvider UserService

// InitUserService 初始化用戶服務
func InitUserService(useMock bool) {
	if useMock {
		UserServiceProvider = NewMockUserService()
	} else {
		// UserServiceProvider = NewDBUserService()
	}
}
