package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/dto"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

// FindAllGenres : Find all genres by Status = 1
func FindAllGenres(c *fiber.Ctx) error {
	genres, err := repository.FindAllGenresByStatusNot(util.STATUS_DELETED)

	if err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", genres)
}

// FindGenreById : Find genre by Genre_Id and Status = 1
func FindGenreById(c *fiber.Ctx) error {
	genreId := c.Params("id")
	genre, err := repository.FindGenreByIdAndStatusNot(genreId, util.STATUS_DELETED)

	if err != nil || genre.GenreId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	return util.ResponseSuccess(c, "Thành công", genre)
}

// CreateNewGenre : Create a new genre
func CreateNewGenre(c *fiber.Ctx) error {
	genreRequest := new(dto.GenreRequest)

	if err := c.BodyParser(genreRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	genre := model.Genre{
		Name:   genreRequest.Name,
		Slug:   genreRequest.Slug,
		Status: genreRequest.Status,
	}

	if _, err := repository.SaveGenre(genre); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// UpdateGenreById : Update genre by Genre_Id and Status = 1
func UpdateGenreById(c *fiber.Ctx) error {
	genreId := c.Params("id")

	genre, err := repository.FindGenreByIdAndStatusNot(genreId, util.STATUS_DELETED)

	if err != nil || genre.GenreId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	genreRequest := new(dto.GenreRequest)
	if err = c.BodyParser(genreRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	genre.Name = genreRequest.Name
	genre.Slug = genreRequest.Slug
	genre.Status = genreRequest.Status

	if _, err = repository.SaveGenre(*genre); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// DeleteGenreById : Delete genre by Genre_Id and Status = 1
func DeleteGenreById(c *fiber.Ctx) error {
	genreId := c.Params("id")
	genre, err := repository.FindGenreByIdAndStatusNot(genreId, util.STATUS_DELETED)

	if err != nil || genre.GenreId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	genre.Status = util.STATUS_DELETED

	if _, err = repository.SaveGenre(*genre); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}
