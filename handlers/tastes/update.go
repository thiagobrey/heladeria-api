package tastes

import (
	"clean_code/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *TasteHandler) Update(c *gin.Context) {
	var taste domain.Taste
	if err := c.ShouldBindJSON(&taste); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"Invalid Id: ": err.Error()})
		return
	}
	taste.Id = id

	updatedTaste, err := h.Services.Update(&taste)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, updatedTaste)
}
