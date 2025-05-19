package model

import (
	"github.com/Cyuhsuan/stray_map_back_end/mockdb"
)

// 定義 interface
type UserStore interface {
	GetAllUsers() []mockdb.ExampleUser
	// 其他 CRUD 方法
}

// mockdb 實作 UserStore
type MockUserStore struct{}

func (m *MockUserStore) GetAllUsers() []mockdb.ExampleUser { return mockdb.GetAllUsers() }

func (m *MockUserStore) GetUserByEmail(email string) (mockdb.ExampleUser, error) {
	return mockdb.GetUserByEmail(email)
}

// 真實 db 也實作 UserStore
// type RealUserStore struct{}
// func (r *RealUserStore) GetAllUsers() []ExampleUser { /* 連資料庫查詢 */ }
