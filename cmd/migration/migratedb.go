package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"encoding/json"
	//"strings"
	"github.com/friday182/gql-api-server/internal/orm"
	"github.com/friday182/gql-api-server/internal/orm/models"
)

func main() {
	db := orm.ConnectDb()
	defer db.Close()

	// db.CreateTable(&models.Todo{})
	// db.CreateTable(&models.User{})
/*
	db.AutoMigrate(
		&models.Todo{},
		&models.Test{},
		&models.User{},
	)
*/
	testDb(db)
}

func testDb(db *gorm.DB) {
	//todo := models.Todo{TextA: "test12", Done: false}
	//db.Create(&todo)

	todos := []models.Todo {
		models.Todo{TextA: "test12345", Done: false},
		models.Todo{TextA: "test12345", Done: false},
		models.Todo{TextA: "test12345", Done: false},
	}
/*
	tests := `[
		{"first":"one"},
		{"Second_k":2},
	]`*/
	mytest,_ := json.Marshal(todos)
	test := models.Test{Adata: 1, Bdata: "one"}
	user := models.User {
		Email:    "1@gmail.com",
		Password: "123456",
		Name:     "Lucas",
		Role:     "Admin",
		KpLog:    "na",
		Mytodo:   &test,
		Tests:    mytest,
	}
	//var existUser []models.User
	//db.Find(&existUser)
	//fmt.Println("exist todos: ", existUser[0].Todos[0].Text)

	//user := models.User{Email: "1@gmail.com", Password: "1234"}
	//var tmp = []models.Todo{todo}
	//user.Todos = tmp
	err := db.Create(&user).Error
	if err != nil {
		fmt.Println("create user failed: ", err)
	}

	//item := models.Todo{}
	//db.Last(&item)
	//fmt.Println("read last obj is: ", item.TextA)

	var items []models.User
	db.Find(&items)
	fmt.Println("total number of objs: ", len(items))
	check := []models.Todo{}
	if e := json.Unmarshal(items[2].Tests, &check); e == nil {
		fmt.Println("successful", check[0].TextA)
	}
}
