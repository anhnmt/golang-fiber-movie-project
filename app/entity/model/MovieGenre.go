package model

type MovieGenre struct {
	MovieId int64 `gorm:"index:,not null" json:"movie_id"`
	GenreId int64 `gorm:"index:,not null" json:"genre_id"`
}
