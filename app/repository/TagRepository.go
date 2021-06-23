package repository

import (
	"errors"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

// FindAllTagsByStatus : Find tag by TagId and Status = 1
func FindAllTagsByStatus(status int8) (*[]model.Tag, error) {
	tags := make([]model.Tag, 0)

	if err := db.Find(&tags, "status = ?", 1).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &tags, nil
}

// FindTagByIdAndStatus : Find tag by TagId and Status = 1
func FindTagByIdAndStatus(id string, status int8) (*model.Tag, error) {
	var tag model.Tag

	uid := util.ParseStringToUInt(id)

	if err := db.Where(&model.Tag{TagId: uid, Status: status}).Find(&tag).Error; err != nil {
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
