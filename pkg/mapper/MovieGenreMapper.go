package mapper

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
)

func MovieGenres(movieId *uint, genresId *[]uint) []model.MovieGenre {
	result := make([]model.MovieGenre, 0)

	for _, genreId := range *genresId {
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
