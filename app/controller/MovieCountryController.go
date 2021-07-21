package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

func FindAllMovieCountryById(c *fiber.Ctx) error {
	movieId := util.ParseStringToUInt(c.Params("id"))
	countries, err := repository.FindAllCountriesByMovieIdAndStatusNotIn(movieId, []int{util.StatusDeleted, util.StatusDraft})

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", countries)
}
