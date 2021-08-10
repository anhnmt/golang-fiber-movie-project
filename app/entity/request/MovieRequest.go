package request

type MovieRequest struct {
	OriginName  string `json:"origin_name"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Trailer     string `json:"trailer"`
	ImdbId      string `json:"imdb_id"`
	Rating      string `json:"rating"`
	ReleaseDate string `json:"release_date"`
	Runtime     string `json:"runtime"`
	Poster      string `json:"poster"`
	SeoTitle    string `json:"seo_title"`
	SeoKeywords string `json:"seo_keywords"`
	MovieTypeId uint   `json:"movie_type_id"`
	GenreIds    []uint `json:"genre_ids"`
	CountryIds  []uint `json:"country_ids"`
	Status      int    `json:"status"`
}
