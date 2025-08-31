package admins

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *AdminHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"Invalid Id": err.Error()})
		return
	}

	admin, err := h.AdminService.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, admin)
}
