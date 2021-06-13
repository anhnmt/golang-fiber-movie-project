package database

import (
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"gorm.io/gorm"
	"log"
)

// Migrate updates the db with new columns, and tables
func Migrate(database *gorm.DB) {
	err := database.AutoMigrate(
		model.User{},
		model.Tag{},
		model.Genre{},
		model.Country{},
		model.People{},
		model.Movie{},
		model.MovieTag{},
		model.MovieGenre{},
		model.MovieCountry{},
		model.MoviePeople{},
	)

	if err != nil {
		log.Println(err)
	}
}
