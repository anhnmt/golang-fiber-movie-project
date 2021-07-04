package model

type MovieCountry struct {
	MovieId   uint `gorm:"not null" json:"movie_id"`
	CountryId uint `gorm:"not null" json:"country_id"`
	Status    int  `gorm:"default:1" json:"status"`
}
