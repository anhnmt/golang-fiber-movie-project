package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
	"github.com/xdorro/golang-fiber-movie-project/app/entity/request"
	"github.com/xdorro/golang-fiber-movie-project/app/repository"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
	"log"
	"sync"
)

type CountryController struct {
	countryRepository *repository.CountryRepository
}

func NewCountryController() *CountryController {
	if countryController == nil {
		once = &sync.Once{}

		once.Do(func() {
			countryController = &CountryController{
				countryRepository: repository.NewCountryRepository(),
			}
			log.Println("Create new CountryController")
		})
	}

	return countryController
}

// FindAllCountries : Find all countries by Status = 1
func (obj *CountryController) FindAllCountries(c *fiber.Ctx) error {
	countries, err := obj.countryRepository.FindAllCountriesByStatusNot(util.StatusDeleted)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", countries)
}

func (obj *CountryController) ClientFindAllCountries(c *fiber.Ctx) error {
	countries, err := obj.countryRepository.FindAllCountriesByStatusNotIn([]int{util.StatusDraft, util.StatusDeleted})

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", countries)
}

// FindCountryById : Find country by Country_Id and Status = 1
func (obj *CountryController) FindCountryById(c *fiber.Ctx) error {
	countryId := c.Params("countryId")
	country, err := obj.countryRepository.FindCountryByIdAndStatusNot(countryId, util.StatusDeleted)

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

	if _, err := obj.countryRepository.SaveCountry(country); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// UpdateCountryById : Update country by Country_Id and Status = 1
func (obj *CountryController) UpdateCountryById(c *fiber.Ctx) error {
	countryId := c.Params("countryId")

	country, err := obj.countryRepository.FindCountryByIdAndStatusNot(countryId, util.StatusDeleted)

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

	if _, err = obj.countryRepository.UpdateCountry(countryId, *country); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// DeleteCountryById : Delete country by Country_Id and Status = 1
func (obj *CountryController) DeleteCountryById(c *fiber.Ctx) error {
	countryId := c.Params("countryId")
	country, err := obj.countryRepository.FindCountryByIdAndStatusNot(countryId, util.StatusDeleted)

	if err != nil || country.CountryId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	country.Status = util.StatusDeleted

	if _, err = obj.countryRepository.UpdateCountry(countryId, *country); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

func (obj *CountryController) CheckIsExistCountrySlug(c *fiber.Ctx) error {
	var err error
	var country *model.Country

	slug := c.Query("slug")
	countryId := c.Query("country_id")
	status := []int{util.StatusDraft, util.StatusDeleted}

	if countryId != "" {
		country, err = obj.countryRepository.FindCountryBySlugAndCountryIdNotAndStatusNotIn(slug, countryId, status)
	} else {
		country, err = obj.countryRepository.FindCountryBySlugAndStatusNotIn(slug, status)
	}

	if err != nil || country.CountryId == 0 {
		return util.ResponseSuccess("Thành công", false)
	}

	return util.ResponseSuccess("Thành công", true)
}
