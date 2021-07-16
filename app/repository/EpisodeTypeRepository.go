package repository

import (
	"errors"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

func FindAllEpisodeTypesByStatusNot(status int) (*[]model.EpisodeType, error) {
	episodeTypes := make([]model.EpisodeType, 0)

	if err := db.Model(model.EpisodeType{}).Find(&episodeTypes, "status <> ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &episodeTypes, nil
}

func FindEpisodeTypeByIdAndStatusNot(id string, status int) (*model.EpisodeType, error) {
	uid := util.ParseStringToUInt(id)

	var episodeType model.EpisodeType
	if err := db.Where("episode_type_id = ? AND status <> ?", uid, status).Find(&episodeType).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &episodeType, nil
}

// SaveEpisodeType : Save Episode Type
func SaveEpisodeType(episodeType model.EpisodeType) (*model.EpisodeType, error) {
	if err := db.Save(&episodeType).Error; err != nil {
		return nil, err
	}

	return &episodeType, nil
}

func UpdateStatusByEpisodeTypeId(episodeTypeId uint, status int) error {
	if err := db.Model(&model.EpisodeDetail{}).
		Where("episode_type_id = ?", episodeTypeId).
		Update("status", status).Error; err != nil {
		return err
	}

	return nil
}
