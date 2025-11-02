package auth

import (
		"github.com/golang-jwt/jwt/v5"
		"time"
	)


func GenerateToken(userId uint, secret string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userId,
        "exp":	 time.Now().Add(72 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString([]byte(secret)) // ใช้ secret ที่รับมา
}