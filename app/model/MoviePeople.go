package model

type MoviePeople struct {
	MovieId  uint `gorm:"not null" json:"movie_id"`
	PeopleId uint `gorm:"not null" json:"people_id"`
	Status   int8 `gorm:"default:1" json:"status"`
}
