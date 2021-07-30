package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/dto"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
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
		Select("movies.movie_id, movies.name, movies.slug, movies.status, movies.movie_type_id, movie_types.name movie_type_name").
		Joins("LEFT JOIN movie_types on movies.movie_type_id = movie_types.movie_type_id").
		Where("movies.status <> ? AND movie_types.status <> ?", status, status).
		Find(&movies).Error

	return &movies, err
}

func (obj *MovieRepository) FindAllTopMoviesByMovieTypeIdAndStatusNotInAndLimit(movieTypeId uint, status []int, limit int) (*[]dto.SearchMovieDTO, error) {
	movies := make([]dto.SearchMovieDTO, 0)

	err := db.
		Model(&model.Movie{}).
		Select("movies.movie_id, movies.name, movies.slug, movies.status, movies.movie_type_id, movie_types.name movie_type_name").
		Joins("LEFT JOIN movie_types on movies.movie_type_id = movie_types.movie_type_id").
		Where("movies.status NOT IN ?", status).
		Where("movie_types.status NOT IN ?", status).
		Where("movies.movie_type_id = ?", movieTypeId).
		Order("movies.updated_at DESC").
		Limit(limit).
		Find(&movies).Error

	return &movies, err
}

func (obj *MovieRepository) FindMovieByIdAndStatusNot(id string, status int) (*model.Movie, error) {
	uid := util.ParseStringToUInt(id)

	var movie model.Movie
	err := db.Model(model.Movie{}).
		Where("movie_id = ? AND status <> ?", uid, status).
		Find(&movie).Error

	return &movie, err
}

func (obj *MovieRepository) FindMovieByIdAndStatusNotJoinMovieType(id string, status int) (*dto.MovieDetailDTO, error) {
	uid := util.ParseStringToUInt(id)

	var movie dto.MovieDetailDTO

	err := db.
		Model(&model.Movie{}).
		Select("movies.*, movie_types.name movie_type_name").
		Joins("LEFT JOIN movie_types on movies.movie_type_id = movie_types.movie_type_id").
		Where("movie_id = ? AND movies.status <> ? AND movie_types.status <> ?", uid, status, status).
		Find(&movie).Error

	return &movie, err
}

// SaveMovie : Save Movie
func (obj *MovieRepository) SaveMovie(movie model.Movie) (*model.Movie, error) {
	err := db.Save(&movie).Error

	return &movie, err
}
