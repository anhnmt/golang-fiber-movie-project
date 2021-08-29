package model

type MovieTag struct {
	MovieId int64 `gorm:"index:,not null" json:"movie_id"`
	TagId   int64 `gorm:"index:,not null" json:"tag_id"`
}
