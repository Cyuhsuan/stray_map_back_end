package mock

import (
	"stray_map_back_end/internal/models"
)

var (
	// Users 模擬的用戶數據
	StrayMaps = []models.StrayMap{
		{
			ID:          1,
			UserID:      1,
			Title:       "測試地圖",
			Description: "測試地圖描述",
			Location:    "台北市",
		},
		{
			ID:          2,
			UserID:      2,
			Title:       "測試地圖2",
			Description: "測試地圖描述2",
			Location:    "台北市",
		},
	}
)
