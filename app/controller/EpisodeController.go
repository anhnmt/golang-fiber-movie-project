package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/dto"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/mapper"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

func FindAllEpisodesByMovieId(c *fiber.Ctx) error {
	movieId := c.Params("id")

	if _, err := validateMovieId(movieId); err != nil {
		return err
	}

	episodes, err := repository.FindAllEpisodesByMovieIdAndStatusNot(movieId, util.StatusDeleted)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", episodes)
}

func CreateEpisodesByMovieId(c *fiber.Ctx) error {
	movieId := c.Params("id")

	movie, err := validateMovieId(movieId)
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

	episode, err := repository.SaveEpisode(newEpisode)
	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	if episode.EpisodeId == 0 {
		return util.ResponseBadRequest("Thêm mới thất bại", nil)
	}

	if err = createEpisodeDetails(&episode.EpisodeId, &episodeRequest.EpisodeDetail); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

func UpdateEpisodesByMovieIdAndEpisodeId(c *fiber.Ctx) error {
	movieId := c.Params("id")
	//episodeId := c.Params("episodeId")

	movie, err := validateMovieId(movieId)
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

	episode, err := repository.SaveEpisode(newEpisode)
	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	if episode.EpisodeId == 0 {
		return util.ResponseBadRequest("Thêm mới thất bại", nil)
	}

	if err = createEpisodeDetails(&episode.EpisodeId, &episodeRequest.EpisodeDetail); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

func createEpisodeDetails(episodeId *uint, newEpisodeDetails *[]request.EpisodeDetailRequest) error {
	if len(*newEpisodeDetails) > 0 {
		// Create Movie Genre
		episodeDetails := mapper.EpisodeDetails(episodeId, newEpisodeDetails)

		return repository.CreateEpisodeDetailsByEpisodeId(episodeDetails)
	}

	return nil
}

func validateMovieId(movieId string) (*model.Movie, error) {
	movie, err := repository.FindMovieByIdAndStatusNot(movieId, util.StatusDeleted)

	if err != nil || movie.MovieId == 0 {
		return nil, util.ResponseBadRequest("ID không tồn tại", err)
	}

	return movie, nil
}

func validateEpisodeId(movieId string) (*dto.MovieDetailDTO, error) {
	movie, err := repository.FindMovieByIdAndStatusNotJoinMovieType(movieId, util.StatusDeleted)

	if err != nil || movie.MovieId == 0 {
		return nil, util.ResponseBadRequest("ID không tồn tại", err)
	}

	return movie, nil
}
