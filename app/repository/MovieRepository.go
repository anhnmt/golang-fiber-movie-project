package repository

import (
	"errors"
	model2 "github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

// FindAllMoviesByStatus : Find movie by MovieId and Status
func FindAllMoviesByStatus(status int) (*[]model2.Movie, error) {
	movies := make([]model2.Movie, 0)

	if err := db.Find(&movies, "status = ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &movies, nil
}

func FindAllMoviesByStatusNot(status int) (*[]model2.Movie, error) {
	movies := make([]model2.Movie, 0)

	if err := db.Find(&movies, "status <> ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &movies, nil
}

// FindMovieByIdAndStatus : Find movie by MovieId and Status
func FindMovieByIdAndStatus(id string, status int) (*model2.Movie, error) {
	uid := util.ParseStringToUInt(id)

	var movie model2.Movie
	if err := db.Where("movie_id = ? AND status = ?", uid, status).Find(&movie).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &movie, nil
}

func FindMovieByIdAndStatusNot(id string, status int) (*model2.Movie, error) {
	uid := util.ParseStringToUInt(id)

	var movie model2.Movie
	if err := db.Where("movie_id = ? AND status <> ?", uid, status).Find(&movie).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &movie, nil
}

// SaveMovie : Save Movie
func SaveMovie(movie model2.Movie) (*model2.Movie, error) {
	if err := db.Save(&movie).Error; err != nil {
		return nil, err
	}

	return &movie, nil
}
