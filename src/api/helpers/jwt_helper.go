package helpers

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go" // Make sure to import the jwt package
)

// Secret key for signing tokens (ensure to keep this secret)
var jwtSecret = []byte("your_secret_key") // Change this to your actual secret key

// GenerateToken generates a JWT token for the given username
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

// ValidateToken checks the validity of the token
func ValidateToken(tokenString string) (string, error) {
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims.Subject, nil
}
