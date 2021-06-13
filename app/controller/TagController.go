package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/dto"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"github.com/xdorro/golang-fiber-base-project/platform/database"
	"gorm.io/gorm"
)

// FindAllTags : Find all tags by Status = 1
func FindAllTags(c *fiber.Ctx) error {
	db := database.GetDB()
	var tags []model.Tag
	db.Find(&tags, "status = ?", 1)
	return util.ResponseSuccess(c, "Thành công", tags)
}

// FindTagById : Find tag by Tag_Id and Status = 1
func FindTagById(c *fiber.Ctx) error {
	var err error
	db := database.GetDB()

	tagId := c.Params("id")
	tag := new(model.Tag)

	tag, err = findTagByIdAndStatus(tagId, tag, db)

	if err != nil || tag.TagId == 0 {
		return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
	}

	return util.ResponseSuccess(c, "Thành công", tag)
}

// CreateNewTag : Create a new tag
func CreateNewTag(c *fiber.Ctx) error {
	db := database.GetDB()
	tagRequest := new(dto.TagRequest)

	if err := c.BodyParser(tagRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	tag := model.Tag{
		Name: tagRequest.Name,
		Slug: tagRequest.Slug,
	}

	if err := db.Create(&tag).Error; err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// UpdateTagById : Update tag by Tag_Id and Status = 1
func UpdateTagById(c *fiber.Ctx) error {
	var err error
	db := database.GetDB()

	tagId := c.Params("id")
	tagRequest := new(dto.TagRequest)
	tag := new(model.Tag)

	tag, err = findTagByIdAndStatus(tagId, tag, db)

	if err != nil || tag.TagId == 0 {
		return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
	}

	if err := c.BodyParser(tagRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	tag.Name = tagRequest.Name
	tag.Slug = tagRequest.Slug

	if err := db.Save(&tag).Error; err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// DeleteTagById : Delete tag by Tag_Id and Status = 1
func DeleteTagById(c *fiber.Ctx) error {
	var err error
	db := database.GetDB()

	tag := new(model.Tag)

	tagId := c.Params("id")

	tag, err = findTagByIdAndStatus(tagId, tag, db)

	if err != nil || tag.TagId == 0 {
		return util.ResponseNotFound(c, "Đường dẫn không tồn tại")
	}

	if result := db.Model(&tag).Update("status", 0); result.Error != nil {
		return util.ResponseError(c, result.Error.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// findTagByIdAndStatus : Find tag by Tag_Id and Status = 1
func findTagByIdAndStatus(tagId string, tag *model.Tag, db *gorm.DB) (*model.Tag, error) {
	if err := db.First(&tag, "tag_id = ? and status = ?", tagId, 1).Error; err != nil {
		return nil, err
	}

	return tag, nil
}
