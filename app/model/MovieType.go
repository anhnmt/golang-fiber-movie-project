package model

type MovieType struct {
	MovieTypeId uint   `gorm:"primaryKey" json:"movie_type_id"`
	Name        string `gorm:"not null;unique" json:"name"`
	Slug        string `gorm:"not null;unique" json:"slug"`
	Status      int8   `gorm:"default:1" json:"status"`
	BaseModel
}
