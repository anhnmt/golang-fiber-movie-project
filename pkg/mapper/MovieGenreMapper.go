package mapper

import (
	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
	"github.com/xdorro/golang-fiber-movie-project/pkg/validator"
)

func MovieGenres(movieId *int64, genreIds *[]int64) []model.MovieGenre {
	result := make([]model.MovieGenre, 0)

	for _, genreId := range *genreIds {
		mapper := MovieGenre(movieId, &genreId)
		result = append(result, *mapper)
	}

	return result
}

func MovieGenre(movieId *int64, genreId *int64) *model.MovieGenre {
	return &model.MovieGenre{
		MovieId: *movieId,
		GenreId: *genreId,
	}
}

func GetGenreIdsNotExistInNewGenreIds(newGenreIds []int64, genres []model.Genre) *[]int64 {
	result := make([]int64, 0)

	for _, genre := range genres {
		if !validator.ExistGenreIdInGenreIds(genre.GenreId, newGenreIds) {
			result = append(result, genre.GenreId)
		}
	}

	return &result
}

func GetNewGenreIdsNotExistInGenres(genreIds []int64, genres []model.Genre) *[]int64 {
	result := make([]int64, 0)

	for _, genreId := range genreIds {
		if !validator.ExistGenreIdInGenres(genreId, genres) {
			result = append(result, genreId)
		}
	}

	return &result
}
