package handlers

import (
	"clean_code/internal/domain"

	"github.com/gin-gonic/gin"
)

func (u *UserHandler) Create(c *gin.Context) {
	var user domain.UserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := u.Services.Create(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, createdUser)
}
