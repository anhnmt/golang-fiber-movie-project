package model

type MovieCountry struct {
	MovieCountryId uint `gorm:"primarykey" json:"movie_country_id"`
	MovieId        uint `gorm:"not null" json:"movie_id"`
	CountryId      uint `gorm:"not null" json:"country_id"`
	Status         int8 `gorm:"default:1" json:"status"`
	BaseModel
}
