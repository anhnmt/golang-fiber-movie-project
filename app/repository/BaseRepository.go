package repository

import (
	"sync"

	"gorm.io/gorm"

	"github.com/xdorro/golang-fiber-movie-project/platform/database"
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
	userRepository          *UserRepository
)
