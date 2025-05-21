package service

import "github.com/Cyuhsuan/stray_map_back_end/internal/models"

// UserService 定義用戶服務接口
type StrayMapService interface {
	// 用戶相關操作
	CreateStrayMap(strayMap *CreateStrayMapRequest) error
	GetStrayMapList() ([]models.StrayMap, error)
	GetStrayMapDetail(id uint) (*models.StrayMap, error)
	UpdateStrayMap(id uint, strayMap *UpdateStrayMapRequest) error
	DeleteStrayMap(id uint) error
}

type CreateStrayMapRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
}

type UpdateStrayMapRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

// UserServiceProvider 用於獲取 UserService 實例
var StrayMapServiceProvider StrayMapService

// InitStrayMapService 初始化用戶服務
func InitStrayMapService(useMock bool) {
	if useMock {
		StrayMapServiceProvider = NewMockStrayMapService()
	} else {
		// UserServiceProvider = NewDBUserService()
	}
}
