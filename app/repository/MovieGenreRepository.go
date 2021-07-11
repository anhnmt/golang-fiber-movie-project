package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
)

// CreateMovieGenreByMovieId : Create MovieGenre By MovieId
func CreateMovieGenreByMovieId(movieGenres []model.MovieGenre) error {
	err := db.Create(&movieGenres).Error

	return err
}

//func UpdateStatusByMovieTypeId(movieTypeId uint, status int) error {
//	if err := db.Model(&model.Movie{}).
//		Where("movie_type_id = ?", movieTypeId).
//		Update("status", status).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
