package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

// Users
type User struct {
	gorm.Model
	Email string  `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"not null"`
  Name string
	Role string  `gorm:"size:255"`
	Todos []Todo
}
