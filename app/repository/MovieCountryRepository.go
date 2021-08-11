package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"gorm.io/gorm"
	"log"
	"sync"
)

type MovieCountryRepository struct {
	db *gorm.DB
}

func NewMovieCountryRepository() *MovieCountryRepository {
	if movieCountryRepository == nil {
		once = &sync.Once{}

		once.Do(func() {
			if movieCountryRepository == nil {
				movieCountryRepository = &MovieCountryRepository{
					db: db,
				}

				log.Println("Create new MovieCountryRepository")
			}
		})
	}

	return movieCountryRepository
}

// CreateMovieCountryByMovieId : Create MovieCountry By MovieId
func (obj *MovieCountryRepository) CreateMovieCountryByMovieId(movieCountries []model.MovieCountry) error {
	err := db.
		Model(&model.MovieCountry{}).
		Create(&movieCountries).Error

	return err
}

func (obj *MovieCountryRepository) RemoveMovieCountryByMovieIdAndCountryIds(movieId int64, countryIds []int64) error {
	err := db.
		Model(&model.MovieCountry{}).
		Where("movie_id = ? AND country_id IN ?", movieId, countryIds).
		Delete(&model.MovieCountry{}).Error

	return err
}

func (obj *MovieCountryRepository) FindAllCountriesByMovieIdAndStatusNotIn(movieId int64, status []int) (*[]model.Country, error) {
	countries := make([]model.Country, 0)

	err := db.
		Model(&model.Country{}).
		Select("countries.*").
		Joins("LEFT JOIN movie_countries ON countries.country_id = movie_countries.country_id").
		Where("movie_countries.movie_id = ? AND countries.status NOT IN ?", movieId, status).
		Find(&countries).Error

	return &countries, err
}
