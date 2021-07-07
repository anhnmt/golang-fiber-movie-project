package controller

import (
	"github.com/gofiber/fiber/v2"
	model2 "github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

// FindAllMovieTypes : Find all moveTypes by Status = 1
func FindAllMovieTypes(c *fiber.Ctx) error {
	moveTypes, err := repository.FindAllMovieTypesByStatusNot(util.STATUS_DELETED)

	if err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", moveTypes)
}

// FindMovieTypeById : Find moveType by MovieType_Id and Status = 1
func FindMovieTypeById(c *fiber.Ctx) error {
	moveTypeId := c.Params("id")
	moveType, err := repository.FindMovieTypeByIdAndStatusNot(moveTypeId, util.STATUS_DELETED)

	if err != nil || moveType.MovieTypeId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	return util.ResponseSuccess(c, "Thành công", moveType)
}

// CreateNewMovieType : Create a new moveType
func CreateNewMovieType(c *fiber.Ctx) error {
	moveTypeRequest := new(request.MovieTypeRequest)

	if err := c.BodyParser(moveTypeRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	moveType := model2.MovieType{
		Name:   moveTypeRequest.Name,
		Slug:   moveTypeRequest.Slug,
		Status: moveTypeRequest.Status,
	}

	if _, err := repository.SaveMovieType(moveType); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// UpdateMovieTypeById : Update moveType by MovieType_Id and Status = 1
func UpdateMovieTypeById(c *fiber.Ctx) error {
	moveTypeId := c.Params("id")
	moveType, err := repository.FindMovieTypeByIdAndStatusNot(moveTypeId, util.STATUS_DELETED)

	if err != nil || moveType.MovieTypeId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	moveTypeRequest := new(request.MovieTypeRequest)
	if err = c.BodyParser(moveTypeRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	moveType.Name = moveTypeRequest.Name
	moveType.Slug = moveTypeRequest.Slug
	moveType.Status = moveTypeRequest.Status

	if _, err = repository.SaveMovieType(*moveType); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// DeleteMovieTypeById : Delete moveType by MovieType_Id and Status = 1
func DeleteMovieTypeById(c *fiber.Ctx) error {
	moveTypeId := c.Params("id")
	moveType, err := repository.FindMovieTypeByIdAndStatusNot(moveTypeId, util.STATUS_DELETED)

	if err != nil || moveType.MovieTypeId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	moveType.Status = util.STATUS_DELETED

	if _, err = repository.SaveMovieType(*moveType); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}
