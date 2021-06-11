package models

type Genre struct {
	GenreId uint   `gorm:"primarykey"`
	Name    string `gorm:"unique"`
	Slug    string `gorm:"unique"`
	Status  int8   `gorm:"default:1"`
	BaseModel
}
