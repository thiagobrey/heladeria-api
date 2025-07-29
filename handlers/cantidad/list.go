package cantidad

import "github.com/gin-gonic/gin"

func (h *CantidadHandler) List(c *gin.Context) {
	cantidades, err := h.CantidadService.List()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, cantidades)
}
