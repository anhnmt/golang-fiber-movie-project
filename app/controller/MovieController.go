package controller

import (
	"github.com/gofiber/fiber/v2"
	model2 "github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

// FindAllMovies : Find all movies by Status
func FindAllMovies(c *fiber.Ctx) error {
	movies, err := repository.FindAllMoviesByStatusNot(util.STATUS_DELETED)

	if err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", movies)
}

// FindMovieById : Find movie by Movie_Id and Status
func FindMovieById(c *fiber.Ctx) error {
	movieId := c.Params("id")
	movie, err := repository.FindMovieByIdAndStatusNot(movieId, util.STATUS_DELETED)

	if err != nil || movie.MovieId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	return util.ResponseSuccess(c, "Thành công", movie)
}

// CreateNewMovie : Create a new movie
func CreateNewMovie(c *fiber.Ctx) error {
	movieRequest := new(request.MovieRequest)

	if err := c.BodyParser(movieRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	movie := model2.Movie{
		Name:      movieRequest.Name,
		Slug:      movieRequest.Slug,
		MovieType: movieRequest.MovieType,
		Status:    movieRequest.Status,
	}

	if _, err := repository.SaveMovie(movie); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// UpdateMovieById : Update movie by Movie_Id and Status
func UpdateMovieById(c *fiber.Ctx) error {
	movieId := c.Params("id")
	movie, err := repository.FindMovieByIdAndStatusNot(movieId, util.STATUS_DELETED)

	if err != nil || movie.MovieId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	movieRequest := new(request.MovieRequest)
	if err = c.BodyParser(movieRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	movie.Name = movieRequest.Name
	movie.Slug = movieRequest.Slug
	movie.MovieType = movieRequest.MovieType
	movie.Status = movieRequest.Status

	if _, err = repository.SaveMovie(*movie); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// DeleteMovieById : Delete movie by Movie_Id and Status = 1
func DeleteMovieById(c *fiber.Ctx) error {
	movieId := c.Params("id")
	movie, err := repository.FindMovieByIdAndStatusNot(movieId, util.STATUS_DELETED)

	if err != nil || movie.MovieId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	movie.Status = util.STATUS_DELETED

	if _, err = repository.SaveMovie(*movie); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}
