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

type MovieTypeController struct {
	movieTypeRepository *repository.MovieTypeRepository
}

func NewMovieTypeController() *MovieTypeController {
	if movieTypeController == nil {
		once = &sync.Once{}

		once.Do(func() {
			if movieTypeController == nil {
				movieTypeController = &MovieTypeController{
					movieTypeRepository: repository.NewMovieTypeRepository(),
				}
				log.Println("Create new MovieTypeController")
			}
		})
	}

	return movieTypeController
}

// FindAllMovieTypes : Find all moveTypes by Status = 1
func (obj *MovieTypeController) FindAllMovieTypes(c *fiber.Ctx) error {
	moveTypes, err := obj.movieTypeRepository.FindAllMovieTypesByStatusNot(util.StatusDeleted)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", moveTypes)
}

// FindMovieTypeById : Find moveType by MovieType_Id and Status = 1
func (obj *MovieTypeController) FindMovieTypeById(c *fiber.Ctx) error {
	moveTypeId := c.Params("movieTypeId")
	moveType, err := obj.movieTypeRepository.FindMovieTypeByIdAndStatusNot(moveTypeId, util.StatusDeleted)

	if err != nil || moveType.MovieTypeId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	return util.ResponseSuccess("Thành công", moveType)
}

// CreateNewMovieType : Create a new moveType
func (obj *MovieTypeController) CreateNewMovieType(c *fiber.Ctx) error {
	moveTypeRequest := new(request.MovieTypeRequest)

	if err := c.BodyParser(moveTypeRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	moveType := model.MovieType{
		Name:   moveTypeRequest.Name,
		Slug:   moveTypeRequest.Slug,
		Status: moveTypeRequest.Status,
	}

	if _, err := obj.movieTypeRepository.SaveMovieType(moveType); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// UpdateMovieTypeById : Update moveType by MovieType_Id and Status = 1
func (obj *MovieTypeController) UpdateMovieTypeById(c *fiber.Ctx) error {
	moveTypeId := c.Params("movieTypeId")
	moveType, err := obj.movieTypeRepository.FindMovieTypeByIdAndStatusNot(moveTypeId, util.StatusDeleted)

	if err != nil || moveType.MovieTypeId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	moveTypeRequest := new(request.MovieTypeRequest)
	if err = c.BodyParser(moveTypeRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	moveType.Name = moveTypeRequest.Name
	moveType.Slug = moveTypeRequest.Slug
	moveType.Status = moveTypeRequest.Status

	if _, err = obj.movieTypeRepository.SaveMovieType(*moveType); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// DeleteMovieTypeById : Delete moveType by MovieType_Id and Status = 1
func (obj *MovieTypeController) DeleteMovieTypeById(c *fiber.Ctx) error {
	moveTypeId := c.Params("movieTypeId")
	moveType, err := obj.movieTypeRepository.FindMovieTypeByIdAndStatusNot(moveTypeId, util.StatusDeleted)

	if err != nil || moveType.MovieTypeId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	moveType.Status = util.StatusDeleted

	// Update movieType status
	if _, err = obj.movieTypeRepository.SaveMovieType(*moveType); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	// Update movie status
	if err = obj.movieTypeRepository.UpdateStatusByMovieTypeId(moveType.MovieTypeId, moveType.Status); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}
