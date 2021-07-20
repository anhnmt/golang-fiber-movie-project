package model

type EpisodeDetail struct {
	EpisodeDetailId uint   `gorm:"primaryKey" json:"episode_detail_id"`
	Name            string `gorm:"unique" json:"name"`
	EpisodeId       uint   `gorm:"index:,not null" json:"episode_id"`
	EpisodeTypeId   uint   `gorm:"index:,not null" json:"episode_type_id"`
	Status          int    `gorm:"index:,default:1" json:"status"`
	BaseModel
}
