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
	"log"
	"sync"
)

type MovieController struct {
	movieRepository   *repository.MovieRepository
	countryRepository *repository.CountryRepository
	genreRepository   *repository.GenreRepository
}

func NewMovieController() *MovieController {
	if movieController == nil {
		once = &sync.Once{}

		once.Do(func() {
			if movieController == nil {
				movieController = &MovieController{
					movieRepository:   repository.NewMovieRepository(),
					countryRepository: repository.NewCountryRepository(),
					genreRepository:   repository.NewGenreRepository(),
				}
				log.Println("Create new MovieController")
			}
		})
	}

	return movieController
}

// FindAllMovies : Find all movies by Status
func (obj *MovieController) FindAllMovies(c *fiber.Ctx) error {
	movies, err := obj.movieRepository.FindAllMoviesByStatusNot(util.StatusDeleted)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := mapper.SearchMovies(movies)

	return util.ResponseSuccess("Thành công", result)
}

// FindMovieById : Find movie by Movie_Id and Status
func (obj *MovieController) FindMovieById(c *fiber.Ctx) error {
	movieId := c.Params("id")
	movie, err := obj.movieRepository.FindMovieByIdAndStatusNotJoinMovieType(movieId, util.StatusDeleted)

	if err != nil || movie.MovieId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	result := mapper.MovieDetail(movie)

	return util.ResponseSuccess("Thành công", result)
}

