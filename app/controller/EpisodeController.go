package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/mapper"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

func FindAllEpisodesByMovieId(c *fiber.Ctx) error {
	movieId := c.Params("id")

	movie, err := repository.FindMovieByIdAndStatusNotJoinMovieType(movieId, util.StatusDeleted)

	if err != nil || movie.MovieId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	episodes, err := repository.FindAllEpisodesByMovieIdAndStatusNot(movieId, util.StatusDeleted)

	if err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", episodes)
}

func CreateEpisodesByMovieId(c *fiber.Ctx) error {
	movieId := c.Params("id")

	movie, err := repository.FindMovieByIdAndStatusNotJoinMovieType(movieId, util.StatusDeleted)

	if err != nil || movie.MovieId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	episodeRequest := new(request.EpisodeRequest)
	if err = c.BodyParser(episodeRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	newEpisode := model.Episode{
		MovieId: movie.MovieId,
		Name:    episodeRequest.Name,
		Status:  episodeRequest.Status,
	}

	episode, err := repository.SaveEpisode(newEpisode)
	if err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	if episode.EpisodeId == 0 {
		return util.ResponseBadRequest(c, "Thêm mới thất bại", nil)
	}

	err = createEpisodeDetails(&episode.EpisodeId, &episodeRequest.EpisodeDetail)
	if err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

func createEpisodeDetails(episodeId *uint, newEpisodeDetails *[]request.EpisodeDetailRequest) error {
	if len(*newEpisodeDetails) > 0 {
		// Create Movie Genre
		episodeDetails := mapper.EpisodeDetails(episodeId, newEpisodeDetails)

		if err := repository.CreateEpisodeDetailsByEpisodeId(episodeDetails); err != nil {
			return err
		}
	}

	return nil
}
