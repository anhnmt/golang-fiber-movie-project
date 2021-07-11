package request

type MovieRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	MovieTypeId uint   `json:"movie_type_id"`
	GenresId    []uint `json:"genres_id"`
	CountriesId []uint `json:"countries_id"`
	TagsId      []uint `json:"tags_id"`
	Status      int    `json:"status"`
}
