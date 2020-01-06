package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/friday182/gql-api-server/internal/orm"
	"github.com/friday182/gql-api-server/internal/orm/models"
)

func main() {
	db := orm.ConnectDb()
	defer db.Close()

	// db.CreateTable(&models.Todo{})
	// db.CreateTable(&models.User{})

	db.AutoMigrate(
		&models.Todo{},
		&models.User{},
	)

	testDb(db)
}

func testDb(db *gorm.DB) {
	todo := models.Todo{Text: "test12345", Done: false}
	db.Create(&todo)

	//var existUser []models.User
	//db.Find(&existUser)
	//fmt.Println("exist todos: ", existUser[0].Todos[0].Text)

	user := models.User{Email: "1@gmail.com", Password: "1234"}
	var tmp = []models.Todo{todo}
	user.Todos = tmp
	db.Create(&user)

	item := models.Todo{}
	db.Last(&item)
	fmt.Println("read last obj is: ", item.Text)

	var items []models.Todo
	db.Find(&items)
	fmt.Println("total number of objs: ", len(items))
}
