package dto

type SearchMovieDTO struct {
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	MovieTypeId   uint   `json:"movie_type_id"`
	MovieTypeName string `json:"movie_type_name"`
	Status        int    `json:"status"`
}
