package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"gorm.io/gorm"
)

type EpisodeDetailRepository struct {
	db *gorm.DB
}

func NewEpisodeDetailRepository() *EpisodeDetailRepository {
	return episodeDetailRepository
}

// CreateEpisodeDetailsByEpisodeId : Create MovieGenre By MovieId
func (obj *EpisodeDetailRepository) CreateEpisodeDetailsByEpisodeId(episodeDetails []model.EpisodeDetail) error {
	err := obj.db.Model(model.EpisodeDetail{}).Create(&episodeDetails).Error

	return err
}

func (obj *EpisodeDetailRepository) FindEpisodeDetailsByIdAndStatusNot(id string, status []int) (*[]model.EpisodeDetail, error) {
	episodeDetails := make([]model.EpisodeDetail, 0)

	err := db.
		Model(&model.Episode{}).
		Select("episode_details.*").
		Joins("LEFT JOIN episode_details ON episode_details.episode_id = episodes.episode_id").
		Where("episodes.episode_id = ?", id).
		Where("episodes.status NOT IN ?", status).
		Where("episode_details.status NOT IN ?", status).
		Find(&episodeDetails).Error

	return &episodeDetails, err
}

func (obj *EpisodeDetailRepository) UpdateStatusByEpisodeId(episodeId string, status int) error {
	err := obj.db.Model(&model.Movie{}).
		Where("episode_id = ?", episodeId).
		Update("status", status).Error

	return err
}
