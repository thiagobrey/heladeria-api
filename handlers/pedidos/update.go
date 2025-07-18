package handlers

import (
	"clean_code/internal/domain"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *PedidosHandler) Update(c *gin.Context) {
	var pedido domain.Pedidos
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idstr := c.Param("ID")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	uid := uint(id)
	pedido.ID = uid

	fmt.Println(pedido)
	editedPedido, err := h.Services.Update(&pedido)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, editedPedido)
}
