package repository

import (
	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
	"gorm.io/gorm"
	"log"
	"sync"
)

type MovieTypeRepository struct {
	db *gorm.DB
}

func NewMovieTypeRepository() *MovieTypeRepository {
	if movieTypeRepository == nil {
		once = &sync.Once{}

		once.Do(func() {
			if movieTypeRepository == nil {
				movieTypeRepository = &MovieTypeRepository{
					db: db,
				}
				log.Println("Create new MovieTypeRepository")
			}
		})
	}

	return movieTypeRepository
}

func (obj *MovieTypeRepository) FindAllMovieTypesByStatusNot(status int) (*[]model.MovieType, error) {
	movieTypes := make([]model.MovieType, 0)

	err := db.Model(&model.MovieType{}).
		Find(&movieTypes, "status <> ?", status).Error

	return &movieTypes, err
}

func (obj *MovieTypeRepository) FindMovieTypeByIdAndStatusNot(id string, status int) (*model.MovieType, error) {
	uid := util.ParseStringToInt64(id)

	var movieType model.MovieType
	err := db.Model(&model.MovieType{}).
		Where("movie_type_id = ? AND status <> ?", uid, status).
		Find(&movieType).Error

	return &movieType, err
}

// SaveMovieType : Save Movie Type
func (obj *MovieTypeRepository) SaveMovieType(movieType model.MovieType) (*model.MovieType, error) {
	err := db.Model(&model.MovieType{}).
		Save(&movieType).Error

	return &movieType, err
}

func (obj *MovieTypeRepository) UpdateStatusByMovieTypeId(movieTypeId int64, status int) error {
	err := db.Model(&model.Movie{}).
		Where("movie_type_id = ?", movieTypeId).
		Update("status", status).Error

	return err
}

func (obj *MovieTypeRepository) FindMovieTypeBySlugAndMovieTypeIdNotAndStatusNotIn(slug string, id string, status []int) (*model.MovieType, error) {
	var movieType model.MovieType

	err := obj.db.
		Where("movie_type_id <> ?", id).
		Where("slug = ? AND status NOT IN ?", slug, status).
		Find(&movieType).Error

	return &movieType, err
}

func (obj *MovieTypeRepository) FindMovieTypeBySlugAndStatusNotIn(slug string, status []int) (*model.MovieType, error) {
	var movieType model.MovieType

	err := obj.db.
		Where("slug = ? AND status NOT IN ?", slug, status).
		Find(&movieType).Error

	return &movieType, err
}
