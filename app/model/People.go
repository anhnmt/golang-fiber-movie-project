package model

type People struct {
	PeopleId uint   `gorm:"primaryKey" json:"people_id"`
	Name     string `gorm:"not null;unique" json:"name"`
	Slug     string `gorm:"not null;unique" json:"slug"`
	Status   int    `gorm:"default:1" json:"status"`
	BaseModel
}
