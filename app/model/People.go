package model

type People struct {
	PeopleId uint   `gorm:"primarykey" json:"people_id"`
	Name     string `gorm:"not null;unique" json:"name"`
	Slug     string `gorm:"not null;unique" json:"slug"`
	Status   int8   `gorm:"default:1" json:"status"`
	BaseModel
}
