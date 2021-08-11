package model

type MoviePeople struct {
	MovieId  int64 `gorm:"index:,not null" json:"movie_id"`
	PeopleId int64 `gorm:"index:,not null" json:"people_id"`
}
