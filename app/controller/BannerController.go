package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"log"
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

// FindAllBanners : Find all banners by Status Not
func (obj *BannerController) FindAllBanners(c *fiber.Ctx) error {
	banners, err := obj.bannerRepository.FindAllBannersByStatusNot(util.StatusDeleted)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", banners)
}

// FindBannerById : Find banner by Banner_Id and Status Not
func (obj *BannerController) FindBannerById(c *fiber.Ctx) error {
	bannerId := c.Params("id")
	banner, err := obj.bannerRepository.FindBannerByIdAndStatusNot(bannerId, util.StatusDeleted)

	if err != nil || banner.BannerId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	return util.ResponseSuccess("Thành công", banner)
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

	banner := model.Banner{
		Image:  bannerRequest.Image,
		Url:    bannerRequest.Url,
		Status: bannerRequest.Status,
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

	banner.Image = bannerRequest.Image
	banner.Url = bannerRequest.Url
	banner.Status = bannerRequest.Status

	if _, err = obj.bannerRepository.SaveBanner(*banner); err != nil {
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

	if _, err = obj.bannerRepository.SaveBanner(*banner); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}
