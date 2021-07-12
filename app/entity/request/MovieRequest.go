package request

type MovieRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	MovieTypeId uint   `json:"movie_type_id"`
	GenreIds    []uint `json:"genre_ids"`
	CountryIds  []uint `json:"country_ids"`
	TagsId      []uint `json:"tags_id"`
	Status      int    `json:"status"`
}
