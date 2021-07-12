package validator

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
)

func ExistGenreIdInGenres(genreId uint, genres []model.Genre) bool {
	for _, genre := range genres {
		if genre.GenreId == genreId {
			return true
		}
	}

	return false
}

func ExistGenreIdInGenreIds(genreId uint, genreIds []uint) bool {
	for _, genre := range genreIds {
		if genre == genreId {
			return true
		}
	}

	return false
}
