package repository

import (
	"github.com/xdorro/golang-fiber-movie-project/app/entity/dto"
	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
	"gorm.io/gorm"
	"log"
	"sync"
)

type BannerRepository struct {
	db *gorm.DB
}

func NewBannerRepository() *BannerRepository {
	if bannerRepository == nil {
		once = &sync.Once{}

		once.Do(func() {
			if bannerRepository == nil {
				bannerRepository = &BannerRepository{
					db: db,
				}
				log.Println("Create new BannerRepository")
			}
		})
	}

	return bannerRepository
}

// FindAllBannersByStatus : Find banner by BannerId and Status
func (obj *BannerRepository) FindAllBannersByStatus(status int) (*[]model.Banner, error) {
	banners := make([]model.Banner, 0)

	err := db.Model(model.Banner{}).
		Find(&banners, "status = ?", status).Error

	return &banners, err
}

func (obj *BannerRepository) CountAllBannersStatusNotIn(status []int) (int64, error) {
	var count int64

	err := db.
		Model(&model.Banner{}).
		Select("banners.banner_id").
		Where("banners.status NOT IN ?", status).
		Count(&count).Error

	return count, err
}

func (obj *BannerRepository) FindAllBannersByStatusNot(status int) (*[]model.Banner, error) {
	banners := make([]model.Banner, 0)

	err := db.Model(model.Banner{}).
		Find(&banners, "status <> ?", status).Error

	return &banners, err
}

func (obj *BannerRepository) FindAllBannersByStatusNotIn(status []int) (*[]model.Banner, error) {
	banners := make([]model.Banner, 0)

	err := db.Model(model.Banner{}).
		Find(&banners, "status NOT IN ?", status).Error

	return &banners, err
}

func (obj *BannerRepository) FindAllBannersByStatusNotInAndJoinMovie(status []int) (*[]dto.SearchBannerDTO, error) {
	banners := make([]dto.SearchBannerDTO, 0)

	err := db.Model(model.Banner{}).
		Select("banners.*, movies.movie_id, movies.origin_name, movies.name, movies.slug, movies.release_date").
		Joins("JOIN movies on movies.movie_id = banners.movie_id").
		Where("movies.status NOT IN ?", status).
		Where("banners.status NOT IN ?", status).
		Order("banners.updated_at DESC").
		Find(&banners).Error

	return &banners, err
}

// FindBannerByIdAndStatus : Find banner by BannerId and Status
func (obj *BannerRepository) FindBannerByIdAndStatus(id string, status int) (*model.Banner, error) {
	var banner model.Banner

	err := obj.db.Model(model.Banner{}).
		Where("banner_id = ? AND status = ?", id, status).
		Find(&banner).Error

	return &banner, err
}

func (obj *BannerRepository) FindBannerByIdAndStatusNot(id string, status int) (*model.Banner, error) {
	var banner model.Banner

	err := obj.db.Model(model.Banner{}).
		Where("banner_id = ? AND status <> ?", id, status).
		Find(&banner).Error

	return &banner, err
}

func (obj *BannerRepository) FindBannerByIdAndStatusNotAndJoinMovie(id string, status int) (*dto.SearchBannerDTO, error) {
	var banner dto.SearchBannerDTO

	err := db.Model(model.Banner{}).
		Select("banners.*, movies.movie_id, movies.origin_name, movies.name, movies.slug, movies.release_date").
		Joins("JOIN movies on movies.movie_id = banners.movie_id").
		Where("movies.status <> ?", status).
		Where("banners.status <> ?", status).
		Where("banners.banner_id = ?", id).
		Find(&banner).Error

	return &banner, err
}

func (obj *BannerRepository) SaveBanner(banner model.Banner) (*model.Banner, error) {
	err := obj.db.Model(model.Banner{}).
		Save(&banner).Error

	return &banner, err
}

func (obj *BannerRepository) UpdateBanner(bannerId string, banner model.Banner) (*model.Banner, error) {
	err := obj.db.Model(model.Banner{}).
		Where("banner_id = ?", bannerId).
		Save(&banner).Error

	return &banner, err
}
