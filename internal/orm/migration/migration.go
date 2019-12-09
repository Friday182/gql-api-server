package migration

import (
	"fmt"
	log "log"
	"github.com/jinzhu/gorm"
	"github.com/friday182/gql-api-server/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

func RunMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
	).Error
}