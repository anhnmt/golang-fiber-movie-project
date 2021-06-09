package database

import (
	"github.com/xdorro/golang-fiber-base-project/models"
	"gorm.io/gorm"
	"log"
)

// Migrate updates the db with new columns, and tables
func Migrate(database *gorm.DB) {
	err := database.AutoMigrate(
		models.User{},
	)

	if err != nil {
		log.Println(err)
	}
}
