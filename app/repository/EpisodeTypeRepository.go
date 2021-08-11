package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"gorm.io/gorm"
	"log"
	"sync"
)

type EpisodeTypeRepository struct {
	db *gorm.DB
}

func NewEpisodeTypeRepository() *EpisodeTypeRepository {
	if episodeTypeRepository == nil {
		once = &sync.Once{}

		once.Do(func() {
			if episodeTypeRepository == nil {
				episodeTypeRepository = &EpisodeTypeRepository{
					db: db,
				}

				log.Println("Create new EpisodeTypeRepository")
			}
		})
	}

	return episodeTypeRepository
}

func (obj *EpisodeTypeRepository) FindAllEpisodeTypesByStatusNot(status int) (*[]model.EpisodeType, error) {
	episodeTypes := make([]model.EpisodeType, 0)

	err := obj.db.Model(model.EpisodeType{}).
		Find(&episodeTypes, "status <> ?", status).Error

	return &episodeTypes, err
}

func (obj *EpisodeTypeRepository) FindEpisodeTypeByIdAndStatusNot(id string, status int) (*model.EpisodeType, error) {
	var episodeType model.EpisodeType

	err := obj.db.Model(&model.EpisodeType{}).
		Where("episode_type_id = ? AND status <> ?", id, status).
		Find(&episodeType).Error

	return &episodeType, err
}

// SaveEpisodeType : Save Episode Type
func (obj *EpisodeTypeRepository) SaveEpisodeType(episodeType model.EpisodeType) (*model.EpisodeType, error) {
	err := obj.db.Model(&model.EpisodeType{}).
		Save(&episodeType).Error

	return &episodeType, err
}

func (obj *EpisodeTypeRepository) UpdateStatusByEpisodeTypeId(episodeTypeId int64, status int) error {
	err := obj.db.Model(&model.EpisodeType{}).
		Where("episode_type_id = ?", episodeTypeId).
		Update("status", status).Error

	return err
}
