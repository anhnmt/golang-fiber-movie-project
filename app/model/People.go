package model

type People struct {
	PeopleId uint   `gorm:"primarykey"`
	Name     string `gorm:"not null;unique"`
	Slug     string `gorm:"not null;unique"`
	Status   int8   `gorm:"default:1"`
	BaseModel
}
