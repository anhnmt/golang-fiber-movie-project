package repository

import (
	"errors"
	model "github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

func FindAllCountriesByStatusNot(status int) (*[]model.Country, error) {
	countries := make([]model.Country, 0)

	if err := db.Find(&countries, "status <> ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &countries, nil
}

func FindCountryByIdAndStatusNot(id string, status int) (*model.Country, error) {
	uid := util.ParseStringToUInt(id)

	var country model.Country
	if err := db.Where("country_id = ? AND status <> ?", uid, status).Find(&country).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &country, nil
}

func FindAllCountriesByCountryIdsInAndStatusNotIn(countryIds []uint, status []int) (*[]model.Country, error) {
	countries := make([]model.Country, 0)

	if err := db.Find(&countries, "country_id IN ? AND status NOT IN ?", countryIds, status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &countries, nil
}

func SaveCountry(country model.Country) (*model.Country, error) {
	if err := db.Save(&country).Error; err != nil {
		return nil, err
	}

	return &country, nil
}
