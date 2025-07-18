package handlers

import "github.com/gin-gonic/gin"

func (h *PedidosHandler) List(c *gin.Context) {
	pedidos, err := h.Services.List()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, pedidos)
}
