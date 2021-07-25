package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"log"
	"sync"
)

type CountryController struct {
	//countryRepository *repository.CountryRepository
}

func NewCountryController() *CountryController {
	if countryController == nil {
		once = &sync.Once{}

		once.Do(func() {
			countryController = &CountryController{
				//countryRepository: repository.NewCountryRepository(),
			}
			log.Println("Create new CountryController")
		})
	}

	return countryController
}

// FindAllCountries : Find all countries by Status = 1
func (obj *CountryController) FindAllCountries(c *fiber.Ctx) error {
	countries, err := repository.FindAllCountriesByStatusNot(util.StatusDeleted)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", countries)
}

// FindCountryById : Find country by Country_Id and Status = 1
func (obj *CountryController) FindCountryById(c *fiber.Ctx) error {
	countryId := c.Params("id")
	country, err := repository.FindCountryByIdAndStatusNot(countryId, util.StatusDeleted)

	if err != nil || country.CountryId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	return util.ResponseSuccess("Thành công", country)
}

// CreateNewCountry : Create a new country
func (obj *CountryController) CreateNewCountry(c *fiber.Ctx) error {
	countryRequest := new(request.CountryRequest)

	if err := c.BodyParser(countryRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	country := model.Country{
		Name:   countryRequest.Name,
		Slug:   countryRequest.Slug,
		Status: countryRequest.Status,
	}

	if _, err := repository.SaveCountry(country); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// UpdateCountryById : Update country by Country_Id and Status = 1
func (obj *CountryController) UpdateCountryById(c *fiber.Ctx) error {
	countryId := c.Params("id")

	country, err := repository.FindCountryByIdAndStatusNot(countryId, util.StatusDeleted)

	if err != nil || country.CountryId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	countryRequest := new(request.CountryRequest)
	if err = c.BodyParser(countryRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	country.Name = countryRequest.Name
	country.Slug = countryRequest.Slug
	country.Status = countryRequest.Status

	if _, err = repository.SaveCountry(*country); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// DeleteCountryById : Delete country by Country_Id and Status = 1
func (obj *CountryController) DeleteCountryById(c *fiber.Ctx) error {
	countryId := c.Params("id")
	country, err := repository.FindCountryByIdAndStatusNot(countryId, util.StatusDeleted)

	if err != nil || country.CountryId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	country.Status = util.StatusDeleted

	if _, err = repository.SaveCountry(*country); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}
