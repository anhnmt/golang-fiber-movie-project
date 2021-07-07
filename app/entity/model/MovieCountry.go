package model

type MovieCountry struct {
	MovieId   uint `gorm:"index:,not null" json:"movie_id"`
	CountryId uint `gorm:"index:,not null" json:"country_id"`
}
