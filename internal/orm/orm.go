package orm

import (
	"log"

	"github.com/jinzhu/gorm"
)

<<<<<<< HEAD
func ConnectDb() *gorm.DB {
	env := "debug"
	db, err := gorm.Open("sqlite3", "./gorm.db")
=======
func InitOrm() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "../../gorm.db")
>>>>>>> a90de4222185d7b88e1a4b4ee6f5070fa2251e5b
	if err != nil {
		panic("[ORM] Connect Database Failed ")
	}

<<<<<<< HEAD
	if env != "producton" {
		db.LogMode(true)
	}
	log.Println("[ORM] Database is ready.")
=======
	db.LogMode(true)
	err = RunMigration(&db)

	log.Println("[ORM] Database is ready.")

	return orm, err
}
>>>>>>> a90de4222185d7b88e1a4b4ee6f5070fa2251e5b

	return db
}
