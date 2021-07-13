package mapper

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/validator"
)

func MovieGenres(movieId *uint, genreIds *[]uint) []model.MovieGenre {
	result := make([]model.MovieGenre, 0)

	for _, genreId := range *genreIds {
		mapper := MovieGenre(movieId, &genreId)
		result = append(result, *mapper)
	}

	return result
}

func MovieGenre(movieId *uint, genreId *uint) *model.MovieGenre {
	return &model.MovieGenre{
		MovieId: *movieId,
		GenreId: *genreId,
	}
}

func GetGenreIdsNotExistInNewGenreIds(newGenreIds []uint, genres []model.Genre) *[]uint {
	result := make([]uint, 0)

	for _, genre := range genres {
		if !validator.ExistGenreIdInGenreIds(genre.GenreId, newGenreIds) {
			result = append(result, genre.GenreId)
		}
	}

	return &result
}

func GetNewGenreIdsNotExistInGenres(genreIds []uint, genres []model.Genre) *[]uint {
	result := make([]uint, 0)

	for _, genreId := range genreIds {
		if !validator.ExistGenreIdInGenres(genreId, genres) {
			result = append(result, genreId)
		}
	}

	return &result
}
