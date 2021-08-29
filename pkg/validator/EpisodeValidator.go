package validator

import (
	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
	"github.com/xdorro/golang-fiber-movie-project/app/repository"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
)

func ValidateEpisodeId(episodeId string) (*model.Episode, error) {
	episodeRepository := repository.NewEpisodeRepository()

	episode, err := episodeRepository.FindEpisodeByIdAndStatusNot(episodeId, util.StatusDeleted)

	if err != nil || episode.EpisodeId == 0 {
		return nil, util.ResponseBadRequest("EpisodeId không tồn tại", err)
	}

	return episode, nil
}
