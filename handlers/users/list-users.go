package handlers

import "github.com/gin-gonic/gin"

func (u *UserHandler) List(c *gin.Context) {
	users, err := u.Services.List()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}
