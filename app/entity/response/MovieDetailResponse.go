package response

import "github.com/xdorro/golang-fiber-base-project/app/entity/dto"

type MovieDetailResponse struct {
	MovieId     uint             `json:"movie_id"`
	Name        string           `json:"name"`
	Slug        string           `json:"slug"`
	Description string           `json:"description"`
	Trailer     string           `json:"trailer"`
	ImdbId      string           `json:"imdb_id"`
	Rating      string           `json:"rating"`
	ReleaseDate string           `json:"release_date"`
	Runtime     string           `json:"runtime"`
	SeoTitle    string           `json:"seo_title"`
	SeoKeywords string           `json:"seo_keywords"`
	Status      int              `json:"status"`
	MovieType   dto.MovieTypeDTO `json:"movie_type"`
}
