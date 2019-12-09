package orm

import (
	"log"

	"github.com/friday182/gql-api-server/internal/orm/models"
	"github.com/jinzhu/gorm"
)

func InitOrm() (*ORM, error) {
	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		log.Panic("[ORM] err: ", err)
	}

	// Set log mode
	db.LogMode(logMode)
	err = RunMigration(&db)
	log.Println("[ORM] Database connection initialized.")
	return orm, err
}

func RunMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
	).Error
}
