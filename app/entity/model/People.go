package model

type People struct {
	PeopleId uint   `gorm:"primaryKey" json:"people_id"`
	Name     string `gorm:"not null" json:"name"`
	Slug     string `gorm:"index:,not null" json:"slug"`
	Status   int    `gorm:"index:,default:1" json:"status"`
	BaseModel
}
