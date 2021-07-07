package controller

import (
	"github.com/gofiber/fiber/v2"
	model2 "github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

// FindAllCountries : Find all countries by Status = 1
func FindAllCountries(c *fiber.Ctx) error {
	countries, err := repository.FindAllCountriesByStatusNot(util.STATUS_DELETED)

	if err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", countries)
}

// FindCountryById : Find country by Country_Id and Status = 1
func FindCountryById(c *fiber.Ctx) error {
	countryId := c.Params("id")
	country, err := repository.FindCountryByIdAndStatusNot(countryId, util.STATUS_DELETED)

	if err != nil || country.CountryId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	return util.ResponseSuccess(c, "Thành công", country)
}

// CreateNewCountry : Create a new country
func CreateNewCountry(c *fiber.Ctx) error {
	countryRequest := new(request.CountryRequest)

	if err := c.BodyParser(countryRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	country := model2.Country{
		Name:   countryRequest.Name,
		Slug:   countryRequest.Slug,
		Status: countryRequest.Status,
	}

	if _, err := repository.SaveCountry(country); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// UpdateCountryById : Update country by Country_Id and Status = 1
func UpdateCountryById(c *fiber.Ctx) error {
	countryId := c.Params("id")

	country, err := repository.FindCountryByIdAndStatusNot(countryId, util.STATUS_DELETED)

	if err != nil || country.CountryId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	countryRequest := new(request.CountryRequest)
	if err = c.BodyParser(countryRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	country.Name = countryRequest.Name
	country.Slug = countryRequest.Slug
	country.Status = countryRequest.Status

	if _, err = repository.SaveCountry(*country); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// DeleteCountryById : Delete country by Country_Id and Status = 1
func DeleteCountryById(c *fiber.Ctx) error {
	countryId := c.Params("id")
	country, err := repository.FindCountryByIdAndStatusNot(countryId, util.STATUS_DELETED)

	if err != nil || country.CountryId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	country.Status = util.STATUS_DELETED

	if _, err = repository.SaveCountry(*country); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}
