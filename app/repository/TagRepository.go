package repository

import (
	"log"
	"sync"

	"gorm.io/gorm"

	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository() *TagRepository {
	if tagRepository == nil {
		once = &sync.Once{}

		once.Do(func() {
			if tagRepository == nil {
				tagRepository = &TagRepository{
					db: db,
				}
				log.Println("Create new TagRepository")
			}
		})
	}

	return tagRepository
}

// FindAllTagsByStatus : Find tag by TagId and Status
func (obj *TagRepository) FindAllTagsByStatus(status int) (*[]model.Tag, error) {
	tags := make([]model.Tag, 0)

	err := db.Model(model.Tag{}).
		Find(&tags, "status = ?", status).Error

	return &tags, err
}

func (obj *TagRepository) FindAllTagsByStatusNot(status int) (*[]model.Tag, error) {
	tags := make([]model.Tag, 0)

	err := db.Model(model.Tag{}).
		Find(&tags, "status <> ?", status).Error

	return &tags, err
}

// FindTagByIdAndStatus : Find tag by TagId and Status
func (obj *TagRepository) FindTagByIdAndStatus(id string, status int) (*model.Tag, error) {
	var tag model.Tag

	err := obj.db.Model(model.Tag{}).
		Where("tag_id = ? AND status = ?", id, status).
		Find(&tag).Error

	return &tag, err
}

func (obj *TagRepository) FindTagByIdAndStatusNot(id string, status int) (*model.Tag, error) {
	var tag model.Tag

	err := obj.db.Model(model.Tag{}).
		Where("tag_id = ? AND status <> ?", id, status).
		Find(&tag).Error

	return &tag, err
}

func (obj *TagRepository) SaveTag(tag model.Tag) (*model.Tag, error) {
	err := obj.db.Model(model.Tag{}).
		Save(&tag).Error

	return &tag, err
}
