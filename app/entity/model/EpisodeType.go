package model

type EpisodeType struct {
	EpisodeTypeId int64  `gorm:"primaryKey" json:"episode_type_id"`
	Name          string `gorm:"not null" json:"name"`
	Status        int    `gorm:"index:,default:1" json:"status"`
	BaseModel
}
