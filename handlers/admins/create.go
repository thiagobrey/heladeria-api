package admins

import (
	"clean_code/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h *AdminHandler) Create(c *gin.Context) {
	var admin domain.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	createdAdmin, err := h.AdminService.Create(&admin)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, createdAdmin)
}
