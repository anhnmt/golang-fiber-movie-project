package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-movie-project/app/entity/response"
	"github.com/xdorro/golang-fiber-movie-project/app/repository"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
	"log"
	"sync"
)

type DashboardController struct {
	bannerRepository  *repository.BannerRepository
	genreRepository   *repository.GenreRepository
	countryRepository *repository.CountryRepository
	movieRepository   *repository.MovieRepository
	userRepository    *repository.UserRepository
}

func NewDashboardController() *DashboardController {
	if dashboardController == nil {
		once = &sync.Once{}

		once.Do(func() {
			if dashboardController == nil {
				dashboardController = &DashboardController{
					bannerRepository:  repository.NewBannerRepository(),
					genreRepository:   repository.NewGenreRepository(),
					countryRepository: repository.NewCountryRepository(),
					movieRepository:   repository.NewMovieRepository(),
				}
				log.Println("Create new DashboardController")
			}
		})
	}

	return dashboardController
}

func (obj *DashboardController) DashboardAnalyzer(*fiber.Ctx) error {
	status := []int{util.StatusDeleted}

	movies, _ := obj.movieRepository.CountAllMoviesStatusNotIn(status)
	banners, _ := obj.bannerRepository.CountAllBannersStatusNotIn(status)
	countries, _ := obj.countryRepository.CountAllCountriesStatusNotIn(status)
	genres, _ := obj.genreRepository.CountAllGenresStatusNotIn(status)

	latestMovies, _ := obj.movieRepository.FindAllTopMoviesByStatusNotInAndLimit(status, 5)

	result := response.DashboardAnalyzerResponse{
		Movies:       movies,
		Banners:      banners,
		Countries:    countries,
		Genres:       genres,
		LatestMovies: latestMovies,
	}

	return util.ResponseSuccess("Thành công", result)
}
