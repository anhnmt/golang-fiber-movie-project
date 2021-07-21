package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

func FindAllEpisodesByMovieIdAndStatusNot(id string, status int) (*[]model.Episode, error) {
	movieId := util.ParseStringToUInt(id)
	episodes := make([]model.Episode, 0)

	err := db.Model(model.Episode{}).Find(&episodes, "movie_id = ? AND status <> ?", movieId, status).Error

	return &episodes, err
}

func FindEpisodeByIdAndStatusNot(id string, status int) (*model.Episode, error) {
	movieId := util.ParseStringToUInt(id)
	episode := new(model.Episode)

	err := db.Model(model.Episode{}).Find(&episode, "episode_id = ? AND status <> ?", movieId, status).Error

	return episode, err
}

// SaveEpisode : Save Episode
func SaveEpisode(episode model.Episode) (*model.Episode, error) {
	if err := db.Model(model.Episode{}).Save(&episode).Error; err != nil {
		return nil, err
	}

	return &episode, nil
}
