package service

import (
	"errors"

	"github.com/Cyuhsuan/stray_map_back_end/internal/mock"
	"github.com/Cyuhsuan/stray_map_back_end/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type mockUserService struct {
	users []models.User
}

// NewMockUserService 創建一個新的 mock 用戶服務實例
func NewMockUserService() UserService {
	return &mockUserService{
		users: mock.Users,
	}
}

func (s *mockUserService) CreateUser(username, password, email string) (*models.User, error) {
	// 檢查用戶名是否已存在
	if _, err := s.GetUserByUsername(username); err == nil {
		return nil, errors.New("username already exists")
	}

	// 檢查郵箱是否已存在
	if _, err := s.GetUserByEmail(email); err == nil {
		return nil, errors.New("email already exists")
	}

	// 創建新用戶
	hashedPassword := generateHashPassword(password)
	newUser := &models.User{
		ID:       uint(len(s.users) + 1),
		Username: username,
		Password: hashedPassword,
		Email:    email,
	}

	s.users = append(s.users, *newUser)
	return newUser, nil
}

func (s *mockUserService) GetUserByUsername(username string) (*models.User, error) {
	for _, user := range s.users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *mockUserService) GetUserByEmail(email string) (*models.User, error) {
	for _, user := range s.users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *mockUserService) GetUserByID(id uint) (*models.User, error) {
	for _, user := range s.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *mockUserService) ValidatePassword(user *models.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// generateHashPassword 生成密碼哈希
func generateHashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
