package dto

type MovieDetailDTO struct {
	MovieId       uint   `json:"movie_id"`
	OriginName    string `json:"origin_name"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	Description   string `json:"description"`
	Trailer       string `json:"trailer"`
	ImdbId        string `json:"imdb_id"`
	Rating        string `json:"rating"`
	ReleaseDate   string `json:"release_date"`
	Runtime       string `json:"runtime"`
	Poster        string `json:"poster"`
	SeoTitle      string `json:"seo_title"`
	SeoKeywords   string `json:"seo_keywords"`
	Status        int    `json:"status"`
	MovieTypeId   uint   `json:"movie_type_id"`
	MovieTypeName string `json:"movie_type_name"`
}
