package repository

import (
	"errors"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

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

func FindAllGenresByGenreIdsInAndStatusNotIn(genreIds []uint, status []int) (*[]model.Genre, error) {
	genres := make([]model.Genre, 0)

	if err := db.Find(&genres, "genre_id IN ? AND status NOT IN ?", genreIds, status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &genres, nil
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
