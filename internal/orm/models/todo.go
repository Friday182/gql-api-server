package models

type Todo struct {
	ID     uint `gorm:"primary_key"`
	TextA   string
	Done   bool
	TestAgain   string
}