package repository

import (
	"github.com/xdorro/golang-fiber-movie-project/platform/database"
	"gorm.io/gorm"
	"sync"
)

var db *gorm.DB

func init() {
	db = database.GetDB()
}

var (
	once                    *sync.Once
	tagRepository           *TagRepository
	episodeRepository       *EpisodeRepository
	episodeDetailRepository *EpisodeDetailRepository
	episodeTypeRepository   *EpisodeTypeRepository
	movieRepository         *MovieRepository
	movieTypeRepository     *MovieTypeRepository
	countryRepository       *CountryRepository
	genreRepository         *GenreRepository
	movieGenreRepository    *MovieGenreRepository
	movieCountryRepository  *MovieCountryRepository
	bannerRepository        *BannerRepository
)
