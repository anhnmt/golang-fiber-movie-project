package repository

import (
	"errors"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

func FindAllMovieTypesByStatusNot(status int) (*[]model.MovieType, error) {
	movieTypes := make([]model.MovieType, 0)

	if err := db.Find(&movieTypes, "status <> ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &movieTypes, nil
}

func FindMovieTypeByIdAndStatusNot(id string, status int) (*model.MovieType, error) {
	uid := util.ParseStringToUInt(id)

	var movieType model.MovieType
	if err := db.Where("movie_type_id = ? AND status <> ?", uid, status).Find(&movieType).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &movieType, nil
}

// SaveMovieType : Save Movie Type
func SaveMovieType(movieType model.MovieType) (*model.MovieType, error) {
	if err := db.Save(&movieType).Error; err != nil {
		return nil, err
	}

	return &movieType, nil
}

func UpdateStatusByMovieTypeId(movieTypeId uint, status int) error {
	if err := db.Model(&model.Movie{}).
		Where("movie_type_id = ?", movieTypeId).
		Update("status", status).Error; err != nil {
		return err
	}

	return nil
}
