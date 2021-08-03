package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/controller"
)

func publicRoute(a fiber.Router) {
	clients := a.Group("/clients")

	// Countries Controller
	countryController := controller.NewCountryController()
	clients.Get("/countries", countryController.ClientFindAllCountries)

	// Genres Controller
	genreController := controller.NewGenreController()
	clients.Get("/genres", genreController.ClientFindAllGenres)

	// Movie Controller
	movieController := controller.NewMovieController()
	clients.Get("/top-movies-sidebar", movieController.ClientTopMovieSidebar)
	clients.Get("/top-movies-body", movieController.ClientTopMoviesBody)
	clients.Get("/find-movie-detail/:movieSlug", movieController.ClientFindMovieDetail)
}
