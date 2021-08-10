package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

func FindAllMovieGenreById(c *fiber.Ctx) error {
	movieGenreRepository := repository.NewMovieGenreRepository()
	movieId := util.ParseStringToUInt(c.Params("id"))
	genres, err := movieGenreRepository.FindAllGenresByMovieIdAndStatusNotIn(movieId, []int{util.StatusDeleted, util.StatusDraft})

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", genres)
}
