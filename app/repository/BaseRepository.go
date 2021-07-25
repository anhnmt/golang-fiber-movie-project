package repository

import (
	"github.com/xdorro/golang-fiber-base-project/platform/database"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = database.GetDB()
}

var (
	tagRepository = &TagRepository{
		db: db,
	}

	episodeRepository = &EpisodeRepository{
		db: db,
	}

	episodeDetailRepository = &EpisodeDetailRepository{
		db: db,
	}
)
