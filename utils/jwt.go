package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)


var jwtSecret = []byte("ULASIN_SECRET_KEY")

func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role": role,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}