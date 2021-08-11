package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"gorm.io/gorm"
	"log"
	"sync"
)

type EpisodeDetailRepository struct {
	db *gorm.DB
}

func NewEpisodeDetailRepository() *EpisodeDetailRepository {
	if episodeDetailRepository == nil {
		once = &sync.Once{}

		once.Do(func() {
			if episodeDetailRepository == nil {
				episodeDetailRepository = &EpisodeDetailRepository{
					db: db,
				}
				log.Println("Create new EpisodeDetailRepository")
			}
		})
	}

	return episodeDetailRepository
}

// CreateEpisodeDetailsByEpisodeId : Create MovieGenre By MovieId
func (obj *EpisodeDetailRepository) CreateEpisodeDetailsByEpisodeId(episodeDetail model.EpisodeDetail) error {
	err := obj.db.
		Model(model.EpisodeDetail{}).
		Create(&episodeDetail).Error

	return err
}

func (obj *EpisodeDetailRepository) UpdateEpisodeDetail(episodeDetailId string, episodeDetail model.EpisodeDetail) (*model.EpisodeDetail, error) {
	err := db.Model(model.EpisodeDetail{}).
		Where("episode_detail_id = ?", episodeDetailId).
		Save(&episodeDetail).Error

	return &episodeDetail, err
}

func (obj *EpisodeDetailRepository) FindEpisodeDetailsByIdAndStatusNotIn(id string, status []int) (*[]model.EpisodeDetail, error) {
	episodeDetails := make([]model.EpisodeDetail, 0)

	err := obj.db.
		Model(&model.Episode{}).
		Select("episode_details.*").
		Joins("LEFT JOIN episode_details ON episode_details.episode_id = episodes.episode_id").
		Where("episodes.episode_id = ?", id).
		Where("episodes.status NOT IN ?", status).
		Where("episode_details.status NOT IN ?", status).
		Find(&episodeDetails).Error

	return &episodeDetails, err
}

func (obj *EpisodeDetailRepository) FindEpisodeDetailByEpisodeIdAndEpisodeDetailIdAndStatusNot(episodeId string, episodeDetailId string, status []int) (*model.EpisodeDetail, error) {
	episodeDetail := new(model.EpisodeDetail)

	err := obj.db.
		Model(&model.EpisodeDetail{}).
		Where("episode_id = ?", episodeId).
		Where("episode_detail_id = ?", episodeDetailId).
		Where("status NOT IN ?", status).
		Find(&episodeDetail).Error

	return episodeDetail, err
}

func (obj *EpisodeDetailRepository) FindEpisodeDetailByIdAndStatusNot(episodeDetailId string, status []int) (*model.EpisodeDetail, error) {
	episodeDetail := new(model.EpisodeDetail)

	err := obj.db.
		Model(&model.EpisodeDetail{}).
		Where("episode_detail_id = ?", episodeDetailId).
		Where("status NOT IN ?", status).
		Find(&episodeDetail).Error

	return episodeDetail, err
}

func (obj *EpisodeDetailRepository) UpdateStatusByEpisodeId(episodeId string, status int) error {
	err := obj.db.
		Model(&model.EpisodeDetail{}).
		Where("episode_id = ?", episodeId).
		Update("status", status).Error

	return err
}

func (obj *EpisodeDetailRepository) UpdateStatusByEpisodeDetailId(episodeDetailId string, status int) error {
	err := obj.db.
		Model(&model.EpisodeDetail{}).
		Where("episode_detail_id = ?", episodeDetailId).
		Update("status", status).Error

	return err
}
