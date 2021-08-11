package model

type MovieCountry struct {
	MovieId   int64 `gorm:"index:,not null" json:"movie_id"`
	CountryId int64 `gorm:"index:,not null" json:"country_id"`
}
