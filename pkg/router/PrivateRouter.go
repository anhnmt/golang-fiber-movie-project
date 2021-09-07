package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-movie-project/app/controller"
	"github.com/xdorro/golang-fiber-movie-project/pkg/middleware"
)

func privateRoute(a fiber.Router) {
	// Dashboard Controller
	dashboardController := controller.NewDashboardController()
	{
		dashboard := a.Group("/dashboard", middleware.Protected())
		dashboard.Get("/", dashboardController.DashboardAnalyzer)
	}

	// Banners Controller
	bannerController := controller.NewBannerController()
	{
		banners := a.Group("/banners", middleware.Protected())
		banners.Get("/", bannerController.FindAllBanners)
		banners.Post("/", bannerController.CreateNewBanner)
		banners.Get("/:bannerId", bannerController.FindBannerById)
		banners.Put("/:bannerId", bannerController.UpdateBannerById)
		banners.Delete("/:bannerId", bannerController.DeleteBannerById)
	}

	// Tags Controller
	tagController := controller.NewTagController()
	{
		tags := a.Group("/tags", middleware.Protected())
		tags.Get("/", tagController.FindAllTags)
		tags.Post("/", tagController.CreateNewTag)
		tags.Get("/:tagId", tagController.FindTagById)
		tags.Put("/:tagId", tagController.UpdateTagById)
		tags.Delete("/:tagId", tagController.DeleteTagById)
	}

	// Genres Controller
	genreController := controller.NewGenreController()
	{
		genres := a.Group("/genres", middleware.Protected())
		genres.Get("/check_slug", genreController.CheckIsExistGenreSlug)
		genres.Get("/", genreController.FindAllGenres)
		genres.Post("/", genreController.CreateNewGenre)
		genres.Get("/:genreId", genreController.FindGenreById)
		genres.Put("/:genreId", genreController.UpdateGenreById)
		genres.Delete("/:genreId", genreController.DeleteGenreById)
	}

	// Countries Controller
	countryController := controller.NewCountryController()
	{
		countries := a.Group("/countries", middleware.Protected())
		countries.Get("/check_slug", countryController.CheckIsExistCountrySlug)
		countries.Get("/", countryController.FindAllCountries)
		countries.Post("/", countryController.CreateNewCountry)
		countries.Get("/:countryId", countryController.FindCountryById)
		countries.Put("/:countryId", countryController.UpdateCountryById)
		countries.Delete("/:countryId", countryController.DeleteCountryById)
	}

	// Users Controller
	userController := controller.NewUserController()
	{
		users := a.Group("/users", middleware.Protected())
		users.Get("/", userController.FindAllUsers)
		users.Post("/", userController.CreateNewUser)
		users.Get("/check_username", userController.CheckIsExistUsername)
		users.Get("/:userId", userController.FindUserById)
		users.Put("/:userId", userController.UpdateUserById)
		users.Delete("/:userId", userController.DeleteUserById)
	}

	// UserRoles Controller
	//users.Get("/:id/roles", controller.FindAllUserRoles)

	// Roles Controller
	{
		roles := a.Group("/roles", middleware.Protected())
		roles.Get("/", controller.FindAllRoles)
		roles.Post("/", controller.CreateNewRole)
		roles.Get("/:roleId", controller.FindRoleById)
		roles.Put("/:roleId", controller.UpdateRoleById)
		roles.Delete("/:roleId", controller.DeleteRoleById)
	}

	// Movies Group
	movies := a.Group("/movies", middleware.Protected())

	// Movies Controller
	movieTypeController := controller.NewMovieTypeController()
	movieTypes := movies.Group("/types")
	movieTypes.Get("/", movieTypeController.FindAllMovieTypes)
	movieTypes.Post("/", movieTypeController.CreateNewMovieType)
	movieTypes.Get("/check_slug", movieTypeController.CheckIsExistMovieTypeSlug)
	movieTypes.Get("/:movieTypeId", movieTypeController.FindMovieTypeById)
	movieTypes.Put("/:movieTypeId", movieTypeController.UpdateMovieTypeById)
	movieTypes.Delete("/:movieTypeId", movieTypeController.DeleteMovieTypeById)

	// Movies Controller
	movieController := controller.NewMovieController()
	movieTypes.Get("/:movieSlug", movieController.FindAllMovies)
	movies.Get("/", movieController.FindAllMovies)
	movies.Post("/", movieController.CreateNewMovie)
	movies.Get("/check_slug", movieController.CheckIsExistMovieSlug)

	// Movie Detail Controller
	movieDetails := movies.Group("/:movieId")
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

	// Episode Type
	episodeTypeController := controller.NewEpisodeTypeController()
	episodeTypes := a.Group("/episode-types", middleware.Protected())
	{
		episodeTypes.Get("/", episodeTypeController.FindAllEpisodeTypes)
		episodeTypes.Post("/", episodeTypeController.CreateNewEpisodeType)
		episodeTypes.Get("/:episodeTypeId", episodeTypeController.FindEpisodeTypeById)
		episodeTypes.Put("/:episodeTypeId", episodeTypeController.UpdateEpisodeTypeById)
		episodeTypes.Delete("/:episodeTypeId", episodeTypeController.DeleteEpisodeTypeById)
	}

	// Episodes Group
	episodes := a.Group("/episodes", middleware.Protected())

	episodeDetail := episodes.Group("/:episodeId")

	episodeDetail.Get("/", episodeController.FindEpisodeByEpisodeId)
	episodeDetail.Put("/", episodeController.UpdateEpisodesByEpisodeId)
	episodeDetail.Delete("/", episodeController.DeleteEpisodesByEpisodeId)

	// Episode Detail
	episodeDetailController := controller.NewEpisodeDetailController()
	episodeDetail.Post("/", episodeDetailController.CreateEpisodeDetailById)

	episodeDetails := episodeDetail.Group("/details")
	episodeDetails.Get("/:episodeDetailId", episodeDetailController.FindEpisodeDetailById)
	episodeDetails.Put("/:episodeDetailId", episodeDetailController.UpdateEpisodeDetailById)
	episodeDetails.Delete("/:episodeDetailId", episodeDetailController.DeleteEpisodeDetailById)

}
