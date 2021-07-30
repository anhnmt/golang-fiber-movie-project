package response

import "github.com/xdorro/golang-fiber-base-project/app/entity/model"

type MovieEpisodeDetailResponse struct {
	model.Episode
	EpisodeDetails []model.EpisodeDetail `json:"episode_details"`
}
