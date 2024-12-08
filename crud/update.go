package crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateTodo(c *gin.Context) {
	var todo Todo
	id := c.Param("id")

	if err := DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "To-Do not found"})
		return
	}

	// Read data from query parameters
	todo.Title = c.Query("title")
	todo.Description = c.Query("description")
	status := c.Query("status")
	if status == "true" {
		todo.Status = true
	} else {
		todo.Status = false
	}

	DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}
