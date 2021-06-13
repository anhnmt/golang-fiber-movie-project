package model

type Genre struct {
	GenreId uint   `gorm:"primarykey" json:"genre_id"`
	Name    string `gorm:"unique" json:"name"`
	Slug    string `gorm:"unique" json:"slug"`
	Status  int8   `gorm:"default:1" json:"status"`
	BaseModel
}
