package mapper

import (
	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
	"github.com/xdorro/golang-fiber-movie-project/pkg/validator"
)

func MovieCountries(movieId *int64, countryIds *[]int64) []model.MovieCountry {
	result := make([]model.MovieCountry, 0)

	for _, countryId := range *countryIds {
		mapper := MovieCountry(movieId, &countryId)
		result = append(result, *mapper)
	}

	return result
}

func MovieCountry(movieId *int64, countryId *int64) *model.MovieCountry {
	return &model.MovieCountry{
		MovieId:   *movieId,
		CountryId: *countryId,
	}
}

func GetCountryIdsNotExistInNewCountryIds(newCountryIds []int64, countries []model.Country) *[]int64 {
	result := make([]int64, 0)

	for _, country := range countries {
		if !validator.ExistCountryIdInCountryIds(country.CountryId, newCountryIds) {
			result = append(result, country.CountryId)
		}
	}

	return &result
}

func GetNewCountryIdsNotExistInCountries(countryIds []int64, countries []model.Country) *[]int64 {
	result := make([]int64, 0)

	for _, countryId := range countryIds {
		if !validator.ExistCountryIdInCountries(countryId, countries) {
			result = append(result, countryId)
		}
	}

	return &result
}
