package mapper

import (
	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
)

func GetEpisodeIds(episodes []model.Episode) *[]int64 {
	result := make([]int64, 0)

	for _, episode := range episodes {
		result = append(result, episode.EpisodeId)
	}

	return &result
}
