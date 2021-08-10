package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/mapper"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"github.com/xdorro/golang-fiber-base-project/pkg/validator"
	"log"
	"strconv"
	"sync"
)

type BannerController struct {
	bannerRepository *repository.BannerRepository
}

func NewBannerController() *BannerController {
	if bannerController == nil {
		once = &sync.Once{}

		once.Do(func() {
			if bannerController == nil {
				bannerController = &BannerController{
					bannerRepository: repository.NewBannerRepository(),
				}
				log.Println("Create new BannerController")
			}
		})
	}

	return bannerController
}

func (obj *BannerController) ClientFindAllBanners(c *fiber.Ctx) error {
	banners, err := obj.bannerRepository.FindAllBannersByStatusNotInAndJoinMovie([]int{util.StatusDraft, util.StatusDeleted})

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := mapper.SearchBannerMapper(banners)

	return util.ResponseSuccess("Thành công", result)
}

// FindAllBanners : Find all banners by Status Not
func (obj *BannerController) FindAllBanners(c *fiber.Ctx) error {
	banners, err := obj.bannerRepository.FindAllBannersByStatusNotInAndJoinMovie([]int{util.StatusDeleted})

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := mapper.SearchBannerMapper(banners)

	return util.ResponseSuccess("Thành công", result)
}

// FindBannerById : Find banner by Banner_Id and Status Not
func (obj *BannerController) FindBannerById(c *fiber.Ctx) error {
	bannerId := c.Params("id")
	banner, err := obj.bannerRepository.FindBannerByIdAndStatusNotAndJoinMovie(bannerId, util.StatusDeleted)

	if err != nil || banner.BannerId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	result := mapper.SearchBanner(banner)

	return util.ResponseSuccess("Thành công", result)
}

// CreateNewBanner : Create a new banner
func (obj *BannerController) CreateNewBanner(c *fiber.Ctx) error {
	bannerForm := c.FormValue("banner")

	bannerRequest := new(request.BannerRequest)

	if err := util.JSONUnmarshal([]byte(bannerForm), bannerRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	// Get first file from form field "image":
	file, err := c.FormFile("image")

	if file != nil {
		// 👷 Save file inside uploads folder under current working directory:
		image := util.StorageBanner(file.Filename)

		if err = c.SaveFile(file, util.Storage(image)); err != nil {
			return err
		}

		bannerRequest.Image = image
	}

	if bannerRequest.MovieId != 0 {
		if _, err = validator.ValidateMovieId(strconv.Itoa(int(bannerRequest.MovieId))); err != nil {
			return err
		}
	}

	banner := model.Banner{
		MovieId: bannerRequest.MovieId,
		Image:   bannerRequest.Image,
		Status:  bannerRequest.Status,
	}

	if _, err = obj.bannerRepository.SaveBanner(banner); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// UpdateBannerById : Update banner by Banner_Id and Status = 1
func (obj *BannerController) UpdateBannerById(c *fiber.Ctx) error {
	bannerId := c.Params("id")
	banner, err := obj.bannerRepository.FindBannerByIdAndStatusNot(bannerId, util.StatusDeleted)

	if err != nil || banner.BannerId == 0 {
		return util.ResponseBadRequest("BannerId không tồn tại", err)
	}

	bannerForm := c.FormValue("banner")

	bannerRequest := new(request.BannerRequest)

	if err = util.JSONUnmarshal([]byte(bannerForm), bannerRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	// Get first file from form field "image":
	file, err := c.FormFile("image")

	if file != nil {
		// 👷 Save file inside uploads folder under current working directory:
		image := util.StorageBanner(file.Filename)

		if err = c.SaveFile(file, util.Storage(image)); err != nil {
			return err
		}

		bannerRequest.Image = image
	}

	if bannerRequest.MovieId != 0 {
		if _, err = validator.ValidateMovieId(strconv.Itoa(int(bannerRequest.MovieId))); err != nil {
			return err
		}
	}

	banner.MovieId = bannerRequest.MovieId
	banner.Image = bannerRequest.Image
	banner.Status = bannerRequest.Status

	if _, err = obj.bannerRepository.UpdateBanner(bannerId, *banner); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// DeleteBannerById : Delete banner by Banner_Id and Status = 1
func (obj *BannerController) DeleteBannerById(c *fiber.Ctx) error {
	bannerId := c.Params("id")
	banner, err := obj.bannerRepository.FindBannerByIdAndStatusNot(bannerId, util.StatusDeleted)

	if err != nil || banner.BannerId == 0 {
		return util.ResponseBadRequest("BannerId không tồn tại", err)
	}

	banner.Status = util.StatusDeleted

	if _, err = obj.bannerRepository.UpdateBanner(bannerId, *banner); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}
