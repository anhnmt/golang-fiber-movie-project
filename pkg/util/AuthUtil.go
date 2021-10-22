package util

import (
	"strconv"

	"github.com/form3tech-oss/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HashPassword hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ValidToken(token *jwt.Token, id string) bool {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	if uid != userId {
		return false
	}

	return true
}
