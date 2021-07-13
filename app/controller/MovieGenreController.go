package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

func FindAllMovieGenreById(c *fiber.Ctx) error {
	movieId := util.ParseStringToUInt(c.Params("id"))
	genres, err := repository.FindAllGenresByMovieIdAndStatusNotIn(movieId, []int{util.StatusDeleted, util.StatusDraft})

	if err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", genres)
}