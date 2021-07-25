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
)

type EpisodeController struct {
	episodeRepository       *repository.EpisodeRepository
	episodeDetailRepository *repository.EpisodeDetailRepository
}

func NewEpisodeController() *EpisodeController {
	return episodeController
}

func (obj *EpisodeController) FindAllEpisodesByMovieId(c *fiber.Ctx) error {
	movieId := c.Params("id")

	if _, err := validator.ValidateMovieId(movieId); err != nil {
		return err
	}

	episodes, err := obj.episodeRepository.FindAllEpisodesByMovieIdAndStatusNot(movieId, util.StatusDeleted)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", episodes)
}

func (obj *EpisodeController) FindEpisodeByMovieIdAndEpisodeId(c *fiber.Ctx) error {
	movieId := c.Params("id")
	episodeId := c.Params("episodeId")

	if _, err := validator.ValidateMovieId(movieId); err != nil {
		return err
	}

	episode, err := validator.ValidateEpisodeId(episodeId)

	if err != nil {
		return err
	}

	episodeDetails, err := obj.episodeDetailRepository.FindEpisodeDetailsByIdAndStatusNot(episodeId, []int{util.StatusDeleted})

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := &response.MovieEpisodeDetailResponse{
		Episode: model.Episode{
			EpisodeId: episode.EpisodeId,
			Name:      episode.Name,
			MovieId:   episode.MovieId,
			Status:    episode.Status,
		},
		EpisodeDetails: *episodeDetails,
	}

	return util.ResponseSuccess("Thành công", result)
}

func (obj *EpisodeController) CreateEpisodesByMovieId(c *fiber.Ctx) error {
	movieId := c.Params("id")

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

	if err = obj.createEpisodeDetails(&episode.EpisodeId, &episodeRequest.EpisodeDetail); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

func (obj *EpisodeController) UpdateEpisodesByMovieIdAndEpisodeId(c *fiber.Ctx) error {
	movieId := c.Params("id")
	//episodeId := c.Params("episodeId")

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

	if err = obj.createEpisodeDetails(&episode.EpisodeId, &episodeRequest.EpisodeDetail); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

func (obj *EpisodeController) createEpisodeDetails(episodeId *uint, newEpisodeDetails *[]request.EpisodeDetailRequest) error {
	if len(*newEpisodeDetails) > 0 {
		// Create Movie Genre
		episodeDetails := mapper.EpisodeDetails(episodeId, newEpisodeDetails)

		return obj.episodeDetailRepository.CreateEpisodeDetailsByEpisodeId(episodeDetails)
	}

	return nil
}

func (obj *EpisodeController) DeleteEpisodesByMovieIdAndEpisodeId(c *fiber.Ctx) error {
	movieId := c.Params("id")
	episodeId := c.Params("episodeId")

	_, err := validator.ValidateMovieId(movieId)
	if err != nil {
		return err
	}

	episode, err := validator.ValidateEpisodeId(episodeId)
	if err != nil {
		return err
	}

	episode.Status = util.StatusDeleted

	// Update episode status
	if _, err = obj.episodeRepository.SaveEpisode(*episode); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	// Delete episodeDetails
	if err = obj.episodeDetailRepository.UpdateStatusByEpisodeId(episodeId, util.StatusDeleted); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}
