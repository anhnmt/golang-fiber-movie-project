package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
)

// CreateMovieGenreByMovieId : Create MovieGenre By MovieId
func CreateMovieGenreByMovieId(movieGenres []model.MovieGenre) error {
	err := db.
		Model(&model.Genre{}).
		Create(&movieGenres).Error

	return err
}

func RemoveMovieGenreByMovieIdAndGenreIds(movieId uint, genreIds []uint) error {
	err := db.
		Model(&model.Genre{}).
		Where("movie_id = ? AND genre_id IN ?", movieId, genreIds).
		Delete(&model.MovieGenre{}).Error

	return err
}

func FindAllGenresByMovieIdAndStatusNotIn(movieId uint, status []int) (*[]model.Genre, error) {
	genres := make([]model.Genre, 0)

	err := db.
		Model(&model.Genre{}).
		Select("genres.*").
		Joins("LEFT JOIN movie_genres ON genres.genre_id = movie_genres.genre_id").
		Where("movie_genres.movie_id = ? AND genres.status NOT IN ?", movieId, status).
		Find(&genres).Error

	return &genres, err
}
