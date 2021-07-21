package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

// FindAllMovieTypes : Find all moveTypes by Status = 1
func FindAllMovieTypes(c *fiber.Ctx) error {
	moveTypes, err := repository.FindAllMovieTypesByStatusNot(util.StatusDeleted)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", moveTypes)
}

// FindMovieTypeById : Find moveType by MovieType_Id and Status = 1
func FindMovieTypeById(c *fiber.Ctx) error {
	moveTypeId := c.Params("id")
	moveType, err := repository.FindMovieTypeByIdAndStatusNot(moveTypeId, util.StatusDeleted)

	if err != nil || moveType.MovieTypeId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	return util.ResponseSuccess("Thành công", moveType)
}

// CreateNewMovieType : Create a new moveType
func CreateNewMovieType(c *fiber.Ctx) error {
	moveTypeRequest := new(request.MovieTypeRequest)

	if err := c.BodyParser(moveTypeRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	moveType := model.MovieType{
		Name:   moveTypeRequest.Name,
		Slug:   moveTypeRequest.Slug,
		Status: moveTypeRequest.Status,
	}

	if _, err := repository.SaveMovieType(moveType); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// UpdateMovieTypeById : Update moveType by MovieType_Id and Status = 1
func UpdateMovieTypeById(c *fiber.Ctx) error {
	moveTypeId := c.Params("id")
	moveType, err := repository.FindMovieTypeByIdAndStatusNot(moveTypeId, util.StatusDeleted)

	if err != nil || moveType.MovieTypeId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	moveTypeRequest := new(request.MovieTypeRequest)
	if err = c.BodyParser(moveTypeRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	moveType.Name = moveTypeRequest.Name
	moveType.Slug = moveTypeRequest.Slug
	moveType.Status = moveTypeRequest.Status

	if _, err = repository.SaveMovieType(*moveType); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// DeleteMovieTypeById : Delete moveType by MovieType_Id and Status = 1
func DeleteMovieTypeById(c *fiber.Ctx) error {
	moveTypeId := c.Params("id")
	moveType, err := repository.FindMovieTypeByIdAndStatusNot(moveTypeId, util.StatusDeleted)

	if err != nil || moveType.MovieTypeId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	moveType.Status = util.StatusDeleted

	// Update movieType status
	if _, err = repository.SaveMovieType(*moveType); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	// Update movie status
	if err = repository.UpdateStatusByMovieTypeId(moveType.MovieTypeId, moveType.Status); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}
