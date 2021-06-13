package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/dto"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"github.com/xdorro/golang-fiber-base-project/platform/database"
	"gorm.io/gorm"
)

// FindAllCountries : Find all countries by Status = 1
func FindAllCountries(c *fiber.Ctx) error {
	db := database.GetDB()
	var countries []model.Country
	db.Find(&countries, "status = ?", 1)
	return util.ResponseSuccess(c, "Thành công", countries)
}

// FindCountryById : Find country by Country_Id and Status = 1
func FindCountryById(c *fiber.Ctx) error {
	var err error
	db := database.GetDB()

	countryId := c.Params("id")
	country := new(model.Country)

	country, err = findCountryByIdAndStatus(countryId, country, db)

	if err != nil || country.CountryId == 0 {
		return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
	}

	return util.ResponseSuccess(c, "Thành công", country)
}

// CreateNewCountry : Create a new country
func CreateNewCountry(c *fiber.Ctx) error {
	db := database.GetDB()

	countryRequest := new(dto.CountryRequest)

	if err := c.BodyParser(countryRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	country := model.Country{
		Name: countryRequest.Name,
		Slug: countryRequest.Slug,
	}

	if err := db.Create(&country).Error; err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// UpdateCountryById : Update country by Country_Id and Status = 1
func UpdateCountryById(c *fiber.Ctx) error {
	var err error
	db := database.GetDB()

	countryId := c.Params("id")
	countryRequest := new(dto.CountryRequest)
	country := new(model.Country)

	country, err = findCountryByIdAndStatus(countryId, country, db)

	if err != nil || country.CountryId == 0 {
		return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
	}

	if err := c.BodyParser(countryRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	country.Name = countryRequest.Name
	country.Slug = countryRequest.Slug

	if err := db.Save(&country).Error; err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// DeleteCountryById : Delete country by Country_Id and Status = 1
func DeleteCountryById(c *fiber.Ctx) error {
	var err error
	db := database.GetDB()

	country := new(model.Country)

	countryId := c.Params("id")

	country, err = findCountryByIdAndStatus(countryId, country, db)

	if err != nil || country.CountryId == 0 {
		return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
	}

	if result := db.Model(&country).Update("status", 0); result.Error != nil {
		return util.ResponseError(c, result.Error.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// findCountryByIdAndStatus : Find country by Country_Id and Status = 1
func findCountryByIdAndStatus(countryId string, country *model.Country, db *gorm.DB) (*model.Country, error) {
	if result := db.First(&country, "country_id = ? and status = ?", countryId, 1); result.Error != nil {
		return nil, result.Error
	}

	return country, nil
}
