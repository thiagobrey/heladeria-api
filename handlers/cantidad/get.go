package cantidad

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *CantidadHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	cantidad, err := h.CantidadService.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, cantidad)
}
