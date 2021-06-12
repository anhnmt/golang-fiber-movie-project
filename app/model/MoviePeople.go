package model

type MoviePeople struct {
	MoviePeopleId uint `gorm:"primarykey" json:"movie_people_id"`
	MovieId       uint `gorm:"not null" json:"movie_id"`
	PeopleId      uint `gorm:"not null" json:"people_id"`
	Status        int8 `gorm:"default:1" json:"status"`
	BaseModel
}
