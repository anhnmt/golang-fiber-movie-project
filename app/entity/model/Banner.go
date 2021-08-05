package model

type Banner struct {
	BannerId uint   `gorm:"primary_key" json:"banner_id"`
	Image    string `gorm:"not null" json:"image"`
	Url      uint   `gorm:"index:,not null" json:"url"`
	Status   int    `gorm:"index:,default:1" json:"status"`
	BaseModel
}
