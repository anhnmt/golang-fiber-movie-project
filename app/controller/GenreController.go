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

type GenreController struct {
	genreRepository *repository.GenreRepository
}

func NewGenreController() *GenreController {
	if genreController == nil {
		once = &sync.Once{}

		once.Do(func() {
			genreController = &GenreController{
				genreRepository: repository.NewGenreRepository(),
			}
			log.Println("Create new GenreController")
		})
	}

	return genreController
}

// FindAllGenres : Find all genres by Status = 1
func (obj *GenreController) FindAllGenres(c *fiber.Ctx) error {
	genres, err := obj.genreRepository.FindAllGenresByStatusNot(util.StatusDeleted)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", genres)
}

func (obj *GenreController) ClientFindAllGenres(c *fiber.Ctx) error {
	genres, err := obj.genreRepository.FindAllGenresByStatusNotIn([]int{util.StatusDraft, util.StatusDeleted})

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", genres)
}

// FindGenreById : Find genre by Genre_Id and Status = 1
func (obj *GenreController) FindGenreById(c *fiber.Ctx) error {
	genreId := c.Params("id")
	genre, err := obj.genreRepository.FindGenreByIdAndStatusNot(genreId, util.StatusDeleted)

	if err != nil || genre.GenreId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	return util.ResponseSuccess("Thành công", genre)
}

// CreateNewGenre : Create a new genre
func (obj *GenreController) CreateNewGenre(c *fiber.Ctx) error {
	genreRequest := new(request.GenreRequest)

	if err := c.BodyParser(genreRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	genre := model.Genre{
		Name:   genreRequest.Name,
		Slug:   genreRequest.Slug,
		Status: genreRequest.Status,
	}

	if _, err := obj.genreRepository.SaveGenre(genre); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// UpdateGenreById : Update genre by Genre_Id and Status = 1
func (obj *GenreController) UpdateGenreById(c *fiber.Ctx) error {
	genreId := c.Params("id")

	genre, err := obj.genreRepository.FindGenreByIdAndStatusNot(genreId, util.StatusDeleted)

	if err != nil || genre.GenreId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	genreRequest := new(request.GenreRequest)
	if err = c.BodyParser(genreRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	genre.Name = genreRequest.Name
	genre.Slug = genreRequest.Slug
	genre.Status = genreRequest.Status

	if _, err = obj.genreRepository.SaveGenre(*genre); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// DeleteGenreById : Delete genre by Genre_Id and Status = 1
func (obj *GenreController) DeleteGenreById(c *fiber.Ctx) error {
	genreId := c.Params("id")
	genre, err := obj.genreRepository.FindGenreByIdAndStatusNot(genreId, util.StatusDeleted)

	if err != nil || genre.GenreId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	genre.Status = util.StatusDeleted

	if _, err = obj.genreRepository.SaveGenre(*genre); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}
