package model

type MovieType struct {
	MovieTypeId int64  `gorm:"primaryKey" json:"movie_type_id"`
	Name        string `gorm:"not null" json:"name"`
	Slug        string `gorm:"not null" json:"slug"`
	Status      int    `gorm:"index:,default:1" json:"status"`
	BaseModel
}
