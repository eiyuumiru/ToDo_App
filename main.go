package main

import (
	"todo-app/crud"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() {
	var err error
	crud.DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	crud.DB.AutoMigrate(&crud.Todo{})
}

func main() {
	InitDatabase()
	router := gin.Default()

	router.POST("/todos", crud.CreateTodo)
	router.GET("/todos", crud.GetTodos)
	router.PUT("/todos/:id", crud.UpdateTodo)
	router.DELETE("/todos/:id", crud.DeleteTodo)

	router.Run(":8080")
}
