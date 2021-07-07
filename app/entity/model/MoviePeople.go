package model

type MoviePeople struct {
	MovieId  uint `gorm:"index:,not null" json:"movie_id"`
	PeopleId uint `gorm:"index:,not null" json:"people_id"`
}
