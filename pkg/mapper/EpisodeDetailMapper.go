package mapper

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
)

func EpisodeDetails(episodeId *uint, episodeDetails *[]request.EpisodeDetailRequest) []model.EpisodeDetail {
	result := make([]model.EpisodeDetail, 0)

	for _, episodeDetail := range *episodeDetails {
		mapper := EpisodeDetail(episodeId, &episodeDetail)
		result = append(result, *mapper)
	}

	return result
}

func EpisodeDetail(episodeId *uint, episodeDetail *request.EpisodeDetailRequest) *model.EpisodeDetail {
	return &model.EpisodeDetail{
		EpisodeDetailId: episodeDetail.EpisodeDetailId,
		Name:            episodeDetail.Name,
		EpisodeId:       *episodeId,
		EpisodeTypeId:   episodeDetail.EpisodeTypeId,
		Status:          episodeDetail.Status,
	}
}
