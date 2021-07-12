package repository

import (
	"errors"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"gorm.io/gorm"
)

// CreateMovieGenreByMovieId : Create MovieGenre By MovieId
func CreateMovieGenreByMovieId(movieGenres []model.MovieGenre) error {
	err := db.Create(&movieGenres).Error

	return err
}

func RemoveMovieGenreByMovieIdAndGenreIds(movieId uint, genreIds []uint) error {
	if err := db.Where("movie_id = ? AND genre_id IN ?", movieId, genreIds).
		Delete(&model.MovieGenre{}).Error; err != nil {
		return err
	}

	return nil
}

func FindAllGenresByMovieIdAndStatusNotIn(movieId uint, status []int) (*[]model.Genre, error) {
	genres := make([]model.Genre, 0)

	if err := db.
		Model(&model.Genre{}).
		Select("genres.*").
		Joins("LEFT JOIN movie_genres ON genres.genre_id = movie_genres.genre_id").
		Where("movie_genres.movie_id = ? AND genres.status NOT IN ?", movieId, status).
		Find(&genres).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &genres, nil
}