// CreateNewMovie : Create a new movie
func (obj *MovieController) CreateNewMovie(c *fiber.Ctx) error {
	movieRequest := new(request.MovieRequest)

	if err := c.BodyParser(movieRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	newMovie := model.Movie{
		Name:        movieRequest.Name,
		Slug:        movieRequest.Slug,
		MovieTypeId: movieRequest.MovieTypeId,
		Status:      movieRequest.Status,
	}

	movie, err := obj.movieRepository.SaveMovie(newMovie)
	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	if movie.MovieId == 0 {
		return util.ResponseBadRequest("Thêm mới thất bại", nil)
	}

	// Create Movie Genres
	if err = obj.createMovieGenres(&movie.MovieId, &movieRequest.GenreIds); err != nil {
		return err
	}

	// Create Movie Countries
	if err = obj.createMovieCountries(&movie.MovieId, &movieRequest.CountryIds); err != nil {
		return err
	}

	return util.ResponseSuccess("Thành công", nil)
}

// UpdateMovieById : Update movie by Movie_Id and Status
func (obj *MovieController) UpdateMovieById(c *fiber.Ctx) error {
	movieId := c.Params("id")
	movie, err := obj.movieRepository.FindMovieByIdAndStatusNot(movieId, util.StatusDeleted)

	if err != nil || movie.MovieId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	movieRequest := new(request.MovieRequest)
	if err = c.BodyParser(movieRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	movie.Name = movieRequest.Name
	movie.Slug = movieRequest.Slug
	movie.MovieTypeId = movieRequest.MovieTypeId
	movie.Status = movieRequest.Status

	if _, err = obj.movieRepository.SaveMovie(*movie); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	// Update movie genres
	if err = obj.updateMovieGenres(&movie.MovieId, &movieRequest.GenreIds); err != nil {
		return err
	}

	// Update movie countries
	if err = obj.updateMovieCountries(&movie.MovieId, &movieRequest.CountryIds); err != nil {
		return err
	}

	return util.ResponseSuccess("Thành công", nil)
}

// DeleteMovieById : Delete movie by Movie_Id and Status = 1
func (obj *MovieController) DeleteMovieById(c *fiber.Ctx) error {
	movieId := c.Params("id")
	movie, err := obj.movieRepository.FindMovieByIdAndStatusNot(movieId, util.StatusDeleted)

	if err != nil || movie.MovieId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	movie.Status = util.StatusDeleted

	if _, err = obj.movieRepository.SaveMovie(*movie); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// createMovieGenres: Handler create movie genre
func (obj *MovieController) createMovieGenres(movieId *uint, newGenreIds *[]uint) error {
	if len(*newGenreIds) > 0 {
		// Find genreIds in Genres
		genres, err := obj.genreRepository.FindAllGenresByGenreIdsInAndStatusNotIn(*newGenreIds, []int{util.StatusDeleted, util.StatusDraft})
		if err != nil {
			return util.ResponseError(err.Error(), nil)
		}

		// Validate genreIds
		for _, genreId := range *newGenreIds {
			if validator.ExistGenreIdInGenres(genreId, *genres) != true {
				return util.ResponseError(fmt.Sprintf("Không tìm thấy genre_id=%d", genreId), nil)
			}
		}

		// Create Movie Genre
		movieGenres := mapper.MovieGenres(movieId, newGenreIds)

		if len(movieGenres) > 0 {
			if err = repository.CreateMovieGenreByMovieId(movieGenres); err != nil {
				return util.ResponseError(err.Error(), nil)
			}
		}
	}

	return nil
}

// updateMovieGenres: Handler update movie genre
func (obj *MovieController) updateMovieGenres(movieId *uint, newGenreIds *[]uint) error {
	if len(*newGenreIds) > 0 {
		// Find all genres by movieId
		genres, err := repository.FindAllGenresByMovieIdAndStatusNotIn(*movieId, []int{util.StatusDeleted, util.StatusDraft})
		if err != nil {
			return util.ResponseError(err.Error(), nil)
		}

		// Get list genreIds not exist and remove
		removeGenreIds := mapper.GetGenreIdsNotExistInNewGenreIds(*newGenreIds, *genres)
		if len(*removeGenreIds) > 0 {
			if err = repository.RemoveMovieGenreByMovieIdAndGenreIds(*movieId, *removeGenreIds); err != nil {
				return util.ResponseError(err.Error(), nil)
			}
		}

		//  Get list genreIds new and create
		createGenreIds := mapper.GetNewGenreIdsNotExistInGenres(*newGenreIds, *genres)

		return obj.createMovieGenres(movieId, createGenreIds)
	}

	return nil
}

// createMovieCountries: Handler create movie genre
func (obj *MovieController) createMovieCountries(movieId *uint, newCountryIds *[]uint) error {
	if len(*newCountryIds) > 0 {
		// Find countryIds in Countries
		countries, err := obj.countryRepository.FindAllCountriesByCountryIdsInAndStatusNotIn(*newCountryIds, []int{util.StatusDeleted, util.StatusDraft})
		if err != nil {
			return util.ResponseError(err.Error(), nil)
		}

		// Validate countryIds
		for _, countryId := range *newCountryIds {
			if validator.ExistCountryIdInCountries(countryId, *countries) != true {
				return util.ResponseError(fmt.Sprintf("Không tìm thấy country_id=%d", countryId), nil)
			}
		}

		// Create Movie Country
		movieCountries := mapper.MovieCountries(movieId, newCountryIds)

		if len(movieCountries) > 0 {
			if err = repository.CreateMovieCountryByMovieId(movieCountries); err != nil {
				return util.ResponseError(err.Error(), nil)
			}
		}
	}

	return nil
}

// updateMovieCountries: Handler update movie countries
func (obj *MovieController) updateMovieCountries(movieId *uint, newCountryIds *[]uint) error {
	if len(*newCountryIds) > 0 {
		// Find all genres by movieId
		countries, err := repository.FindAllCountriesByMovieIdAndStatusNotIn(*movieId, []int{util.StatusDeleted, util.StatusDraft})
		if err != nil {
			return util.ResponseError(err.Error(), nil)
		}

		// Get list genreIds not exist and remove
		removeCountryIds := mapper.GetCountryIdsNotExistInNewCountryIds(*newCountryIds, *countries)
		if len(*removeCountryIds) > 0 {
			if err = repository.RemoveMovieCountryByMovieIdAndCountryIds(*movieId, *removeCountryIds); err != nil {
				return util.ResponseError(err.Error(), nil)
			}
		}

		//  Get list genreIds new and create
		createCountryIds := mapper.GetNewCountryIdsNotExistInCountries(*newCountryIds, *countries)

		return obj.createMovieCountries(movieId, createCountryIds)
	}

	return nil
}
