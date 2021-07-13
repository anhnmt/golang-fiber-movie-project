package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/controller"
)

func privateRoute(a fiber.Router) {

	// Tags Controller
	tags := a.Group("/tags")
	tags.Get("/", controller.FindAllTags)
	tags.Post("/", controller.CreateNewTag)
	tags.Get("/:id", controller.FindTagById)
	tags.Put("/:id", controller.UpdateTagById)
	tags.Delete("/:id", controller.DeleteTagById)

	// Genres Controller
	genres := a.Group("/genres")
	genres.Get("/", controller.FindAllGenres)
	genres.Post("/", controller.CreateNewGenre)
	genres.Get("/:id", controller.FindGenreById)
	genres.Put("/:id", controller.UpdateGenreById)
	genres.Delete("/:id", controller.DeleteGenreById)

	// Countries Controller
	countries := a.Group("/countries")
	countries.Get("/", controller.FindAllCountries)
	countries.Post("/", controller.CreateNewCountry)
	countries.Get("/:id", controller.FindCountryById)
	countries.Put("/:id", controller.UpdateCountryById)
	countries.Delete("/:id", controller.DeleteCountryById)

	// Users Controller
	users := a.Group("/users")
	users.Get("/", controller.FindAllUsers)
	users.Post("/", controller.CreateNewUser)
	users.Get("/:id", controller.FindUserById)
	users.Put("/:id", controller.UpdateUserById)
	users.Delete("/:id", controller.DeleteUserById)

	// UserRoles Controller
	//users.Get("/:id/roles", controller.FindAllUserRoles)

	// Permissions Controller
	permissions := a.Group("/permissions")
	permissions.Get("/", controller.FindAllPermissions)
	permissions.Post("/", controller.CreateNewPermission)
	permissions.Get("/:id", controller.FindPermissionById)
	permissions.Put("/:id", controller.UpdatePermissionById)
	permissions.Delete("/:id", controller.DeletePermissionById)

	// Roles Controller
	roles := a.Group("/roles")
	roles.Get("/", controller.FindAllRoles)
	roles.Post("/", controller.CreateNewRole)
	roles.Get("/:id", controller.FindRoleById)
	roles.Put("/:id", controller.UpdateRoleById)
	roles.Delete("/:id", controller.DeleteRoleById)

	// Movies Group
	movies := a.Group("/movies")

	// Movies Controller
	movieTypes := movies.Group("/types")
	movieTypes.Get("/", controller.FindAllMovieTypes)
	movieTypes.Post("/", controller.CreateNewMovieType)
	movieTypes.Get("/:id", controller.FindMovieTypeById)
	movieTypes.Put("/:id", controller.UpdateMovieTypeById)
	movieTypes.Delete("/:id", controller.DeleteMovieTypeById)

	// Movies Controller
	movies.Get("/", controller.FindAllMovies)
	movies.Post("/", controller.CreateNewMovie)

	// Movie detail
	movieDetails := movies.Group("/:id")
	movieDetails.Get("/", controller.FindMovieById)
	movieDetails.Put("/", controller.UpdateMovieById)
	movieDetails.Delete("/", controller.DeleteMovieById)

	movieDetails.Get("/genres", controller.FindAllMovieGenreById)
	movieDetails.Get("/countries", controller.FindAllMovieCountryById)
	movieDetails.Get("/tags", controller.FindMovieById)
}
