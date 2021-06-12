package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"github.com/xdorro/golang-fiber-base-project/platform/database"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db = database.GetDB()
}

// FindAllTags : Find all tags by Status = 1
func FindAllTags(c *fiber.Ctx) error {
	var tags []model.Tag
	db.Find(&tags, "status = ?", 1)
	return util.ResponseSuccessData(c, tags, "Thành công")
}

// FindTagById : Find tag by Tag_Id and Status = 1
func FindTagById(c *fiber.Ctx) error {
	tagId := c.Params("id")
	tag := new(model.Tag)

	tag, err = findTagByIdAndStatus(tagId, tag, db)

	if err != nil || tag.TagId == 0 {
		return util.ResponseNotFound(c, "Không tìm thấy ID này")
	}

	return util.ResponseSuccessData(c, tag, "Thành công")
}

// CreateNewTag : Create a new tag
func CreateNewTag(c *fiber.Ctx) error {
	tag := new(model.Tag)

	if err := c.BodyParser(tag); err != nil {
		return util.ResponseError(c, err.Error())
	}

	result := db.Create(&tag)

	return util.ResponseSuccessData(c, result, "Thành công")
}

// UpdateTagById : Update tag by Tag_Id and Status = 1
func UpdateTagById(c *fiber.Ctx) error {
	tagId := c.Params("id")
	tag := new(model.Tag)

	tag, err = findTagByIdAndStatus(tagId, tag, db)

	if err != nil || tag.TagId == 0 {
		return util.ResponseNotFound(c, "Không tìm thấy ID này")
	}

	if err := c.BodyParser(tag); err != nil {
		return util.ResponseError(c, err.Error())
	}

	db.Save(&tag)

	return util.ResponseSuccess(c, "Thành công")
}

// DeleteTagById : Delete tag by Tag_Id and Status = 1
func DeleteTagById(c *fiber.Ctx) error {
	tag := new(model.Tag)

	tagId := c.Params("id")
	db := database.GetDB()

	tag, err = findTagByIdAndStatus(tagId, tag, db)

	if err != nil || tag.TagId == 0 {
		return util.ResponseNotFound(c, "Không tìm thấy ID này")
	}

	db.Model(&tag).Update("status", 0)

	return util.ResponseSuccess(c, "Thành công")
}

// findTagByIdAndStatus : Find tag by Tag_Id and Status = 1
func findTagByIdAndStatus(tagId string, tag *model.Tag, db *gorm.DB) (*model.Tag, error) {
	if result := db.First(&tag, "tag_id = ? and status = ?", tagId, 1); result.Error != nil {
		return nil, result.Error
	}

	return tag, nil
}
