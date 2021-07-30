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

type TagController struct {
	tagRepository *repository.TagRepository
}

func NewTagController() *TagController {
	if tagController == nil {
		once := &sync.Once{}

		once.Do(func() {
			if tagController == nil {
				tagController = &TagController{
					tagRepository: repository.NewTagRepository(),
				}
				log.Println("Create new TagController")
			}
		})
	}

	return tagController
}

// FindAllTags : Find all tags by Status Not
func (obj *TagController) FindAllTags(c *fiber.Ctx) error {
	tags, err := obj.tagRepository.FindAllTagsByStatusNot(util.StatusDeleted)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", tags)
}

// FindTagById : Find tag by Tag_Id and Status Not
func (obj *TagController) FindTagById(c *fiber.Ctx) error {
	tagId := c.Params("id")
	tag, err := obj.tagRepository.FindTagByIdAndStatusNot(tagId, util.StatusDeleted)

	if err != nil || tag.TagId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	return util.ResponseSuccess("Thành công", tag)
}

// CreateNewTag : Create a new tag
func (obj *TagController) CreateNewTag(c *fiber.Ctx) error {
	tagRequest := new(request.TagRequest)

	if err := c.BodyParser(tagRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	tag := model.Tag{
		Name:   tagRequest.Name,
		Slug:   tagRequest.Slug,
		Status: tagRequest.Status,
	}

	if _, err := obj.tagRepository.SaveTag(tag); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// UpdateTagById : Update tag by Tag_Id and Status = 1
func (obj *TagController) UpdateTagById(c *fiber.Ctx) error {
	tagId := c.Params("id")
	tag, err := obj.tagRepository.FindTagByIdAndStatusNot(tagId, util.StatusDeleted)

	if err != nil || tag.TagId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	tagRequest := new(request.TagRequest)
	if err = c.BodyParser(tagRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	tag.Name = tagRequest.Name
	tag.Slug = tagRequest.Slug
	tag.Status = tagRequest.Status

	if _, err = obj.tagRepository.SaveTag(*tag); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

// DeleteTagById : Delete tag by Tag_Id and Status = 1
func (obj *TagController) DeleteTagById(c *fiber.Ctx) error {
	tagId := c.Params("id")
	tag, err := obj.tagRepository.FindTagByIdAndStatusNot(tagId, util.StatusDeleted)

	if err != nil || tag.TagId == 0 {
		return util.ResponseBadRequest("ID không tồn tại", err)
	}

	tag.Status = util.StatusDeleted

	if _, err = obj.tagRepository.SaveTag(*tag); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}
