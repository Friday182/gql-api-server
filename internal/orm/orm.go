package orm

import (
	"log"

	"github.com/friday182/gql-api-server/internal/orm/models"
	"github.com/jinzhu/gorm"
)

func InitOrm() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "../../gorm.db")
	if err != nil {
		log.Panic("[ORM] err: ", err)
	}

	db.LogMode(true)
	err = RunMigration(&db)

	log.Println("[ORM] Database is ready.")

	return orm, err
}

func RunMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
	).Error
}
