package services

import (
	models "TechHunterClone/src/models/user"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

const secretKey = "your-secret-key"

func CreateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"email":   user.Email,
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 48).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexp signing method")
		}
		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := &models.User{
			Email: claims["email"].(string),
			ID:    uint(claims["user_id"].(float64)),
			Role:  models.Role(claims["role"].(string)),
		}
		return user, nil
	} else {
		return nil, err
	}
}
func ValidateUser(user *models.User) error {
	if len(user.FirstName) == 0 {
		return errors.New("first name is required")
	}

	if len(user.LastName) == 0 {
		return errors.New("last name is required")
	}

	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	return nil
}

func isValidEmail(email string) bool {
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return strings.Contains(email, "@")
}
