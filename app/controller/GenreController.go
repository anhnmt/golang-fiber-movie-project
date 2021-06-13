package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/dto"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"github.com/xdorro/golang-fiber-base-project/platform/database"
	"gorm.io/gorm"
)

// FindAllGenres : Find all genres by Status = 1
func FindAllGenres(c *fiber.Ctx) error {
	db := database.GetDB()
	var genres []model.Genre
	db.Find(&genres, "status = ?", 1)
	return util.ResponseSuccess(c, "Thành công", genres)
}

// FindGenreById : Find genre by Genre_Id and Status = 1
func FindGenreById(c *fiber.Ctx) error {
	var err error
	db := database.GetDB()

	genreId := c.Params("id")
	genre := new(model.Genre)

	genre, err = findGenreByIdAndStatus(genreId, genre, db)

	if err != nil || genre.GenreId == 0 {
		return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
	}

	return util.ResponseSuccess(c, "Thành công", genre)
}

// CreateNewGenre : Create a new genre
func CreateNewGenre(c *fiber.Ctx) error {
	db := database.GetDB()

	genreRequest := new(dto.GenreRequest)

	if err := c.BodyParser(genreRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	genre := model.Genre{
		Name: genreRequest.Name,
		Slug: genreRequest.Slug,
	}

	if err := db.Create(&genre).Error; err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// UpdateGenreById : Update genre by Genre_Id and Status = 1
func UpdateGenreById(c *fiber.Ctx) error {
	var err error
	db := database.GetDB()

	genreId := c.Params("id")
	genreRequest := new(dto.GenreRequest)
	genre := new(model.Genre)

	genre, err = findGenreByIdAndStatus(genreId, genre, db)

	if err != nil || genre.GenreId == 0 {
		return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
	}

	if err := c.BodyParser(genreRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	genre.Name = genreRequest.Name
	genre.Slug = genreRequest.Slug

	if err := db.Save(&genre).Error; err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// DeleteGenreById : Delete genre by Genre_Id and Status = 1
func DeleteGenreById(c *fiber.Ctx) error {
	var err error
	db := database.GetDB()

	genre := new(model.Genre)

	genreId := c.Params("id")

	genre, err = findGenreByIdAndStatus(genreId, genre, db)

	if err != nil || genre.GenreId == 0 {
		return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
	}

	if result := db.Model(&genre).Update("status", 0); result.Error != nil {
		return util.ResponseError(c, result.Error.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// findGenreByIdAndStatus : Find genre by Genre_Id and Status = 1
func findGenreByIdAndStatus(genreId string, genre *model.Genre, db *gorm.DB) (*model.Genre, error) {
	if result := db.First(&genre, "genre_id = ? and status = ?", genreId, 1); result.Error != nil {
		return nil, result.Error
	}

	return genre, nil
}
