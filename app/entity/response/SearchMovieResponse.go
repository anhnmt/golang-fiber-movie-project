package response

import "github.com/xdorro/golang-fiber-base-project/app/entity/dto"

type SearchMovieResponse struct {
	Name      string           `json:"name"`
	Slug      string           `json:"slug"`
	Status    int              `json:"status"`
	MovieType dto.MovieTypeDTO `json:"movie_type"`
}
