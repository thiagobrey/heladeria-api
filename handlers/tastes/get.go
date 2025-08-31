package tastes

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *TasteHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"Invalid Id: ": err.Error()})
		return
	}

	taste, err := h.Services.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, taste)
}
