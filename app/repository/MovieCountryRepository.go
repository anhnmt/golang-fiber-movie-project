package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
)

// CreateMovieCountryByMovieId : Create MovieCountry By MovieId
func CreateMovieCountryByMovieId(movieCountries []model.MovieCountry) error {
	err := db.
		Model(&model.Country{}).
		Create(&movieCountries).Error

	return err
}

func RemoveMovieCountryByMovieIdAndCountryIds(movieId uint, countryIds []uint) error {
	err := db.
		Model(&model.Country{}).
		Where("movie_id = ? AND country_id IN ?", movieId, countryIds).
		Delete(&model.MovieCountry{}).Error

	return err
}

func FindAllCountriesByMovieIdAndStatusNotIn(movieId uint, status []int) (*[]model.Country, error) {
	countries := make([]model.Country, 0)

	err := db.
		Model(&model.Country{}).
		Select("countries.*").
		Joins("LEFT JOIN movie_countries ON countries.country_id = movie_countries.country_id").
		Where("movie_countries.movie_id = ? AND countries.status NOT IN ?", movieId, status).
		Find(&countries).Error

	return &countries, err
}
