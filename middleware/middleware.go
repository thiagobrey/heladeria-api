package middleware

import (
    "clean_code/internal/domain"
    "fmt"
    "net/http"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {


        
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
            return
        }
        token := strings.TrimPrefix(authHeader, "Bearer ")

        var sesion domain.Sesions
        if err := db.Where("token = ?", token).First(&sesion).Error; err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
            return
        }

        if time.Since(sesion.CreatedAt) > time.Hour {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expirado"})
            return
        }

        fmt.Println("Token válido:", token)
        c.Next()
    }
}