package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/xdorro/golang-fiber-movie-project/app/repository"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
)

func FindAllMovieCountryById(c *fiber.Ctx) error {
	movieCountryRepository := repository.NewMovieCountryRepository()
	movieId := util.ParseStringToInt64(c.Params("movieId"))
	countries, err := movieCountryRepository.FindAllCountriesByMovieIdAndStatusNotIn(movieId, []int{
		util.StatusDeleted, util.StatusDraft,
	})

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", countries)
}
