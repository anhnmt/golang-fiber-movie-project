package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/entity/response"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/mapper"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"github.com/xdorro/golang-fiber-base-project/pkg/validator"
	"log"
	"sync"
)

type EpisodeController struct {
	episodeRepository       *repository.EpisodeRepository
	episodeDetailRepository *repository.EpisodeDetailRepository
}

func NewEpisodeController() *EpisodeController {
	if episodeController == nil {
		once = &sync.Once{}

		once.Do(func() {
			episodeController = &EpisodeController{
				episodeRepository:       repository.NewEpisodeRepository(),
				episodeDetailRepository: repository.NewEpisodeDetailRepository(),
			}
			log.Println("Create new EpisodeController")
		})
	}

	return episodeController
}

func (obj *EpisodeController) ClientFindEpisodesByMovieId(c *fiber.Ctx) error {
	movieId := c.Params("movieId")

	if _, err := validator.ValidateMovieId(movieId); err != nil {
		return err
	}

	episodes, err := obj.episodeRepository.FindAllEpisodesByMovieIdAndStatusNotIn(movieId, []int{util.StatusDraft, util.StatusDeleted})

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	episodeIds := mapper.GetEpisodeIds(*episodes)

	var episodeDetails *[]model.EpisodeDetail

	if len(*episodeIds) > 0 {
		episodeDetails, err = obj.episodeDetailRepository.
			FindEpisodeDetailsByEpisodeIdInAndStatusNotIn(*episodeIds, []int{util.StatusDraft, util.StatusDeleted})

		if err != nil {
			return util.ResponseError(err.Error(), nil)
		}
	}

	result := mapper.EpisodeDetailsMapper(episodes, episodeDetails)

	return util.ResponseSuccess("Thành công", result)
}

func (obj *EpisodeController) FindAllEpisodesByMovieId(c *fiber.Ctx) error {
	movieId := c.Params("movieId")

	if _, err := validator.ValidateMovieId(movieId); err != nil {
		return err
	}

	episodes, err := obj.episodeRepository.FindAllEpisodesByMovieIdAndStatusNot(movieId, util.StatusDeleted)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", episodes)
}

func (obj *EpisodeController) FindEpisodeByEpisodeId(c *fiber.Ctx) error {
	episodeId := c.Params("episodeId")

	episode, err := validator.ValidateEpisodeId(episodeId)

	if err != nil {
		return err
	}

	episodeDetails, err := obj.episodeDetailRepository.FindEpisodeDetailsByIdAndStatusNotIn(episodeId, []int{util.StatusDeleted})

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := &response.MovieEpisodeDetailResponse{
		Episode:        *episode,
		EpisodeDetails: *episodeDetails,
	}

	return util.ResponseSuccess("Thành công", result)
}

func (obj *EpisodeController) CreateEpisodesByMovieId(c *fiber.Ctx) error {
	movieId := c.Params("movieId")

	movie, err := validator.ValidateMovieId(movieId)
	if err != nil {
		return err
	}

	episodeRequest := new(request.EpisodeRequest)
	if err = c.BodyParser(episodeRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	newEpisode := model.Episode{
		MovieId: movie.MovieId,
		Name:    episodeRequest.Name,
		Status:  episodeRequest.Status,
	}

	episode, err := obj.episodeRepository.SaveEpisode(newEpisode)
	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	if episode.EpisodeId == 0 {
		return util.ResponseBadRequest("Thêm mới thất bại", nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

func (obj *EpisodeController) UpdateEpisodesByEpisodeId(c *fiber.Ctx) error {
	episodeId := c.Params("episodeId")

	episode, err := validator.ValidateEpisodeId(episodeId)
	if err != nil {
		return err
	}

	episodeRequest := new(request.EpisodeRequest)
	if err = c.BodyParser(episodeRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	episode.Name = episodeRequest.Name
	episode.Status = episodeRequest.Status

	episode, err = obj.episodeRepository.UpdateEpisode(episodeId, *episode)
	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	if episode.EpisodeId == 0 {
		return util.ResponseBadRequest("Thêm mới thất bại", nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

func (obj *EpisodeController) DeleteEpisodesByEpisodeId(c *fiber.Ctx) error {
	episodeId := c.Params("episodeId")

	_, err := validator.ValidateEpisodeId(episodeId)
	if err != nil {
		return err
	}

	// Update episode status
	if err = obj.episodeRepository.UpdateStatusByEpisodeId(episodeId, util.StatusDeleted); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	// Delete episodeDetails
	if err = obj.episodeDetailRepository.UpdateStatusByEpisodeId(episodeId, util.StatusDeleted); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}
