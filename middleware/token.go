package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
            return
        }
        token := strings.TrimPrefix(authHeader, "Bearer ")
		fmt.Println("Token recibido:", token)
        /* if token != tokenSecret {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
            return
        }
		*/
        c.Next()
    }
}