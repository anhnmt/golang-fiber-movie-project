package repository

import (
	"errors"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

// FindAllMovieTypesByStatus : Find tag by MovieTypeId and Status
func FindAllMovieTypesByStatus(status int) (*[]model.MovieType, error) {
	tags := make([]model.MovieType, 0)

	if err := db.Find(&tags, "status = ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &tags, nil
}

func FindAllMovieTypesByStatusNot(status int) (*[]model.MovieType, error) {
	tags := make([]model.MovieType, 0)

	if err := db.Find(&tags, "status <> ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &tags, nil
}

// FindMovieTypeByIdAndStatus : Find tag by MovieTypeId and Status
func FindMovieTypeByIdAndStatus(id string, status int) (*model.MovieType, error) {
	uid := util.ParseStringToUInt(id)

	var tag model.MovieType
	if err := db.Where("movie_type_id = ? AND status = ?", uid, status).Find(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &tag, nil
}

func FindMovieTypeByIdAndStatusNot(id string, status int) (*model.MovieType, error) {
	uid := util.ParseStringToUInt(id)

	var tag model.MovieType
	if err := db.Where("movie_type_id = ? AND status <> ?", uid, status).Find(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &tag, nil
}

// SaveMovieType : Save Movie Type
func SaveMovieType(tag model.MovieType) (*model.MovieType, error) {
	if err := db.Save(&tag).Error; err != nil {
		return nil, err
	}

	return &tag, nil
}
