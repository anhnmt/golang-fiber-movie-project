package repository

import (
	"errors"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"gorm.io/gorm"
)

// CreateEpisodeDetailsByEpisodeId : Create MovieGenre By MovieId
func CreateEpisodeDetailsByEpisodeId(episodeDetails []model.EpisodeDetail) error {
	err := db.Model(model.EpisodeDetail{}).Create(&episodeDetails).Error

	return err
}

func FindEpisodeDetailsByIdAndStatusNot(id string, status []int) (*[]model.EpisodeDetail, error) {
	episodeDetails := make([]model.EpisodeDetail, 0)

	if err := db.
		Model(&model.Episode{}).
		Select("episode_details.*").
		Joins("LEFT JOIN episode_details ON episode_details.episode_id = episodes.episode_id").
		Where("episodes.episode_id = ?", id).
		Where("episodes.status NOT IN ?", status).
		Where("episode_details.status NOT IN ?", status).
		Find(&episodeDetails).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &episodeDetails, nil
}
