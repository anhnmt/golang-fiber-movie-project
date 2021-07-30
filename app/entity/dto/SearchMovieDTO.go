package dto

type SearchMovieDTO struct {
	MovieId       uint   `json:"movie_id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	MovieTypeId   uint   `json:"movie_type_id,omitempty"`
	MovieTypeName string `json:"movie_type_name,omitempty"`
	Status        int    `json:"status"`
}
