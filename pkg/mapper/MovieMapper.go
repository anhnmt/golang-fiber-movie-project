package mapper

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/dto"
	"github.com/xdorro/golang-fiber-base-project/app/entity/response"
)

func SearchMovies(movies *[]dto.SearchMovieDTO) []response.SearchMovieResponse {
	result := make([]response.SearchMovieResponse, 0)

	for _, movie := range *movies {
		mapper := SearchMovie(&movie)
		result = append(result, *mapper)
	}
	return result
}

func SearchMovie(movie *dto.SearchMovieDTO) *response.SearchMovieResponse {
	return &response.SearchMovieResponse{
		MovieId: movie.MovieId,
		Name:    movie.Name,
		Slug:    movie.Slug,
		Status:  movie.Status,
		MovieType: dto.MovieTypeDTO{
			MovieTypeId: movie.MovieTypeId,
			Name:        movie.MovieTypeName,
		},
	}
}

func MovieDetail(movie *dto.MovieDetailDTO) *response.MovieDetailResponse {
	return &response.MovieDetailResponse{
		MovieId:     movie.MovieId,
		Name:        movie.Name,
		Slug:        movie.Slug,
		Description: movie.Description,
		Trailer:     movie.Trailer,
		ImdbId:      movie.ImdbId,
		Rating:      movie.Rating,
		ReleaseDate: movie.ReleaseDate,
		Runtime:     movie.Runtime,
		SeoTitle:    movie.SeoTitle,
		SeoKeywords: movie.SeoKeywords,
		Status:      movie.Status,
		MovieType: dto.MovieTypeDTO{
			MovieTypeId: movie.MovieTypeId,
			Name:        movie.MovieTypeName,
		},
	}
}
