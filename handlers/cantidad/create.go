package cantidad

import (
	"clean_code/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h *CantidadHandler) Create(c *gin.Context) {
	var cr domain.CantidadRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	cantidad, err := h.CantidadService.Create(&cr)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, cantidad)
}
