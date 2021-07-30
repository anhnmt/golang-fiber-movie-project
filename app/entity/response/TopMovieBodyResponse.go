package response

import "github.com/xdorro/golang-fiber-base-project/app/entity/dto"

type TopMovieBodyResponse struct {
	Cinemas  []dto.SearchMovieDTO `json:"cinemas,omitempty"`
	Movies   []dto.SearchMovieDTO `json:"movies,omitempty"`
	Series   []dto.SearchMovieDTO `json:"series,omitempty"`
	Cartoons []dto.SearchMovieDTO `json:"cartoons,omitempty"`
}
