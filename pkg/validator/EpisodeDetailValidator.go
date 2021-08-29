package validator

import (
	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
	"github.com/xdorro/golang-fiber-movie-project/app/repository"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
)

func ValidateEpisodeDetailId(episodeDetailId string) (*model.EpisodeDetail, error) {
	episodeDetailRepository := repository.NewEpisodeDetailRepository()

	episodeDetail, err := episodeDetailRepository.FindEpisodeDetailByIdAndStatusNot(episodeDetailId, []int{util.StatusDeleted})

	if err != nil || episodeDetail.EpisodeDetailId == 0 {
		return nil, util.ResponseBadRequest("EpisodeDetailId không tồn tại", err)
	}

	return episodeDetail, nil
}

func ValidateEpisodeIdAndEpisodeDetailId(episodeId string, episodeDetailId string) (*model.EpisodeDetail, error) {
	episodeDetailRepository := repository.NewEpisodeDetailRepository()

	episodeDetail, err := episodeDetailRepository.FindEpisodeDetailByEpisodeIdAndEpisodeDetailIdAndStatusNot(episodeId, episodeDetailId, []int{util.StatusDeleted})

	if err != nil || episodeDetail.EpisodeDetailId == 0 {
		return nil, util.ResponseBadRequest("EpisodeDetailId không tồn tại", err)
	}

	return episodeDetail, nil
}
