package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-movie-project/app/repository"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
)

func FindAllMovieGenreById(c *fiber.Ctx) error {
	movieGenreRepository := repository.NewMovieGenreRepository()
	movieId := util.ParseStringToInt64(c.Params("movieId"))
	genres, err := movieGenreRepository.FindAllGenresByMovieIdAndStatusNotIn(movieId, []int{util.StatusDeleted, util.StatusDraft})

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", genres)
}
