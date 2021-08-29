package validator

import (
	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
)

func ExistGenreIdInGenres(genreId int64, genres []model.Genre) bool {
	for _, genre := range genres {
		if genre.GenreId == genreId {
			return true
		}
	}

	return false
}

func ExistGenreIdInGenreIds(genreId int64, genreIds []int64) bool {
	for _, genre := range genreIds {
		if genre == genreId {
			return true
		}
	}

	return false
}
