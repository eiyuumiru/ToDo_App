package crud

import (
	"gorm.io/gorm"
)

type Todo struct {
	ID          uint   `json:"id" gorm:"primaryKey" form:"id" binding:"-"`
	Title       string `json:"title" form:"title" binding:"-"`
	Description string `json:"description" form:"description" binding:"-"`
	Status      bool   `json:"status" form:"completed" binding:"-"`
}

var DB *gorm.DB
