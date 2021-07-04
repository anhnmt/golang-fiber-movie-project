package model

type Movie struct {
	MovieId     uint   `gorm:"primaryKey" json:"movie_id"`
	Name        string `gorm:"not null;unique" json:"name"`
	Slug        string `gorm:"not null;unique" json:"slug"`
	Description string `json:"description"`
	MovieType   uint   `gorm:"not null" json:"movie_type"`
	Trailer     string `json:"trailer"`
	ImdbId      string `json:"imdb_id"`
	Rating      string `json:"rating"`
	ReleaseDate string `json:"release_date"`
	Runtime     string `json:"runtime"`
	SeoTitle    string `json:"seo_title"`
	SeoKeywords string `json:"seo_keywords"`
	Status      int    `gorm:"default:1" json:"status"`
	BaseModel
}
