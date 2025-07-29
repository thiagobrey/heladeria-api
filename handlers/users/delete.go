package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (u *UserHandler) Delete(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := u.Services.Delete(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}
