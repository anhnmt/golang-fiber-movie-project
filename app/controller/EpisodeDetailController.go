package controller

import (
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"

	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
	"github.com/xdorro/golang-fiber-movie-project/app/entity/request"
	"github.com/xdorro/golang-fiber-movie-project/app/repository"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
	"github.com/xdorro/golang-fiber-movie-project/pkg/validator"
)

type EpisodeDetailController struct {
	episodeDetailRepository *repository.EpisodeDetailRepository
}

func NewEpisodeDetailController() *EpisodeDetailController {
	if episodeDetailController == nil {
		once = &sync.Once{}

		once.Do(func() {
			if episodeDetailController == nil {
				episodeDetailController = &EpisodeDetailController{
					episodeDetailRepository: repository.NewEpisodeDetailRepository(),
				}
			}

			log.Println("Create new EpisodeDetailController")
		})
	}

	return episodeDetailController
}

func (obj *EpisodeDetailController) ClientFindEpisodeDetailByEpisodeId(c *fiber.Ctx) error {
	episodeId := c.Params("episodeId")

	_, err := validator.ValidateEpisodeId(episodeId)

	if err != nil {
		return err
	}

	episodeDetails, err := obj.episodeDetailRepository.FindEpisodeDetailsByIdAndStatusNotIn(episodeId, []int{
		util.StatusDraft, util.StatusDeleted,
	})

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", episodeDetails)
}

func (obj *EpisodeDetailController) FindEpisodeDetailById(c *fiber.Ctx) error {
	episodeId := c.Params("episodeId")
	episodeDetailId := c.Params("episodeDetailId")

	_, err := validator.ValidateEpisodeId(episodeId)

	if err != nil {
		return err
	}

	episodeDetail, err := validator.ValidateEpisodeIdAndEpisodeDetailId(episodeId, episodeDetailId)

	if err != nil {
		return err
	}

	return util.ResponseSuccess("Thành công", episodeDetail)
}

func (obj *EpisodeDetailController) CreateEpisodeDetailById(c *fiber.Ctx) error {
	episodeId := c.Params("episodeId")

	episode, err := validator.ValidateEpisodeId(episodeId)

	if err != nil {
		return err
	}

	episodeDetailRequest := new(request.EpisodeDetailRequest)
	if err = c.BodyParser(episodeDetailRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	newEpisodeDetail := model.EpisodeDetail{
		EpisodeId:     episode.EpisodeId,
		Name:          episodeDetailRequest.Name,
		Link:          episodeDetailRequest.Link,
		EpisodeTypeId: episodeDetailRequest.EpisodeTypeId,
		Status:        episodeDetailRequest.Status,
	}

	if err = obj.episodeDetailRepository.CreateEpisodeDetailsByEpisodeId(newEpisodeDetail); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

func (obj *EpisodeDetailController) UpdateEpisodeDetailById(c *fiber.Ctx) error {
	episodeId := c.Params("episodeId")
	episodeDetailId := c.Params("episodeDetailId")

	_, err := validator.ValidateEpisodeId(episodeId)

	if err != nil {
		return err
	}

	episodeDetail, err := validator.ValidateEpisodeIdAndEpisodeDetailId(episodeId, episodeDetailId)

	if err != nil {
		return err
	}

	episodeDetailRequest := new(request.EpisodeDetailRequest)
	if err = c.BodyParser(episodeDetailRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	episodeDetail.Name = episodeDetailRequest.Name
	episodeDetail.Link = episodeDetailRequest.Link
	episodeDetail.EpisodeTypeId = episodeDetailRequest.EpisodeTypeId
	episodeDetail.Status = episodeDetailRequest.Status

	if _, err = obj.episodeDetailRepository.UpdateEpisodeDetail(episodeDetailId, *episodeDetail); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

func (obj *EpisodeDetailController) DeleteEpisodeDetailById(c *fiber.Ctx) error {
	episodeId := c.Params("episodeId")
	episodeDetailId := c.Params("episodeDetailId")

	_, err := validator.ValidateEpisodeId(episodeId)

	if err != nil {
		return err
	}

	_, err = validator.ValidateEpisodeIdAndEpisodeDetailId(episodeId, episodeDetailId)

	if err != nil {
		return err
	}

	if err = obj.episodeDetailRepository.UpdateStatusByEpisodeDetailId(episodeDetailId, util.StatusDeleted); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}
