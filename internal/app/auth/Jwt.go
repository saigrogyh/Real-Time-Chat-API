package auth

import (
		"github.com/golang-jwt/jwt/v5"
		"time"
	)

var jwtSecret = []byte("JWT_SECRET")


func GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":	 time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}