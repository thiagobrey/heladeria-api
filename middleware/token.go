package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("JWT_KEY")

func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(), 
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
