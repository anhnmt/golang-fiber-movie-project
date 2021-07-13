package validator

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
)

func ExistCountryIdInCountries(countryId uint, countries []model.Country) bool {
	for _, country := range countries {
		if country.CountryId == countryId {
			return true
		}
	}

	return false
}

func ExistCountryIdInCountryIds(countryId uint, countryIds []uint) bool {
	for _, country := range countryIds {
		if country == countryId {
			return true
		}
	}

	return false
}
