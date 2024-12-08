package crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []Todo
	DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}
