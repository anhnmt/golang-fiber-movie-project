package model

type MovieTag struct {
	MovieTagId uint `gorm:"primarykey" json:"movie_tag_id"`
	MovieId    uint `gorm:"not null" json:"movie_id"`
	TagId      uint `gorm:"not null" json:"tag_id"`
	Status     int8 `gorm:"default:1" json:"status"`
	BaseModel
}
