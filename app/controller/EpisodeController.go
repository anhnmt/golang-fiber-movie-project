package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

func FindAllEpisodesByMovieId(c *fiber.Ctx) error {
	movieId := c.Params("id")
	episodes, err := repository.FindAllEpisodesByMovieIdAndStatusNot(movieId, util.StatusDeleted)

	if err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", episodes)
}
