package models

type MovieGenre struct {
	MovieGenreId uint `gorm:"primarykey"`
	MovieId      uint `gorm:"not null"`
	GenreId      uint `gorm:"not null"`
	Status       int8 `gorm:"default:1"`
	BaseModel
}
