package model

type MovieTag struct {
	MovieId uint `gorm:"index:,not null" json:"movie_id"`
	TagId   uint `gorm:"index:,not null" json:"tag_id"`
}
