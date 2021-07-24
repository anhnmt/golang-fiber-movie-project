package repository

import (
	"errors"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

// FindAllTagsByStatus : Find tag by TagId and Status
func (obj TagRepository) FindAllTagsByStatus(status int) (*[]model.Tag, error) {
	tags := make([]model.Tag, 0)

	if err := obj.db.Find(&tags, "status = ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &tags, nil
}

func (obj TagRepository) FindAllTagsByStatusNot(status int) (*[]model.Tag, error) {
	tags := make([]model.Tag, 0)

	if err := obj.db.Find(&tags, "status <> ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &tags, nil
}

// FindTagByIdAndStatus : Find tag by TagId and Status
func (obj TagRepository) FindTagByIdAndStatus(id string, status int) (*model.Tag, error) {
	uid := util.ParseStringToUInt(id)

	var tag model.Tag
	if err := obj.db.Where("tag_id = ? AND status = ?", uid, status).Find(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &tag, nil
}

func (obj TagRepository) FindTagByIdAndStatusNot(id string, status int) (*model.Tag, error) {
	uid := util.ParseStringToUInt(id)

	var tag model.Tag
	if err := obj.db.Where("tag_id = ? AND status <> ?", uid, status).Find(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &tag, nil
}

func (obj TagRepository) SaveTag(tag model.Tag) (*model.Tag, error) {
	if err := obj.db.Save(&tag).Error; err != nil {
		return nil, err
	}

	return &tag, nil
}
