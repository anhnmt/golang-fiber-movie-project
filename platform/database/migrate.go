package database

import (
	models2 "github.com/xdorro/golang-fiber-base-project/app/models"
	"gorm.io/gorm"
	"log"
)

// Migrate updates the db with new columns, and tables
func Migrate(database *gorm.DB) {
	err := database.AutoMigrate(
		models2.User{},
		models2.Tag{},
		models2.Genre{},
		models2.Country{},
		models2.People{},
		models2.Movie{},
		models2.MovieTag{},
		models2.MovieGenre{},
		models2.MovieCountry{},
		models2.MoviePeople{},
	)

	if err != nil {
		log.Println(err)
	}
}
