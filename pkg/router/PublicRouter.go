package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/controller"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
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
	clients.Post("/test-upload", func(c *fiber.Ctx) error {
		movie := c.FormValue("movie")

		movieRequest := new(request.MovieRequest)

		if err := util.JSONUnmarshal([]byte(movie), movieRequest); err != nil {
			return util.ResponseError(err.Error(), nil)
		}

		// Get first file from form field "poster":
		file, err := c.FormFile("poster")

		if file != nil {
			// 👷 Save file inside uploads folder under current working directory:
			poster := util.StoragePoster(file.Filename)

			if err = c.SaveFile(file, util.Storage(poster)); err != nil {
				return err
			}

			movieRequest.Poster = poster
		}

		return util.ResponseSuccess("Thành công", movieRequest)
	})
}
