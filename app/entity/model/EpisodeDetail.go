package model

type EpisodeDetail struct {
	EpisodeDetailId int64  `gorm:"primaryKey" json:"episode_detail_id"`
	Name            string `gorm:"not null" json:"name"`
	Link            string `gorm:"not null" json:"link"`
	EpisodeId       int64  `gorm:"index:,not null" json:"episode_id"`
	EpisodeTypeId   int64  `gorm:"index:,not null" json:"episode_type_id"`
	Status          int    `gorm:"index:,default:1" json:"status"`
	BaseModel
}
