package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"log"
	"sync"
)

type EpisodeTypeController struct {
	episodeTypeRepository *repository.EpisodeTypeRepository
}

func NewEpisodeTypeController() *EpisodeTypeController {
	if episodeTypeController == nil {
		once = &sync.Once{}

		once.Do(func() {
			episodeTypeController = &EpisodeTypeController{
				episodeTypeRepository: repository.NewEpisodeTypeRepository(),
			}
			log.Println("Create new EpisodeTypeController")
		})
	}

	return episodeTypeController
}

func (obj *EpisodeTypeController) FindAllEpisodeTypes(c *fiber.Ctx) error {
	episodeTypes, err := obj.episodeTypeRepository.FindAllEpisodeTypesByStatusNot(util.StatusDeleted)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", episodeTypes)
}

func (obj *EpisodeTypeController) FindEpisodeTypeById(c *fiber.Ctx) error {
	episodeTypeId := c.Params("id")
	episodeType, err := obj.episodeTypeRepository.FindEpisodeTypeByIdAndStatusNot(episodeTypeId, util.StatusDeleted)

	if err != nil || episodeType.EpisodeTypeId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	return util.ResponseSuccess("Thành công", episodeType)
}

// CreateNewEpisodeType : Create a new episodeType
func (obj *EpisodeTypeController) CreateNewEpisodeType(c *fiber.Ctx) error {
	episodeTypeRequest := new(request.EpisodeTypeRequest)

	if err := c.BodyParser(episodeTypeRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	episodeType := model.EpisodeType{
		Name:   episodeTypeRequest.Name,
		Status: episodeTypeRequest.Status,
	}

	if _, err := obj.episodeTypeRepository.SaveEpisodeType(episodeType); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// UpdateEpisodeTypeById : Update episodeType by Episode_Type_Id and Status = 1
func (obj *EpisodeTypeController) UpdateEpisodeTypeById(c *fiber.Ctx) error {
	episodeTypeId := c.Params("id")
	episodeType, err := obj.episodeTypeRepository.FindEpisodeTypeByIdAndStatusNot(episodeTypeId, util.StatusDeleted)

	if err != nil || episodeType.EpisodeTypeId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	episodeTypeRequest := new(request.EpisodeTypeRequest)
	if err = c.BodyParser(episodeTypeRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	episodeType.Name = episodeTypeRequest.Name
	episodeType.Status = episodeTypeRequest.Status

	if _, err = obj.episodeTypeRepository.SaveEpisodeType(*episodeType); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// DeleteEpisodeTypeById : Delete episodeType by EpisodeType_Id and Status = 1
func (obj *EpisodeTypeController) DeleteEpisodeTypeById(c *fiber.Ctx) error {
	episodeTypeId := c.Params("id")
	episodeType, err := obj.episodeTypeRepository.FindEpisodeTypeByIdAndStatusNot(episodeTypeId, util.StatusDeleted)

	if err != nil || episodeType.EpisodeTypeId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	episodeType.Status = util.StatusDeleted

	// Update movieType status
	if _, err = obj.episodeTypeRepository.SaveEpisodeType(*episodeType); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	// Update movie status
	if err = obj.episodeTypeRepository.UpdateStatusByEpisodeTypeId(episodeType.EpisodeTypeId, episodeType.Status); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}
