package services

import (
	"golang.org/x/crypto/bcrypt"
)

// AuthService provides authentication-related services.
type AuthService struct{}

// HashPassword hashes the given password using bcrypt.
func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword checks the given password against the hashed password.
func (s *AuthService) CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
