package repository

import (
	"github.com/xdorro/golang-fiber-movie-project/app/entity/dto"
	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
	"gorm.io/gorm"
	"log"
	"sync"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository() *MovieRepository {
	if movieRepository == nil {
		once = &sync.Once{}

		once.Do(func() {
			if movieRepository == nil {
				movieRepository = &MovieRepository{
					db: db,
				}
				log.Println("Create new MovieRepository")
			}
		})
	}

	return movieRepository
}

// FindAllMoviesByStatusNot : Find movie by MovieId and Status NOT
func (obj *MovieRepository) FindAllMoviesByStatusNot(status int) (*[]dto.SearchMovieDTO, error) {
	movies := make([]dto.SearchMovieDTO, 0)

	err := db.
		Model(&model.Movie{}).
		Select("movies.*, movie_types.name movie_type_name, movie_types.slug movie_type_slug").
		Joins("LEFT JOIN movie_types on movies.movie_type_id = movie_types.movie_type_id").
		Where("movies.status <> ? AND movie_types.status <> ?", status, status).
		Order("movies.movie_id DESC").
		Find(&movies).Error

	return &movies, err
}

func (obj *MovieRepository) FindAllTopMoviesByMovieTypeIdAndStatusNotInAndLimit(movieTypeId int64, status []int, limit int) (*[]dto.SearchMovieDTO, error) {
	movies := make([]dto.SearchMovieDTO, 0)

	err := db.
		Model(&model.Movie{}).
		Select("movies.*, movie_types.name movie_type_name, movie_types.slug movie_type_slug").
		Joins("LEFT JOIN movie_types on movies.movie_type_id = movie_types.movie_type_id").
		Where("movies.status NOT IN ?", status).
		Where("movie_types.status NOT IN ?", status).
		Where("movies.movie_type_id = ?", movieTypeId).
		Order("movies.updated_at DESC").
		Limit(limit).
		Find(&movies).Error

	return &movies, err
}

func (obj *MovieRepository) FindAllTopMoviesByGenreSlugAndStatusNotInAndLimit(slug string, status []int, limit int) (*[]dto.SearchMovieDTO, error) {
	movies := make([]dto.SearchMovieDTO, 0)

	err := db.
		Model(&model.Movie{}).
		Select("movies.*, movie_types.name movie_type_name, movie_types.slug movie_type_slug").
		Joins("JOIN movie_types on movies.movie_type_id = movie_types.movie_type_id").
		Joins("JOIN movie_genres on movie_genres.movie_id = movies.movie_id").
		Joins("JOIN genres on genres.genre_id = movie_genres.genre_id").
		Where("movies.status NOT IN ?", status).
		Where("genres.status NOT IN ?", status).
		Where("genres.slug LIKE ?", slug).
		Order("movies.updated_at DESC").
		Limit(limit).
		Find(&movies).Error

	return &movies, err
}

func (obj *MovieRepository) FindAllTopMoviesByMovieTypeSlugAndStatusNotInAndLimit(slug string, status []int, limit int) (*[]dto.SearchMovieDTO, error) {
	movies := make([]dto.SearchMovieDTO, 0)

	err := db.
		Model(&model.Movie{}).
		Select("movies.*, movie_types.name movie_type_name, movie_types.slug movie_type_slug").
		Joins("JOIN movie_types on movie_types.movie_type_id = movies.movie_type_id").
		Where("movies.status NOT IN ?", status).
		Where("movie_types.status NOT IN ?", status).
		Where("movie_types.slug LIKE ?", slug).
		Order("movies.updated_at DESC").
		Limit(limit).
		Find(&movies).Error

	return &movies, err
}

func (obj *MovieRepository) FindMovieByIdAndStatusNot(id string, status int) (*model.Movie, error) {
	uid := util.ParseStringToInt64(id)

	var movie model.Movie
	err := db.Model(model.Movie{}).
		Where("movie_id = ? AND status <> ?", uid, status).
		Find(&movie).Error

	return &movie, err
}

func (obj *MovieRepository) FindMovieByIdAndStatusNotJoinMovieType(movieId string, status int) (*dto.MovieDetailDTO, error) {
	var movie dto.MovieDetailDTO

	err := db.
		Model(&model.Movie{}).
		Select("movies.*, movie_types.name movie_type_name, movie_types.slug movie_type_slug").
		Joins("LEFT JOIN movie_types on movies.movie_type_id = movie_types.movie_type_id").
		Where("movie_id = ? AND movies.status <> ? AND movie_types.status <> ?", movieId, status, status).
		Find(&movie).Error

	return &movie, err
}

func (obj *MovieRepository) FindMovieBySlugAndStatusNotInAndJoinMovieType(movieSlug string, status []int) (*dto.MovieDetailDTO, error) {
	var movie dto.MovieDetailDTO

	err := db.
		Model(&model.Movie{}).
		Select("movies.*, movie_types.name movie_type_name, movie_types.slug movie_type_slug").
		Joins("LEFT JOIN movie_types on movies.movie_type_id = movie_types.movie_type_id").
		Where("movies.slug LIKE ? AND movies.status NOT IN ? AND movie_types.status NOT IN ?", movieSlug, status, status).
		Find(&movie).Error

	return &movie, err
}

func (obj *MovieRepository) CountAllMoviesByMovieNameAndStatusNotIn(movieName string, status []int) (int64, error) {
	var count int64
	movieName = "%" + movieName + "%"

	err := db.
		Model(&model.Movie{}).
		Select("movies.*, movie_types.name movie_type_name, movie_types.slug movie_type_slug").
		Joins("JOIN movie_types on movie_types.movie_type_id = movies.movie_type_id").
		Where("movies.status NOT IN ?", status).
		Where("movie_types.status NOT IN ?", status).
		Where("movies.name LIKE ? OR movies.origin_name LIKE ?", movieName, movieName).
		Order("movies.updated_at DESC").
		Count(&count).Error

	return count, err
}

func (obj *MovieRepository) FindAllMoviesByMovieNameAndStatusNotIn(movieName string, status []int, page int, pageSize int) (*[]dto.SearchMovieDTO, error) {
	movies := make([]dto.SearchMovieDTO, 0)
	movieName = "%" + movieName + "%"
	offset := (page - 1) * pageSize

	err := db.
		Model(&model.Movie{}).
		Select("movies.*, movie_types.name movie_type_name, movie_types.slug movie_type_slug").
		Joins("JOIN movie_types on movie_types.movie_type_id = movies.movie_type_id").
		Where("movies.status NOT IN ?", status).
		Where("movie_types.status NOT IN ?", status).
		Where("movies.name LIKE ? OR movies.origin_name LIKE ?", movieName, movieName).
		Order("movies.updated_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&movies).Error

	return &movies, err
}

func (obj *MovieRepository) CountAllMoviesByMovieTypeSlugAndStatusNotIn(movieType string, status []int) (int64, error) {
	var count int64
	movieType = "%" + movieType + "%"

	err := db.
		Model(&model.Movie{}).
		Select("movies.*, movie_types.name movie_type_name, movie_types.slug movie_type_slug").
		Joins("JOIN movie_types on movie_types.movie_type_id = movies.movie_type_id").
		Where("movies.status NOT IN ?", status).
		Where("movie_types.status NOT IN ?", status).
		Where("movie_types.slug LIKE ?", movieType).
		Order("movies.updated_at DESC").
		Count(&count).Error

	return count, err
}

func (obj *MovieRepository) FindAllMoviesByMovieTypeSlugAndStatusNotIn(movieType string, status []int, page int, pageSize int) (*[]dto.SearchMovieDTO, error) {
	movies := make([]dto.SearchMovieDTO, 0)
	movieType = "%" + movieType + "%"
	offset := (page - 1) * pageSize

	err := db.
		Model(&model.Movie{}).
		Select("movies.*, movie_types.name movie_type_name, movie_types.slug movie_type_slug").
		Joins("JOIN movie_types on movie_types.movie_type_id = movies.movie_type_id").
		Where("movies.status NOT IN ?", status).
		Where("movie_types.status NOT IN ?", status).
		Where("movie_types.slug LIKE ?", movieType).
		Order("movies.updated_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&movies).Error

	return &movies, err
}

func (obj *MovieRepository) CountAllMoviesByGenreSlugAndStatusNotIn(movieGenre string, status []int) (int64, error) {
	var count int64
	movieGenre = "%" + movieGenre + "%"

	err := db.
		Model(&model.Movie{}).
		Select("movies.*, movie_types.name movie_type_name, movie_types.slug movie_type_slug").
		Joins("JOIN movie_types on movie_types.movie_type_id = movies.movie_type_id").
		Joins("JOIN movie_genres on movie_genres.movie_id = movies.movie_id").
		Joins("JOIN genres on genres.genre_id = movie_genres.genre_id").
		Where("movies.status NOT IN ?", status).
		Where("movie_types.status NOT IN ?", status).
		Where("genres.status NOT IN ?", status).
		Where("genres.slug LIKE ?", movieGenre).
		Order("movies.updated_at DESC").
		Count(&count).Error

	return count, err
}

func (obj *MovieRepository) FindAllMoviesByGenreSlugAndStatusNotIn(movieGenre string, status []int, page int, pageSize int) (*[]dto.SearchMovieDTO, error) {
	movies := make([]dto.SearchMovieDTO, 0)
	movieGenre = "%" + movieGenre + "%"
	offset := (page - 1) * pageSize

	err := db.
		Model(&model.Movie{}).
		Select("movies.*, movie_types.name movie_type_name, movie_types.slug movie_type_slug").
		Joins("JOIN movie_types on movie_types.movie_type_id = movies.movie_type_id").
		Joins("JOIN movie_genres on movie_genres.movie_id = movies.movie_id").
		Joins("JOIN genres on genres.genre_id = movie_genres.genre_id").
		Where("movies.status NOT IN ?", status).
		Where("movie_types.status NOT IN ?", status).
		Where("genres.status NOT IN ?", status).
		Where("genres.slug LIKE ?", movieGenre).
		Order("movies.updated_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&movies).Error

	return &movies, err
}

func (obj *MovieRepository) CountAllMoviesByCountrySlugAndStatusNotIn(movieCountry string, status []int) (int64, error) {
	var count int64
	movieCountry = "%" + movieCountry + "%"

	err := db.
		Model(&model.Movie{}).
		Select("movies.*, movie_types.name movie_type_name, movie_types.slug movie_type_slug").
		Joins("JOIN movie_types on movie_types.movie_type_id = movies.movie_type_id").
		Joins("JOIN movie_countries on movie_countries.movie_id = movies.movie_id").
		Joins("JOIN countries on countries.country_id = movie_countries.country_id").
		Where("movies.status NOT IN ?", status).
		Where("movie_types.status NOT IN ?", status).
		Where("countries.status NOT IN ?", status).
		Where("countries.slug LIKE ?", movieCountry).
		Order("movies.updated_at DESC").
		Count(&count).Error

	return count, err
}

func (obj *MovieRepository) FindAllMoviesByCountrySlugAndStatusNotIn(movieCountry string, status []int, page int, pageSize int) (*[]dto.SearchMovieDTO, error) {
	movies := make([]dto.SearchMovieDTO, 0)
	movieCountry = "%" + movieCountry + "%"
	offset := (page - 1) * pageSize

	err := db.
		Model(&model.Movie{}).
		Select("movies.*, movie_types.name movie_type_name, movie_types.slug movie_type_slug").
		Joins("JOIN movie_types on movie_types.movie_type_id = movies.movie_type_id").
		Joins("JOIN movie_countries on movie_countries.movie_id = movies.movie_id").
		Joins("JOIN countries on countries.country_id = movie_countries.country_id").
		Where("movies.status NOT IN ?", status).
		Where("movie_types.status NOT IN ?", status).
		Where("countries.status NOT IN ?", status).
		Where("countries.slug LIKE ?", movieCountry).
		Order("movies.updated_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&movies).Error

	return &movies, err
}

func (obj *MovieRepository) FindAllMoviesRelatedByMovieIdAndStatusNotInAndLimit(movieId int64, status []int, limit int) (*[]dto.SearchMovieDTO, error) {
	movies := make([]dto.SearchMovieDTO, 0)

	err := db.
		Model(&model.Movie{}).
		Select("DISTINCT movies.*, movie_types.name movie_type_name, movie_types.slug movie_type_slug").
		Joins("JOIN movie_genres on movie_genres.movie_id = ?", movieId).
		Joins("JOIN genres on genres.genre_id = movie_genres.genre_id").
		Joins("JOIN movie_countries on movie_countries.movie_id = ?", movieId).
		Joins("JOIN countries on countries.country_id = movie_countries.country_id").
		Joins("JOIN movie_types on movies.movie_type_id = movie_types.movie_type_id").
		Where("movies.status NOT IN ?", status).
		Where("genres.status NOT IN ?", status).
		Where("countries.status NOT IN ?", status).
		Where("movies.movie_id <> ?", movieId).
		Order("movies.updated_at DESC").
		Limit(limit).
		Find(&movies).Error

	return &movies, err
}

// SaveMovie : Save Movie
func (obj *MovieRepository) SaveMovie(movie model.Movie) (*model.Movie, error) {
	err := db.
		Model(&model.Movie{}).
		Save(&movie).Error

	return &movie, err
}

func (obj *MovieRepository) UpdateMovie(movieId string, movie model.Movie) (*model.Movie, error) {
	err := db.Model(model.Movie{}).
		Where("movie_id = ?", movieId).
		Save(&movie).Error

	return &movie, err
}
