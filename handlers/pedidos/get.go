package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *PedidosHandler) GetCommentByUserId(c *gin.Context) {
	userIdStr := c.Param("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pedido, err := h.Services.GetByUserId(userId)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, pedido)
}
