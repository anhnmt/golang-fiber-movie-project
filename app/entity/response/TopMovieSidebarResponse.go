package response

import "github.com/xdorro/golang-fiber-base-project/app/entity/dto"

type TopMovieSidebarResponse struct {
	Movies []dto.SearchMovieDTO `json:"movies"`
	Series []dto.SearchMovieDTO `json:"series"`
}
