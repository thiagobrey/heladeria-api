package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("clave_secreta_super_segura")

func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(), // Expira en 1 hora
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
