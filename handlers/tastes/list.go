package tastes

import (
	"github.com/gin-gonic/gin"
)

func (h *TasteHandler) List(c *gin.Context) {
	tastes, err := h.Services.List()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tastes)
}
