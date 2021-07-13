package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/mapper"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"github.com/xdorro/golang-fiber-base-project/pkg/validator"
)

// FindAllMovies : Find all movies by Status
func FindAllMovies(c *fiber.Ctx) error {
	movies, err := repository.FindAllMoviesByStatusNot(util.StatusDeleted)

	if err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	result := mapper.SearchMovies(movies)

	return util.ResponseSuccess(c, "Thành công", result)
}

// FindMovieById : Find movie by Movie_Id and Status
func FindMovieById(c *fiber.Ctx) error {
	movieId := c.Params("id")
	movie, err := repository.FindMovieByIdAndStatusNotJoinMovieType(movieId, util.StatusDeleted)

	if err != nil || movie.MovieId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	result := mapper.MovieDetail(movie)

	return util.ResponseSuccess(c, "Thành công", result)
}

// CreateNewMovie : Create a new movie
func CreateNewMovie(c *fiber.Ctx) error {
	movieRequest := new(request.MovieRequest)

	if err := c.BodyParser(movieRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	newMovie := model.Movie{
		Name:        movieRequest.Name,
		Slug:        movieRequest.Slug,
		MovieTypeId: movieRequest.MovieTypeId,
		Status:      movieRequest.Status,
	}

	movie, err := repository.SaveMovie(newMovie)
	if err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	if movie.MovieId == 0 {
		return util.ResponseBadRequest(c, "Thêm mới thất bại", nil)
	}

	// Create Movie Genres
	if err = createMovieGenres(c, &movie.MovieId, &movieRequest.GenreIds); err != nil {
		return err
	}

	// Create Movie Countries
	if err = createMovieCountries(c, &movie.MovieId, &movieRequest.CountryIds); err != nil {
		return err
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// UpdateMovieById : Update movie by Movie_Id and Status
func UpdateMovieById(c *fiber.Ctx) error {
	movieId := c.Params("id")
	movie, err := repository.FindMovieByIdAndStatusNot(movieId, util.StatusDeleted)

	if err != nil || movie.MovieId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	movieRequest := new(request.MovieRequest)
	if err = c.BodyParser(movieRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	movie.Name = movieRequest.Name
	movie.Slug = movieRequest.Slug
	movie.MovieTypeId = movieRequest.MovieTypeId
	movie.Status = movieRequest.Status

	if _, err = repository.SaveMovie(*movie); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	// Update movie genres
	if err = updateMovieGenres(c, &movie.MovieId, &movieRequest.GenreIds); err != nil {
		return err
	}

	// Update movie countries
	if err = updateMovieCountries(c, &movie.MovieId, &movieRequest.CountryIds); err != nil {
		return err
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// DeleteMovieById : Delete movie by Movie_Id and Status = 1
func DeleteMovieById(c *fiber.Ctx) error {
	movieId := c.Params("id")
	movie, err := repository.FindMovieByIdAndStatusNot(movieId, util.StatusDeleted)

	if err != nil || movie.MovieId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	movie.Status = util.StatusDeleted

	if _, err = repository.SaveMovie(*movie); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// createMovieGenres: Handler create movie genre
func createMovieGenres(c *fiber.Ctx, movieId *uint, newGenreIds *[]uint) error {
	if len(*newGenreIds) > 0 {
		// Find genreIds in Genres
		genres, err := repository.FindAllGenresByGenreIdsInAndStatusNotIn(*newGenreIds, []int{util.StatusDeleted, util.StatusDraft})
		if err != nil {
			return util.ResponseError(c, err.Error(), nil)
		}

		// Validate genreIds
		for _, genreId := range *newGenreIds {
			if validator.ExistGenreIdInGenres(genreId, *genres) != true {
				return util.ResponseError(c, fmt.Sprintf("Không tìm thấy genre_id=%d", genreId), nil)
			}
		}

		// Create Movie Genre
		movieGenres := mapper.MovieGenres(movieId, newGenreIds)

		if len(movieGenres) > 0 {
			if err = repository.CreateMovieGenreByMovieId(movieGenres); err != nil {
				return util.ResponseError(c, err.Error(), nil)
			}
		}
	}

	return nil
}

// updateMovieGenres: Handler update movie genre
func updateMovieGenres(c *fiber.Ctx, movieId *uint, newGenreIds *[]uint) error {
	if len(*newGenreIds) > 0 {
		// Find all genres by movieId
		genres, err := repository.FindAllGenresByMovieIdAndStatusNotIn(*movieId, []int{util.StatusDeleted, util.StatusDraft})
		if err != nil {
			return util.ResponseError(c, err.Error(), nil)
		}

		// Get list genreIds not exist and remove
		removeGenreIds := mapper.GetGenreIdsNotExistInNewGenreIds(*newGenreIds, *genres)
		if len(*removeGenreIds) > 0 {
			if err = repository.RemoveMovieGenreByMovieIdAndGenreIds(*movieId, *removeGenreIds); err != nil {
				return util.ResponseError(c, err.Error(), nil)
			}
		}

		//  Get list genreIds new and create
		createGenreIds := mapper.GetNewGenreIdsNotExistInGenres(*newGenreIds, *genres)

		return createMovieGenres(c, movieId, createGenreIds)
	}

	return nil
}

// createMovieCountries: Handler create movie genre
func createMovieCountries(c *fiber.Ctx, movieId *uint, newCountryIds *[]uint) error {
	if len(*newCountryIds) > 0 {
		// Find countryIds in Countries
		countries, err := repository.FindAllCountriesByCountryIdsInAndStatusNotIn(*newCountryIds, []int{util.StatusDeleted, util.StatusDraft})
		if err != nil {
			return util.ResponseError(c, err.Error(), nil)
		}

		// Validate countryIds
		for _, countryId := range *newCountryIds {
			if validator.ExistCountryIdInCountries(countryId, *countries) != true {
				return util.ResponseError(c, fmt.Sprintf("Không tìm thấy country_id=%d", countryId), nil)
			}
		}

		// Create Movie Country
		movieCountries := mapper.MovieCountries(movieId, newCountryIds)

		if len(movieCountries) > 0 {
			if err = repository.CreateMovieCountryByMovieId(movieCountries); err != nil {
				return util.ResponseError(c, err.Error(), nil)
			}
		}
	}

	return nil
}

// updateMovieCountries: Handler update movie countries
func updateMovieCountries(c *fiber.Ctx, movieId *uint, newCountryIds *[]uint) error {
	if len(*newCountryIds) > 0 {
		// Find all genres by movieId
		countries, err := repository.FindAllCountriesByMovieIdAndStatusNotIn(*movieId, []int{util.StatusDeleted, util.StatusDraft})
		if err != nil {
			return util.ResponseError(c, err.Error(), nil)
		}

		// Get list genreIds not exist and remove
		removeCountryIds := mapper.GetCountryIdsNotExistInNewCountryIds(*newCountryIds, *countries)
		if len(*removeCountryIds) > 0 {
			if err = repository.RemoveMovieCountryByMovieIdAndCountryIds(*movieId, *removeCountryIds); err != nil {
				return util.ResponseError(c, err.Error(), nil)
			}
		}

		//  Get list genreIds new and create
		createCountryIds := mapper.GetNewCountryIdsNotExistInCountries(*newCountryIds, *countries)

		return createMovieCountries(c, movieId, createCountryIds)
	}

	return nil
}
