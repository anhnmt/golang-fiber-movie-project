package mapper

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/validator"
)

func MovieCountries(movieId *uint, countryIds *[]uint) []model.MovieCountry {
	result := make([]model.MovieCountry, 0)

	for _, countryId := range *countryIds {
		mapper := MovieCountry(movieId, &countryId)
		result = append(result, *mapper)
	}

	return result
}

func MovieCountry(movieId *uint, countryId *uint) *model.MovieCountry {
	return &model.MovieCountry{
		MovieId:   *movieId,
		CountryId: *countryId,
	}
}

func GetCountryIdsNotExistInNewCountryIds(newCountryIds []uint, countries []model.Country) *[]uint {
	result := make([]uint, 0)

	for _, country := range countries {
		if !validator.ExistCountryIdInCountryIds(country.CountryId, newCountryIds) {
			result = append(result, country.CountryId)
		}
	}

	return &result
}

func GetNewCountryIdsNotExistInCountries(countryIds []uint, countries []model.Country) *[]uint {
	result := make([]uint, 0)

	for _, countryId := range countryIds {
		if !validator.ExistCountryIdInCountries(countryId, countries) {
			result = append(result, countryId)
		}
	}

	return &result
}
