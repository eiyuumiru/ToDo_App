package crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(&Todo{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Todo"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "To-Do deleted"})
}
