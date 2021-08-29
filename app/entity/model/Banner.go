package model

type Banner struct {
	BannerId int64  `gorm:"primary_key" json:"banner_id"`
	MovieId  int64  `gorm:"index:,not null" json:"movie_id"`
	Image    string `gorm:"not null" json:"image"`
	Status   int    `gorm:"index:,default:1" json:"status"`
	BaseModel
}
