package models

type MovieCountry struct {
	MovieCountryId uint `gorm:"primarykey"`
	MovieId        uint `gorm:"not null"`
	CountryId      uint `gorm:"not null"`
	Status         int8 `gorm:"default:1"`
	BaseModel
}
