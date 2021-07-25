package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/controller"
)

func privateRoute(a fiber.Router) {

	// Tags Controller
	tagController := controller.NewTagController()
	tags := a.Group("/tags")
	tags.Get("/", tagController.FindAllTags)
	tags.Post("/", tagController.CreateNewTag)
	tags.Get("/:id", tagController.FindTagById)
	tags.Put("/:id", tagController.UpdateTagById)
	tags.Delete("/:id", tagController.DeleteTagById)

	// Genres Controller
	genreController := controller.NewGenreController()
	genres := a.Group("/genres")
	genres.Get("/", genreController.FindAllGenres)
	genres.Post("/", genreController.CreateNewGenre)
	genres.Get("/:id", genreController.FindGenreById)
	genres.Put("/:id", genreController.UpdateGenreById)
	genres.Delete("/:id", genreController.DeleteGenreById)

	// Countries Controller
	countryController := controller.NewCountryController()
	countries := a.Group("/countries")
	countries.Get("/", countryController.FindAllCountries)
	countries.Post("/", countryController.CreateNewCountry)
	countries.Get("/:id", countryController.FindCountryById)
	countries.Put("/:id", countryController.UpdateCountryById)
	countries.Delete("/:id", countryController.DeleteCountryById)

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
	movieTypeController := controller.NewMovieTypeController()
	movieTypes := movies.Group("/types")
	movieTypes.Get("/", movieTypeController.FindAllMovieTypes)
	movieTypes.Post("/", movieTypeController.CreateNewMovieType)
	movieTypes.Get("/:id", movieTypeController.FindMovieTypeById)
	movieTypes.Put("/:id", movieTypeController.UpdateMovieTypeById)
	movieTypes.Delete("/:id", movieTypeController.DeleteMovieTypeById)

	// Movies Controller
	movieController := controller.NewMovieController()
	movies.Get("/", movieController.FindAllMovies)
	movies.Post("/", movieController.CreateNewMovie)

	// Movie Detail Controller
	movieDetails := movies.Group("/:id")
	movieDetails.Get("/", movieController.FindMovieById)
	movieDetails.Put("/", movieController.UpdateMovieById)
	movieDetails.Delete("/", movieController.DeleteMovieById)

	movieDetails.Get("/genres", controller.FindAllMovieGenreById)
	movieDetails.Get("/countries", controller.FindAllMovieCountryById)
	movieDetails.Get("/tags", movieController.FindMovieById)

	// Movie Episode
	episodeController := controller.NewEpisodeController()
	movieDetails.Get("/episodes", episodeController.FindAllEpisodesByMovieId)
	movieDetails.Post("/episodes", episodeController.CreateEpisodesByMovieId)
	movieDetails.Get("/episodes/:episodeId", episodeController.FindEpisodeByMovieIdAndEpisodeId)
	movieDetails.Put("/episodes/:episodeId", episodeController.UpdateEpisodesByMovieIdAndEpisodeId)
	movieDetails.Delete("/episodes/:episodeId", episodeController.DeleteEpisodesByMovieIdAndEpisodeId)

	// Episodes Group
	episodes := a.Group("/episodes")

	// Episode Type
	episodeTypeController := controller.NewEpisodeTypeController()
	episodeTypes := episodes.Group("/types")
	episodeTypes.Get("/", episodeTypeController.FindAllEpisodeTypes)
	episodeTypes.Post("/", episodeTypeController.CreateNewEpisodeType)
	episodeTypes.Get("/:id", episodeTypeController.FindEpisodeTypeById)
	episodeTypes.Put("/:id", episodeTypeController.UpdateEpisodeTypeById)
	episodeTypes.Delete("/:id", episodeTypeController.DeleteEpisodeTypeById)
}
