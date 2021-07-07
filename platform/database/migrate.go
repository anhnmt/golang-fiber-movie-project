package database

import (
	model2 "github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"log"
)

// migrate : updates the db with new columns, and tables
func migrate() {
	if err = db.AutoMigrate(
		model2.User{},
		model2.Role{},
		model2.Permission{},
		model2.RolePermission{},
		model2.UserPermission{},

		model2.Tag{},
		model2.Genre{},
		model2.Country{},
		model2.People{},

		model2.Movie{},
		model2.MovieTag{},
		model2.MovieGenre{},
		model2.MovieCountry{},
		model2.MoviePeople{},
		model2.MovieType{},
	); err != nil {
		log.Println(err)
	}
}
