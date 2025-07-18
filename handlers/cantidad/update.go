package cantidad

import (
	"clean_code/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *CantidadHandler) Update(c *gin.Context) {
	var cantidad domain.CantidadRequest

	if err := c.ShouldBindJSON(&cantidad); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updatedCantidad, err := h.CantidadService.Update(&cantidad, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, updatedCantidad)
}