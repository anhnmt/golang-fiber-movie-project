package model

type MovieTag struct {
	MovieTagId uint `gorm:"primarykey"`
	MovieId    uint `gorm:"not null"`
	TagId      uint `gorm:"not null"`
	Status     int8 `gorm:"default:1"`
	BaseModel
}
