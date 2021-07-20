package repository

import "github.com/xdorro/golang-fiber-base-project/app/entity/model"

// CreateEpisodeDetailsByEpisodeId : Create MovieGenre By MovieId
func CreateEpisodeDetailsByEpisodeId(episodeDetails []model.EpisodeDetail) error {
	err := db.Model(model.EpisodeDetail{}).Create(&episodeDetails).Error

	return err
}
