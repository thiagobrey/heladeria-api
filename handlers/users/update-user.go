package handlers

import (
	"clean_code/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (u *UserHandler) Update(c *gin.Context) {
	var user domain.UserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := u.Services.Update(&user, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, updatedUser)
}
