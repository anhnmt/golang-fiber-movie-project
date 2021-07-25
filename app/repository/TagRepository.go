package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository() *TagRepository {
	return tagRepository
}

// FindAllTagsByStatus : Find tag by TagId and Status
func (obj *TagRepository) FindAllTagsByStatus(status int) (*[]model.Tag, error) {
	tags := make([]model.Tag, 0)

	err := obj.db.Find(&tags, "status = ?", status).Error

	return &tags, err
}

func (obj *TagRepository) FindAllTagsByStatusNot(status int) (*[]model.Tag, error) {
	tags := make([]model.Tag, 0)

	err := obj.db.Find(&tags, "status <> ?", status).Error

	return &tags, err
}

// FindTagByIdAndStatus : Find tag by TagId and Status
func (obj *TagRepository) FindTagByIdAndStatus(id string, status int) (*model.Tag, error) {
	uid := util.ParseStringToUInt(id)

	var tag model.Tag
	err := obj.db.Where("tag_id = ? AND status = ?", uid, status).Find(&tag).Error

	return &tag, err
}

func (obj *TagRepository) FindTagByIdAndStatusNot(id string, status int) (*model.Tag, error) {
	uid := util.ParseStringToUInt(id)

	var tag model.Tag
	err := obj.db.Where("tag_id = ? AND status <> ?", uid, status).Find(&tag).Error

	return &tag, err
}

func (obj *TagRepository) SaveTag(tag model.Tag) (*model.Tag, error) {
	err := obj.db.Save(&tag).Error

	return &tag, err
}
