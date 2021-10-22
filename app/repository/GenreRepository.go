package repository

import (
	"log"
	"sync"

	"gorm.io/gorm"

	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
)

type GenreRepository struct {
	db *gorm.DB
}

func NewGenreRepository() *GenreRepository {
	if genreRepository == nil {
		once = &sync.Once{}

		once.Do(func() {
			if genreRepository == nil {
				genreRepository = &GenreRepository{
					db: db,
				}

				log.Println("Create new GenreRepository")
			}
		})
	}

	return genreRepository
}

func (obj *GenreRepository) CountAllGenresStatusNotIn(status []int) (int64, error) {
	var count int64

	err := db.
		Model(&model.Genre{}).
		Select("genres.genre_id").
		Where("genres.status NOT IN ?", status).
		Count(&count).Error

	return count, err
}

func (obj *GenreRepository) FindAllGenresByStatusNot(status int) (*[]model.Genre, error) {
	genres := make([]model.Genre, 0)

	err := db.Model(model.Genre{}).
		Find(&genres, "status <> ?", status).Error

	return &genres, err
}

func (obj *GenreRepository) FindAllGenresByStatusNotIn(status []int) (*[]model.Genre, error) {
	genres := make([]model.Genre, 0)

	err := db.Model(model.Genre{}).
		Find(&genres, "status NOT IN ?", status).Error

	return &genres, err
}

func (obj *GenreRepository) FindAllGenresByGenreIdsInAndStatusNotIn(genreIds []int64, status []int) (
	*[]model.Genre, error,
) {
	genres := make([]model.Genre, 0)

	err := db.Model(model.Genre{}).
		Find(&genres, "genre_id IN ? AND status NOT IN ?", genreIds, status).Error

	return &genres, err
}

func (obj *GenreRepository) FindGenreByIdAndStatusNot(id string, status int) (*model.Genre, error) {
	uid := util.ParseStringToInt64(id)

	var genre model.Genre
	err := db.Model(model.Genre{}).
		Where("genre_id = ? AND status <> ?", uid, status).
		Find(&genre).Error

	return &genre, err
}

func (obj *GenreRepository) SaveGenre(genre model.Genre) (*model.Genre, error) {
	err := db.Model(model.Genre{}).
		Save(&genre).Error

	return &genre, err
}

func (obj *GenreRepository) UpdateGenre(genreId string, genre model.Genre) (*model.Genre, error) {
	err := db.Model(model.Genre{}).
		Where("genre_id = ?", genreId).
		Save(&genre).Error

	return &genre, err
}

func (obj *GenreRepository) FindGenreBySlugAndGenreIdNotAndStatusNotIn(
	slug string, id string, status []int,
) (*model.Genre, error) {
	var genre model.Genre

	err := obj.db.
		Where("genre_id <> ?", id).
		Where("slug = ? AND status NOT IN ?", slug, status).
		Find(&genre).Error

	return &genre, err
}

func (obj *GenreRepository) FindGenreBySlugAndStatusNotIn(slug string, status []int) (*model.Genre, error) {
	var genre model.Genre

	err := obj.db.
		Where("slug = ? AND status NOT IN ?", slug, status).
		Find(&genre).Error

	return &genre, err
}
