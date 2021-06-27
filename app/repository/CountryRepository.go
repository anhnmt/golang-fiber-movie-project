package repository

import (
	"errors"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

// FindAllCountriesByStatus : Find country by CountryId and Status = 1
func FindAllCountriesByStatus(status int8) (*[]model.Country, error) {
	countries := make([]model.Country, 0)

	if err := db.Find(&countries, "status = ?", 1).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &countries, nil
}

// FindCountryByIdAndStatus : Find country by CountryId and Status = 1
func FindCountryByIdAndStatus(id string, status int8) (*model.Country, error) {
	uid := util.ParseStringToUInt(id)

	var country model.Country
	if err := db.Where(&model.Country{CountryId: uid, Status: status}).Find(&country).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &country, nil
}

func SaveCountry(country model.Country) (*model.Country, error) {
	if err := db.Save(&country).Error; err != nil {
		return nil, err
	}

	return &country, nil
}
