package services

import "user-auth-api/src/api/helpers"

type LoginService struct {
	// Add necessary fields, such as a user repository
}

// NewLoginService creates a new instance of LoginService
func NewLoginService() *LoginService {
	return &LoginService{}
}

func (s *LoginService) Login(username, password string) (string, error) {
	// Your login logic
	token, err := helpers.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}
