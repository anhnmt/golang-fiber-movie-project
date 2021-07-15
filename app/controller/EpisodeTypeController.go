package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

func FindAllEpisodeTypes(c *fiber.Ctx) error {
	episodeTypes, err := repository.FindAllEpisodeTypesByStatusNot(util.StatusDeleted)

	if err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", episodeTypes)
}

func FindEpisodeTypeById(c *fiber.Ctx) error {
	episodeTypeId := c.Params("id")
	episodeType, err := repository.FindEpisodeTypeByIdAndStatusNot(episodeTypeId, util.StatusDeleted)

	if err != nil || episodeType.EpisodeTypeId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	return util.ResponseSuccess(c, "Thành công", episodeType)
}

// CreateNewEpisodeType : Create a new episodeType
func CreateNewEpisodeType(c *fiber.Ctx) error {
	episodeTypeRequest := new(request.EpisodeTypeRequest)

	if err := c.BodyParser(episodeTypeRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	episodeType := model.EpisodeType{
		Name:   episodeTypeRequest.Name,
		Status: episodeTypeRequest.Status,
	}

	if _, err := repository.SaveEpisodeType(episodeType); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// UpdateEpisodeTypeById : Update episodeType by Episode_Type_Id and Status = 1
func UpdateEpisodeTypeById(c *fiber.Ctx) error {
	episodeTypeId := c.Params("id")
	episodeType, err := repository.FindEpisodeTypeByIdAndStatusNot(episodeTypeId, util.StatusDeleted)

	if err != nil || episodeType.EpisodeTypeId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	episodeTypeRequest := new(request.EpisodeTypeRequest)
	if err = c.BodyParser(episodeTypeRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	episodeType.Name = episodeTypeRequest.Name
	episodeType.Status = episodeTypeRequest.Status

	if _, err = repository.SaveEpisodeType(*episodeType); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// DeleteEpisodeTypeById : Delete episodeType by EpisodeType_Id and Status = 1
func DeleteEpisodeTypeById(c *fiber.Ctx) error {
	episodeTypeId := c.Params("id")
	episodeType, err := repository.FindEpisodeTypeByIdAndStatusNot(episodeTypeId, util.StatusDeleted)

	if err != nil || episodeType.EpisodeTypeId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	episodeType.Status = util.StatusDeleted

	// Update movieType status
	if _, err = repository.SaveEpisodeType(*episodeType); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	// Update movie status
	if err = repository.UpdateStatusByEpisodeTypeId(episodeType.EpisodeTypeId, episodeType.Status); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}
