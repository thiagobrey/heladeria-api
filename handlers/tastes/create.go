package tastes

import (
	"clean_code/internal/domain"

	"github.com/gin-gonic/gin"
)

func (t *TasteHandler) Create(c *gin.Context) {
	var taste domain.Taste

	if err := c.ShouldBindJSON(&taste); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newTaste, err := t.Services.Create(&taste)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, newTaste)
}
