package repository

import (
	"errors"
	model2 "github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

// FindAllGenresByStatus : Find genre by GenreId and Status = 1
func FindAllGenresByStatus(status int) (*[]model2.Genre, error) {
	genres := make([]model2.Genre, 0)

	if err := db.Find(&genres, "status = ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &genres, nil
}

func FindAllGenresByStatusNot(status int) (*[]model2.Genre, error) {
	genres := make([]model2.Genre, 0)

	if err := db.Find(&genres, "status <> ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &genres, nil
}

func FindAllGenresByStatusIn(status []int) (*[]model2.Genre, error) {
	genres := make([]model2.Genre, 0)

	if err := db.Find(&genres, "status IN ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &genres, nil
}

func FindAllGenresByStatusNotIn(status []int) (*[]model2.Genre, error) {
	genres := make([]model2.Genre, 0)

	if err := db.Find(&genres, "status NOT IN ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &genres, nil
}

// FindGenreByIdAndStatus : Find genre by GenreId and Status = 1
func FindGenreByIdAndStatus(id string, status int) (*model2.Genre, error) {
	uid := util.ParseStringToUInt(id)

	var genre model2.Genre
	if err := db.Where("genre_id = ? AND status = ?", uid, status).Find(&genre).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &genre, nil
}

func FindGenreByIdAndStatusNot(id string, status int) (*model2.Genre, error) {
	uid := util.ParseStringToUInt(id)

	var genre model2.Genre
	if err := db.Where("genre_id = ? AND status <> ?", uid, status).Find(&genre).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &genre, nil
}

func SaveGenre(genre model2.Genre) (*model2.Genre, error) {
	if err := db.Save(&genre).Error; err != nil {
		return nil, err
	}

	return &genre, nil
}
