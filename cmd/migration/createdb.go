package main

import (
	"github.com/friday182/gql-api-server/internal/orm"
	"github.com/friday182/gql-api-server/internal/orm/models"
)

func main() {
	db := orm.ConnectDb()
	defer db.Close()

	tables := []interface{}{
		&models.Todo{},
		&models.User{},
	}
	
	db.DropTableIfExists(tables...)
	db.CreateTable(tables...)
}

