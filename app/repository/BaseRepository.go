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
	once *sync.Once

	tagRepository *TagRepository

	episodeRepository *EpisodeRepository

	episodeDetailRepository *EpisodeDetailRepository

	episodeTypeRepository *EpisodeTypeRepository

	movieRepository *MovieRepository

	movieTypeRepository *MovieTypeRepository
)
