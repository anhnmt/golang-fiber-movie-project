package dto

type SearchMovieDTO struct {
	MovieId       uint   `json:"movie_id"`
	OriginName    string `json:"origin_name"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	Description   string `json:"description,omitempty"`
	Trailer       string `json:"trailer,omitempty"`
	ImdbId        string `json:"imdb_id,omitempty"`
	Rating        string `json:"rating,omitempty"`
	ReleaseDate   string `json:"release_date"`
	Runtime       string `json:"runtime,omitempty"`
	Poster        string `json:"poster,omitempty"`
	SeoTitle      string `json:"seo_title,omitempty"`
	SeoKeywords   string `json:"seo_keywords,omitempty"`
	MovieTypeId   uint   `json:"movie_type_id,omitempty"`
	MovieTypeName string `json:"movie_type_name,omitempty"`
	Status        int    `json:"status"`
}
