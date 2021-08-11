package response

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/dto"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
)

type MovieDetailResponse struct {
	MovieId        int64                 `json:"movie_id"`
	OriginName     string                `json:"origin_name"`
	Name           string                `json:"name"`
	Slug           string                `json:"slug"`
	Description    string                `json:"description,omitempty"`
	Trailer        string                `json:"trailer,omitempty"`
	ImdbId         string                `json:"imdb_id,omitempty"`
	Rating         string                `json:"rating,omitempty"`
	ReleaseDate    string                `json:"release_date"`
	Runtime        string                `json:"runtime,omitempty"`
	Poster         string                `json:"poster"`
	SeoTitle       string                `json:"seo_title,omitempty"`
	SeoKeywords    string                `json:"seo_keywords,omitempty"`
	Status         int                   `json:"status"`
	MovieType      dto.MovieTypeDTO      `json:"movie_type"`
	MovieRelated   []SearchMovieResponse `json:"movie_related,omitempty"`
	MovieGenres    []model.Genre         `json:"movie_genres,omitempty"`
	MovieCountries []model.Country       `json:"movie_countries,omitempty"`
}
