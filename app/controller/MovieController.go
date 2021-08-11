package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/entity/response"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/mapper"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"github.com/xdorro/golang-fiber-base-project/pkg/validator"
	"log"
	"net/url"
	"sync"
)

type MovieController struct {
	movieRepository        *repository.MovieRepository
	countryRepository      *repository.CountryRepository
	genreRepository        *repository.GenreRepository
	movieGenreRepository   *repository.MovieGenreRepository
	movieCountryRepository *repository.MovieCountryRepository
}

func NewMovieController() *MovieController {
	if movieController == nil {
		once = &sync.Once{}

		once.Do(func() {
			if movieController == nil {
				movieController = &MovieController{
					movieRepository:        repository.NewMovieRepository(),
					countryRepository:      repository.NewCountryRepository(),
					genreRepository:        repository.NewGenreRepository(),
					movieGenreRepository:   repository.NewMovieGenreRepository(),
					movieCountryRepository: repository.NewMovieCountryRepository(),
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

func (obj *MovieController) ClientTopMovieSidebar(c *fiber.Ctx) error {
	status := []int{util.StatusDraft, util.StatusDeleted}

	movies, err := obj.movieRepository.FindAllTopMoviesByMovieTypeIdAndStatusNotInAndLimit(1, status, 5)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	series, err := obj.movieRepository.FindAllTopMoviesByMovieTypeIdAndStatusNotInAndLimit(2, status, 5)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := response.TopMovieSidebarResponse{
		Movies: *movies,
		Series: *series,
	}

	return util.ResponseSuccess("Thành công", result)
}

func (obj *MovieController) ClientTopMoviesBody(c *fiber.Ctx) error {
	status := []int{util.StatusDraft, util.StatusDeleted}

	cinemas, err := obj.movieRepository.FindAllTopMoviesByGenreSlugAndStatusNotInAndLimit("phim-chieu-rap", status, 8)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	movies, err := obj.movieRepository.FindAllTopMoviesByMovieTypeSlugAndStatusNotInAndLimit("phim-le", status, 8)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	series, err := obj.movieRepository.FindAllTopMoviesByMovieTypeSlugAndStatusNotInAndLimit("phim-bo", status, 6)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	cartoons, err := obj.movieRepository.FindAllTopMoviesByGenreSlugAndStatusNotInAndLimit("hoat-hinh", status, 6)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := response.TopMovieBodyResponse{
		Cinemas:  mapper.SearchMovies(cinemas),
		Movies:   mapper.SearchMovies(movies),
		Series:   mapper.SearchMovies(series),
		Cartoons: mapper.SearchMovies(cartoons),
	}

	return util.ResponseSuccess("Thành công", result)
}

func (obj *MovieController) ClientFindMovieDetail(c *fiber.Ctx) error {
	movieSlug := c.Params("movieSlug")
	status := []int{util.StatusDraft, util.StatusDeleted}

	movie, err := obj.movieRepository.FindMovieBySlugAndStatusNotInAndJoinMovieType(movieSlug, status)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	movieRelated, err := obj.movieRepository.FindAllMoviesRelatedByMovieIdAndStatusNotInAndLimit(movie.MovieId, status, 12)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	movieGenres, err := obj.movieGenreRepository.FindAllGenresByMovieIdAndStatusNotIn(movie.MovieId, status)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	movieCountries, err := obj.movieCountryRepository.FindAllCountriesByMovieIdAndStatusNotIn(movie.MovieId, status)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := mapper.ClientMovieDetail(movie, movieRelated, movieGenres, movieCountries)

	return util.ResponseSuccess("Thành công", result)
}

func (obj *MovieController) ClientFindMovieByName(c *fiber.Ctx) error {
	movieName, err := url.PathUnescape(c.Params("movieName"))
	if err != nil {
		log.Fatal(err)
	}

	status := []int{util.StatusDraft, util.StatusDeleted}

	movies, err := obj.movieRepository.FindAllMoviesByMovieNameAndStatusNotIn(movieName, status, 10)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := mapper.SearchMovies(movies)

	return util.ResponseSuccess("Thành công", result)
}

func (obj *MovieController) ClientFindMovieByMovieTypeSlug(c *fiber.Ctx) error {
	movieType, err := url.PathUnescape(c.Params("movieType"))
	if err != nil {
		log.Fatal(err)
	}

	status := []int{util.StatusDraft, util.StatusDeleted}

	movies, err := obj.movieRepository.FindAllMoviesByMovieTypeSlugAndStatusNotIn(movieType, status, 10)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := mapper.SearchMovies(movies)

	return util.ResponseSuccess("Thành công", result)
}

func (obj *MovieController) ClientFindMovieByMovieGenre(c *fiber.Ctx) error {
	movieGenre, err := url.PathUnescape(c.Params("movieGenre"))
	if err != nil {
		log.Fatal(err)
	}

	status := []int{util.StatusDraft, util.StatusDeleted}

	movies, err := obj.movieRepository.FindAllMoviesByGenreSlugAndStatusNotIn(movieGenre, status, 10)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := mapper.SearchMovies(movies)

	return util.ResponseSuccess("Thành công", result)
}

func (obj *MovieController) ClientFindMovieByMovieCountry(c *fiber.Ctx) error {
	movieCountry, err := url.PathUnescape(c.Params("movieCountry"))
	if err != nil {
		log.Fatal(err)
	}

	status := []int{util.StatusDraft, util.StatusDeleted}

	movies, err := obj.movieRepository.FindAllMoviesByCountrySlugAndStatusNotIn(movieCountry, status, 10)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := mapper.SearchMovies(movies)

	return util.ResponseSuccess("Thành công", result)
}

// FindMovieById : Find movie by Movie_Id and Status
func (obj *MovieController) FindMovieById(c *fiber.Ctx) error {
	movieId := c.Params("movieId")
	movie, err := obj.movieRepository.FindMovieByIdAndStatusNotJoinMovieType(movieId, util.StatusDeleted)

	if err != nil || movie.MovieId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	result := mapper.MovieDetail(movie)

	return util.ResponseSuccess("Thành công", result)
}

// CreateNewMovie : Create a new movie
func (obj *MovieController) CreateNewMovie(c *fiber.Ctx) error {
	movieForm := c.FormValue("movie")

	movieRequest := new(request.MovieRequest)

	if err := util.JSONUnmarshal([]byte(movieForm), movieRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	// Get first file from form field "poster":
	file, err := c.FormFile("poster")

	if file != nil {
		// 👷 Save file inside uploads folder under current working directory:
		poster := util.StoragePoster(file.Filename)

		if err = c.SaveFile(file, util.Storage(poster)); err != nil {
			return err
		}

		movieRequest.Poster = poster
	}

	newMovie := model.Movie{
		OriginName:  movieRequest.OriginName,
		Name:        movieRequest.Name,
		Slug:        movieRequest.Slug,
		Description: movieRequest.Description,
		Trailer:     movieRequest.Trailer,
		ImdbId:      movieRequest.ImdbId,
		Rating:      movieRequest.Rating,
		ReleaseDate: movieRequest.ReleaseDate,
		Runtime:     movieRequest.Runtime,
		Poster:      movieRequest.Poster,
		SeoTitle:    movieRequest.SeoTitle,
		SeoKeywords: movieRequest.SeoKeywords,
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
	movieId := c.Params("movieId")
	movie, err := obj.movieRepository.FindMovieByIdAndStatusNot(movieId, util.StatusDeleted)

	if err != nil || movie.MovieId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	movieForm := c.FormValue("movie")

	movieRequest := new(request.MovieRequest)

	if err = util.JSONUnmarshal([]byte(movieForm), movieRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	// Get first file from form field "poster":
	file, err := c.FormFile("poster")

	if file != nil {
		// 👷 Save file inside uploads folder under current working directory:
		poster := util.StoragePoster(file.Filename)

		if err = c.SaveFile(file, util.Storage(poster)); err != nil {
			return err
		}

		movieRequest.Poster = poster
	}

	movie.OriginName = movieRequest.OriginName
	movie.Name = movieRequest.Name
	movie.Slug = movieRequest.Slug
	movie.Description = movieRequest.Description
	movie.Trailer = movieRequest.Trailer
	movie.ImdbId = movieRequest.ImdbId
	movie.Rating = movieRequest.Rating
	movie.ReleaseDate = movieRequest.ReleaseDate
	movie.Runtime = movieRequest.Runtime
	movie.Poster = movieRequest.Poster
	movie.SeoTitle = movieRequest.SeoTitle
	movie.SeoKeywords = movieRequest.SeoKeywords
	movie.MovieTypeId = movieRequest.MovieTypeId
	movie.Status = movieRequest.Status

	if _, err = obj.movieRepository.UpdateMovie(movieId, *movie); err != nil {
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
	movieId := c.Params("movieId")
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
func (obj *MovieController) createMovieGenres(movieId *int64, newGenreIds *[]int64) error {
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
			if err = obj.movieGenreRepository.CreateMovieGenreByMovieId(movieGenres); err != nil {
				return util.ResponseError(err.Error(), nil)
			}
		}
	}

	return nil
}

// updateMovieGenres: Handler update movie genre
func (obj *MovieController) updateMovieGenres(movieId *int64, newGenreIds *[]int64) error {
	if len(*newGenreIds) > 0 {
		// Find all genres by movieId
		genres, err := obj.movieGenreRepository.FindAllGenresByMovieIdAndStatusNotIn(*movieId, []int{util.StatusDeleted, util.StatusDraft})
		if err != nil {
			return util.ResponseError(err.Error(), nil)
		}

		// Get list genreIds not exist and remove
		removeGenreIds := mapper.GetGenreIdsNotExistInNewGenreIds(*newGenreIds, *genres)
		if len(*removeGenreIds) > 0 {
			if err = obj.movieGenreRepository.RemoveMovieGenreByMovieIdAndGenreIds(*movieId, *removeGenreIds); err != nil {
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
func (obj *MovieController) createMovieCountries(movieId *int64, newCountryIds *[]int64) error {
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
			if err = obj.movieCountryRepository.CreateMovieCountryByMovieId(movieCountries); err != nil {
				return util.ResponseError(err.Error(), nil)
			}
		}
	}

	return nil
}

// updateMovieCountries: Handler update movie countries
func (obj *MovieController) updateMovieCountries(movieId *int64, newCountryIds *[]int64) error {
	if len(*newCountryIds) > 0 {
		// Find all genres by movieId
		countries, err := obj.movieCountryRepository.FindAllCountriesByMovieIdAndStatusNotIn(*movieId, []int{util.StatusDeleted, util.StatusDraft})
		if err != nil {
			return util.ResponseError(err.Error(), nil)
		}

		// Get list genreIds not exist and remove
		removeCountryIds := mapper.GetCountryIdsNotExistInNewCountryIds(*newCountryIds, *countries)
		if len(*removeCountryIds) > 0 {
			if err = obj.movieCountryRepository.RemoveMovieCountryByMovieIdAndCountryIds(*movieId, *removeCountryIds); err != nil {
				return util.ResponseError(err.Error(), nil)
			}
		}

		//  Get list genreIds new and create
		createCountryIds := mapper.GetNewCountryIdsNotExistInCountries(*newCountryIds, *countries)

		return obj.createMovieCountries(movieId, createCountryIds)
	}

	return nil
}
