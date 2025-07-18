package handlers

import (
	"clean_code/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *PedidosHandler) Create(c *gin.Context) {
	var pedido domain.PedidosRequest
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPedido, err := h.Services.Create(&pedido)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newPedido)
}
