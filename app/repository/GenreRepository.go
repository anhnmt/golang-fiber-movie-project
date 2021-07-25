package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

func FindAllGenresByStatusNot(status int) (*[]model.Genre, error) {
	genres := make([]model.Genre, 0)

	err := db.Model(model.Genre{}).
		Find(&genres, "status <> ?", status).Error

	return &genres, err
}

func FindAllGenresByGenreIdsInAndStatusNotIn(genreIds []uint, status []int) (*[]model.Genre, error) {
	genres := make([]model.Genre, 0)

	err := db.Model(model.Genre{}).
		Find(&genres, "genre_id IN ? AND status NOT IN ?", genreIds, status).Error

	return &genres, err
}

func FindGenreByIdAndStatusNot(id string, status int) (*model.Genre, error) {
	uid := util.ParseStringToUInt(id)

	var genre model.Genre
	err := db.Model(model.Genre{}).
		Where("genre_id = ? AND status <> ?", uid, status).
		Find(&genre).Error

	return &genre, err
}

func SaveGenre(genre model.Genre) (*model.Genre, error) {
	err := db.Model(model.Genre{}).
		Save(&genre).Error

	return &genre, err
}
