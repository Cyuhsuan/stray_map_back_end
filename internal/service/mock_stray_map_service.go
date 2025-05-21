package service

import (
	"errors"

	"github.com/Cyuhsuan/stray_map_back_end/internal/mock"
	"github.com/Cyuhsuan/stray_map_back_end/internal/models"
)

type mockStrayMapService struct {
	strayMaps []models.StrayMap
}

// NewMockStrayMapService 創建一個新的 mock 用戶服務實例
func NewMockStrayMapService() StrayMapService {
	return &mockStrayMapService{
		strayMaps: mock.StrayMaps,
	}
}

func (s *mockStrayMapService) CreateStrayMap(strayMap *CreateStrayMapRequest) error {
	s.strayMaps = append(s.strayMaps, models.StrayMap{
		ID:          uint(len(s.strayMaps) + 1),
		UserID:      1, // TODO: 需要從 context 中獲取當前登入用戶的 ID
		Title:       strayMap.Title,
		Description: strayMap.Description,
		Location:    strayMap.Location,
	})
	return nil
}

func (s *mockStrayMapService) GetStrayMapList() ([]models.StrayMap, error) {
	return s.strayMaps, nil
}

func (s *mockStrayMapService) GetStrayMapDetail(id uint) (*models.StrayMap, error) {
	for _, strayMap := range s.strayMaps {
		if strayMap.ID == id {
			return &strayMap, nil
		}
	}
	return nil, errors.New("stray map not found")
}

func (s *mockStrayMapService) UpdateStrayMap(id uint, strayMap *UpdateStrayMapRequest) error {
	for i, existingMap := range s.strayMaps {
		if existingMap.ID == id {
			s.strayMaps[i] = models.StrayMap{
				ID:          id,
				UserID:      existingMap.UserID,
				Title:       strayMap.Title,
				Description: strayMap.Description,
				Location:    strayMap.Location,
			}
			return nil
		}
	}
	return errors.New("stray map not found")
}

func (s *mockStrayMapService) DeleteStrayMap(id uint) error {
	for i, strayMap := range s.strayMaps {
		if strayMap.ID == id {
			s.strayMaps = append(s.strayMaps[:i], s.strayMaps[i+1:]...)
			return nil
		}
	}
	return errors.New("stray map not found")
}
