package mockdb

import "fmt"

// ExampleUser 是一個範例結構
type ExampleUser struct {
	ID       string
	Name     string
	Email    string
	Password string
}

// users 是一個模擬的資料表
var users = []ExampleUser{
	{ID: "1", Name: "Alice", Email: "alice@example.com", Password: "password123"},
	{ID: "2", Name: "Bob", Email: "bob@example.com", Password: "password456"},
}

// GetAllUsers 回傳所有 mock 使用者
func GetAllUsers() []ExampleUser {
	return users
}

func GetUserByEmail(email string) (ExampleUser, error) {
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}
	return ExampleUser{}, fmt.Errorf("user not found")
}
