package model

type Tag struct {
	TagId  uint   `gorm:"primaryKey" json:"tag_id"`
	Name   string `gorm:"not null;unique" json:"name"`
	Slug   string `gorm:"not null;unique" json:"slug"`
	Status int    `gorm:"index:,default:1" json:"status"`
	BaseModel
}
