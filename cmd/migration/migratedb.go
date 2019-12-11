package main

import (
	"github.com/friday182/gql-api-server/internal/orm"
	"github.com/friday182/gql-api-server/internal/orm/models"
)

func main() {
	db := orm.ConnectDb()
	defer db.Close()

	db.AutoMigrate(
		&models.Todo{},
		&models.User{},
	)
}
