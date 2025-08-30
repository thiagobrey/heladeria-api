package admins

import (
	"clean_code/internal/domain"
	"clean_code/middleware"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *AdminHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin, err := h.AdminService.GetByEmail(req.Email)
	if err != nil || admin.Password != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	token, err := middleware.GenerateToken(admin.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}
	fmt.Println(token)

	sesion := domain.Sesions{
		Token: token,
		Timestamps: domain.Timestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Data: admin.Email,
	}

	_, err = h.SesionService.Create(&sesion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la sesión"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
