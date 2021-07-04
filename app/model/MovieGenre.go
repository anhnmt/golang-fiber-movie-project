package model

type MovieGenre struct {
	MovieId uint `gorm:"not null" json:"movie_id"`
	GenreId uint `gorm:"not null" json:"genre_id"`
	Status  int  `gorm:"default:1" json:"status"`
}
