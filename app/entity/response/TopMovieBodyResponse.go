package response

import "github.com/xdorro/golang-fiber-base-project/app/entity/dto"

type TopMovieBodyResponse struct {
	Cinemas  []dto.SearchMovieDTO `json:"cinemas"`
	Movies   []dto.SearchMovieDTO `json:"movies"`
	Series   []dto.SearchMovieDTO `json:"series"`
	Cartoons []dto.SearchMovieDTO `json:"cartoons"`
}
