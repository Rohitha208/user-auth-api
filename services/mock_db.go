package services

import (
	"errors"
	"user-auth-api/src/models"
)

type MockDB struct {
	users map[string]models.User
}

func NewMockDB() *MockDB {
	return &MockDB{users: make(map[string]models.User)}
}

func (db *MockDB) SaveUser(user models.User) error {
	if _, exists := db.users[user.Username]; exists {
		return errors.New("user already exists")
	}
	db.users[user.Username] = user
	return nil
}

func (db *MockDB) GetUserByUsername(username string) (models.User, error) {
	user, exists := db.users[username]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}
