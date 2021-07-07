package repository

import (
	"errors"
	model2 "github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

// FindAllMovieTypesByStatus : Find movieType by MovieTypeId and Status
func FindAllMovieTypesByStatus(status int) (*[]model2.MovieType, error) {
	movieTypes := make([]model2.MovieType, 0)

	if err := db.Find(&movieTypes, "status = ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &movieTypes, nil
}

func FindAllMovieTypesByStatusNot(status int) (*[]model2.MovieType, error) {
	movieTypes := make([]model2.MovieType, 0)

	if err := db.Find(&movieTypes, "status <> ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &movieTypes, nil
}

// FindMovieTypeByIdAndStatus : Find movieType by MovieTypeId and Status
func FindMovieTypeByIdAndStatus(id string, status int) (*model2.MovieType, error) {
	uid := util.ParseStringToUInt(id)

	var movieType model2.MovieType
	if err := db.Where("movie_type_id = ? AND status = ?", uid, status).Find(&movieType).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &movieType, nil
}

func FindMovieTypeByIdAndStatusNot(id string, status int) (*model2.MovieType, error) {
	uid := util.ParseStringToUInt(id)

	var movieType model2.MovieType
	if err := db.Where("movie_type_id = ? AND status <> ?", uid, status).Find(&movieType).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &movieType, nil
}

// SaveMovieType : Save Movie Type
func SaveMovieType(movieType model2.MovieType) (*model2.MovieType, error) {
	if err := db.Save(&movieType).Error; err != nil {
		return nil, err
	}

	return &movieType, nil
}
