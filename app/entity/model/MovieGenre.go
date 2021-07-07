package model

type MovieGenre struct {
	MovieId uint `gorm:"index:,not null" json:"movie_id"`
	GenreId uint `gorm:"index:,not null" json:"genre_id"`
}
