package model

type Movie struct {
	MovieId     int64  `gorm:"primaryKey" json:"movie_id"`
	OriginName  string `gorm:"not null" json:"origin_name"`
	Name        string `gorm:"not null" json:"name"`
	Slug        string `gorm:"index:,not null" json:"slug"`
	Description string `json:"description"`
	Trailer     string `json:"trailer"`
	ImdbId      string `json:"imdb_id"`
	Rating      string `json:"rating"`
	ReleaseDate string `json:"release_date"`
	Runtime     string `json:"runtime"`
	Poster      string `json:"poster"`
	SeoTitle    string `json:"seo_title"`
	SeoKeywords string `json:"seo_keywords"`
	Ribbon      string `json:"ribbon"`
	MovieTypeId int64  `gorm:"index:,not null" json:"movie_type_id"`
	Status      int    `gorm:"index:,default:1" json:"status"`
	BaseModel
}
