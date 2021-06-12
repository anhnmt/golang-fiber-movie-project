package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"github.com/xdorro/golang-fiber-base-project/platform/database"
)

func FindAllTags(c *fiber.Ctx) error {
	db := database.GetDB()
	var tags []model.Tag
	db.Find(&tags)
	return util.ResponseSuccess(c, tags, "Thành công")
}
