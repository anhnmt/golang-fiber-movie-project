package mapper

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/response"
)

func EpisodeDetailsMapper(episodes *[]model.Episode, episodeDetails *[]model.EpisodeDetail) []response.MovieEpisodeDetailResponse {
	result := make([]response.MovieEpisodeDetailResponse, 0)

	for _, episode := range *episodes {
		detail := make([]model.EpisodeDetail, 0)

		for _, episodeDetail := range *episodeDetails {
			if episode.EpisodeId == episodeDetail.EpisodeId {
				detail = append(detail, episodeDetail)
			}
		}

		mapper := EpisodeDetails(&episode, &detail)
		result = append(result, *mapper)
	}

	return result
}

func EpisodeDetails(episode *model.Episode, episodeDetails *[]model.EpisodeDetail) *response.MovieEpisodeDetailResponse {
	return &response.MovieEpisodeDetailResponse{
		Episode:        *episode,
		EpisodeDetails: *episodeDetails,
	}
}
