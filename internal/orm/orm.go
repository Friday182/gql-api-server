package orm

import (
	"log"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func ConnectDb() *gorm.DB {
	env := "debug"
	db, err := gorm.Open("sqlite3", "../../internal/orm/gorm.db")
	// db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=ttm password=123456 sslmode=disable")
	if err != nil {
		panic("[ORM] Connect Database Failed ")
	}

	if env != "producton" {
		db.LogMode(true)
	}
	log.Println("[ORM] Database is ready.")

	return db
}
