package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/controller"
)

func privateRoute(a *fiber.App) {
	api := a.Group("/api")

	// Tags Controller
	tags := api.Group("/tags")
	tags.Get("/", controller.FindAllTags)
	tags.Post("/", controller.CreateNewTag)
	tags.Get("/:id", controller.FindTagById)
	tags.Put("/:id", controller.UpdateTagById)
	tags.Delete("/:id", controller.DeleteTagById)

	// Genres Controller
	genres := api.Group("/genres")
	genres.Get("/", controller.FindAllGenres)
	genres.Post("/", controller.CreateNewGenre)
	genres.Get("/:id", controller.FindGenreById)
	genres.Put("/:id", controller.UpdateGenreById)
	genres.Delete("/:id", controller.DeleteGenreById)

	// Countries Controller
	countries := api.Group("/countries")
	countries.Get("/", controller.FindAllCountries)
	countries.Post("/", controller.CreateNewCountry)
	countries.Get("/:id", controller.FindCountryById)
	countries.Put("/:id", controller.UpdateCountryById)
	countries.Delete("/:id", controller.DeleteCountryById)

	// Users Controller
	users := api.Group("/users")
	users.Get("/", controller.FindAllUsers)
	users.Post("/", controller.CreateNewUser)
	users.Get("/:id", controller.FindUserById)
	users.Put("/:id", controller.UpdateUserById)
	users.Delete("/:id", controller.DeleteUserById)
}
