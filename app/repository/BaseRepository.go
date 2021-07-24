package repository

import (
	"github.com/xdorro/golang-fiber-base-project/platform/database"
	"gorm.io/gorm"
	"sync"
)

var db *gorm.DB

func init() {
	db = database.GetDB()
}

var (
	once sync.Once

	tagInstance *TagRepository
)

func NewTagRepository() *TagRepository {
	once.Do(func() {
		tagInstance = &TagRepository{
			db: db,
		}
	})

	return tagInstance
}
