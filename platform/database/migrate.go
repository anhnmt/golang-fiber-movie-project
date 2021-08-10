package database

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"log"
)

// migrate : updates the db with new columns, and tables
func migrate() {
	if err = db.AutoMigrate(
		model.User{},
		model.Role{},

		//model.Tag{},
		model.Genre{},
		model.Country{},
		//model.People{},
		model.Banner{},

		model.Movie{},
		//model.MovieTag{},
		model.MovieGenre{},
		model.MovieCountry{},
		//model.MoviePeople{},
		model.MovieType{},

		model.Episode{},
		model.EpisodeType{},
		model.EpisodeDetail{},
	); err != nil {
		log.Println(err)
	}
}
