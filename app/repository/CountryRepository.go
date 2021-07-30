package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
	"log"
	"sync"
)

type CountryRepository struct {
	db *gorm.DB
}

func NewCountryRepository() *CountryRepository {
	if countryRepository == nil {
		once = &sync.Once{}

		once.Do(func() {
			if countryRepository == nil {
				countryRepository = &CountryRepository{
					db: db,
				}
				log.Println("Create new CountryRepository")
			}
		})
	}

	return countryRepository
}

func (obj *CountryRepository) FindAllCountriesByStatusNot(status int) (*[]model.Country, error) {
	countries := make([]model.Country, 0)

	err := db.Model(model.Country{}).
		Find(&countries, "status <> ?", status).Error

	return &countries, err
}

func (obj *CountryRepository) FindAllCountriesByStatusNotIn(status []int) (*[]model.Country, error) {
	countries := make([]model.Country, 0)

	err := db.Model(model.Country{}).
		Find(&countries, "status NOT IN ?", status).Error

	return &countries, err
}

func (obj *CountryRepository) FindCountryByIdAndStatusNot(id string, status int) (*model.Country, error) {
	uid := util.ParseStringToUInt(id)

	var country model.Country
	err := db.Model(model.Country{}).
		Where("country_id = ? AND status <> ?", uid, status).
		Find(&country).Error

	return &country, err
}

func (obj *CountryRepository) FindAllCountriesByCountryIdsInAndStatusNotIn(countryIds []uint, status []int) (*[]model.Country, error) {
	countries := make([]model.Country, 0)

	err := db.Model(model.Country{}).
		Find(&countries, "country_id IN ? AND status NOT IN ?", countryIds, status).Error

	return &countries, err
}

func (obj *CountryRepository) SaveCountry(country model.Country) (*model.Country, error) {
	err := db.Model(model.Country{}).
		Save(&country).Error

	return &country, err
}
