package repository

import (
	"errors"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"gorm.io/gorm"
)

// CreateMovieCountryByMovieId : Create MovieCountry By MovieId
func CreateMovieCountryByMovieId(movieCountries []model.MovieCountry) error {
	err := db.Create(&movieCountries).Error

	return err
}

func RemoveMovieCountryByMovieIdAndCountryIds(movieId uint, countryIds []uint) error {
	if err := db.Where("movie_id = ? AND country_id IN ?", movieId, countryIds).
		Delete(&model.MovieCountry{}).Error; err != nil {
		return err
	}

	return nil
}

func FindAllCountriesByMovieIdAndStatusNotIn(movieId uint, status []int) (*[]model.Country, error) {
	countries := make([]model.Country, 0)

	if err := db.
		Model(&model.Country{}).
		Select("countries.*").
		Joins("LEFT JOIN movie_countries ON countries.country_id = movie_countries.country_id").
		Where("movie_countries.movie_id = ? AND countries.status NOT IN ?", movieId, status).
		Find(&countries).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &countries, nil
}
