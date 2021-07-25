package validator

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

func ValidateEpisodeId(movieId string) (*model.Episode, error) {
	episodeRepository := repository.NewEpisodeRepository()

	episode, err := episodeRepository.FindEpisodeByIdAndStatusNot(movieId, util.StatusDeleted)

	if err != nil || episode.EpisodeId == 0 {
		return nil, util.ResponseBadRequest("EpisodeId không tồn tại", err)
	}

	return episode, nil
}
