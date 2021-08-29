package validator

import (
	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
)

func ExistCountryIdInCountries(countryId int64, countries []model.Country) bool {
	for _, country := range countries {
		if country.CountryId == countryId {
			return true
		}
	}

	return false
}

func ExistCountryIdInCountryIds(countryId int64, countryIds []int64) bool {
	for _, country := range countryIds {
		if country == countryId {
			return true
		}
	}

	return false
}
