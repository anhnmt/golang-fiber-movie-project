package repository

import (
	"errors"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

// FindAllGenresByStatus : Find genre by GenreId and Status = 1
func FindAllGenresByStatus(status int) (*[]model.Genre, error) {
	genres := make([]model.Genre, 0)

	if err := db.Find(&genres, "status = ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &genres, nil
}

func FindAllGenresByStatusNot(status int) (*[]model.Genre, error) {
	genres := make([]model.Genre, 0)

	if err := db.Find(&genres, "status <> ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &genres, nil
}

func FindAllGenresByStatusIn(status []int) (*[]model.Genre, error) {
	genres := make([]model.Genre, 0)

	if err := db.Find(&genres, "status IN ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &genres, nil
}

func FindAllGenresByStatusNotIn(status []int) (*[]model.Genre, error) {
	genres := make([]model.Genre, 0)

	if err := db.Find(&genres, "status NOT IN ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &genres, nil
}

// FindGenreByIdAndStatus : Find genre by GenreId and Status = 1
func FindGenreByIdAndStatus(id string, status int) (*model.Genre, error) {
	uid := util.ParseStringToUInt(id)

	var genre model.Genre
	if err := db.Where("genre_id = ? AND status = ?", uid, status).Find(&genre).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &genre, nil
}

func FindGenreByIdAndStatusNot(id string, status int) (*model.Genre, error) {
	uid := util.ParseStringToUInt(id)

	var genre model.Genre
	if err := db.Where("genre_id = ? AND status <> ?", uid, status).Find(&genre).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &genre, nil
}

func SaveGenre(genre model.Genre) (*model.Genre, error) {
	if err := db.Save(&genre).Error; err != nil {
		return nil, err
	}

	return &genre, nil
}
