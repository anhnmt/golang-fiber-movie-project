package repository

import (
	"errors"
	model "github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

// FindAllTagsByStatus : Find tag by TagId and Status
func FindAllTagsByStatus(status int) (*[]model.Tag, error) {
	tags := make([]model.Tag, 0)

	if err := db.Find(&tags, "status = ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &tags, nil
}

func FindAllTagsByStatusNot(status int) (*[]model.Tag, error) {
	tags := make([]model.Tag, 0)

	if err := db.Find(&tags, "status <> ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &tags, nil
}

// FindTagByIdAndStatus : Find tag by TagId and Status
func FindTagByIdAndStatus(id string, status int) (*model.Tag, error) {
	uid := util.ParseStringToUInt(id)

	var tag model.Tag
	if err := db.Where("tag_id = ? AND status = ?", uid, status).Find(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &tag, nil
}

func FindTagByIdAndStatusNot(id string, status int) (*model.Tag, error) {
	uid := util.ParseStringToUInt(id)

	var tag model.Tag
	if err := db.Where("tag_id = ? AND status <> ?", uid, status).Find(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &tag, nil
}

func SaveTag(tag model.Tag) (*model.Tag, error) {
	if err := db.Save(&tag).Error; err != nil {
		return nil, err
	}

	return &tag, nil
}
