package model

type Tag struct {
	TagId  uint   `gorm:"primarykey" json:"tag_id"`
	Name   string `gorm:"not null;unique" json:"name"`
	Slug   string `gorm:"not null;unique" json:"slug"`
	Status int8   `gorm:"default:1" json:"status"`
	BaseModel
}
