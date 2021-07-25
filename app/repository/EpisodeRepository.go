package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"gorm.io/gorm"
)

type EpisodeRepository struct {
	db *gorm.DB
}

func NewEpisodeRepository() *EpisodeRepository {
	return episodeRepository
}

func (obj *EpisodeRepository) FindAllEpisodesByMovieIdAndStatusNot(movieId string, status int) (*[]model.Episode, error) {
	episodes := make([]model.Episode, 0)

	err := db.Model(model.Episode{}).Find(&episodes, "movie_id = ? AND status <> ?", movieId, status).Error

	return &episodes, err
}

func (obj *EpisodeRepository) FindEpisodeByIdAndStatusNot(episodeId string, status int) (*model.Episode, error) {
	episode := new(model.Episode)

	err := db.Model(model.Episode{}).Find(&episode, "episode_id = ? AND status <> ?", episodeId, status).Error

	return episode, err
}

// SaveEpisode : Save Episode
func (obj *EpisodeRepository) SaveEpisode(episode model.Episode) (*model.Episode, error) {
	err := db.Model(model.Episode{}).Save(&episode).Error

	return &episode, err
}
