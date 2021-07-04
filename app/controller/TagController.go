package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/dto"
	"github.com/xdorro/golang-fiber-base-project/app/model"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
)

// FindAllTags : Find all tags by Status = 1
func FindAllTags(c *fiber.Ctx) error {
	tags, err := repository.FindAllTagsByStatus(util.STATUS_ACTIVATED)

	if err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", tags)
}

// FindTagById : Find tag by Tag_Id and Status = 1
func FindTagById(c *fiber.Ctx) error {
	tagId := c.Params("id")
	tag, err := repository.FindTagByIdAndStatus(tagId, util.STATUS_ACTIVATED)

	if err != nil || tag.TagId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	return util.ResponseSuccess(c, "Thành công", tag)
}

// CreateNewTag : Create a new tag
func CreateNewTag(c *fiber.Ctx) error {
	tagRequest := new(dto.TagRequest)

	if err := c.BodyParser(tagRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	tag := model.Tag{
		Name:   tagRequest.Name,
		Slug:   tagRequest.Slug,
		Status: util.STATUS_ACTIVATED,
	}

	if _, err := repository.SaveTag(tag); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// UpdateTagById : Update tag by Tag_Id and Status = 1
func UpdateTagById(c *fiber.Ctx) error {
	tagId := c.Params("id")

	tag, err := repository.FindTagByIdAndStatus(tagId, util.STATUS_ACTIVATED)

	if err != nil || tag.TagId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	tagRequest := new(dto.TagRequest)
	if err = c.BodyParser(tagRequest); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	tag.Name = tagRequest.Name
	tag.Slug = tagRequest.Slug

	if _, err = repository.SaveTag(*tag); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}

// DeleteTagById : Delete tag by Tag_Id and Status = 1
func DeleteTagById(c *fiber.Ctx) error {
	tagId := c.Params("id")
	tag, err := repository.FindTagByIdAndStatus(tagId, util.STATUS_ACTIVATED)

	if err != nil || tag.TagId == 0 {
		return util.ResponseBadRequest(c, "ID không tồn tại", err)
	}

	tag.Status = util.STATUS_DELETED

	if _, err = repository.SaveTag(*tag); err != nil {
		return util.ResponseError(c, err.Error(), nil)
	}

	return util.ResponseSuccess(c, "Thành công", nil)
}
