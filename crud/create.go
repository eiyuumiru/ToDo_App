package crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo Todo

	todo.Title = c.Query("title")
	todo.Description = c.Query("description")
	status := c.Query("status")
	if status == "true" {
		todo.Status = true
	} else {
		todo.Status = false
	}

	DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}
