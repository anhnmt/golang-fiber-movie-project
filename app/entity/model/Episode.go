package model

type Episode struct {
	EpisodeId int64  `gorm:"primaryKey" json:"episode_id"`
	Name      string `gorm:"not null" json:"name"`
	MovieId   int64  `gorm:"index:,not null" json:"movie_id"`
	Status    int    `gorm:"index:,default:1" json:"status"`
	BaseModel
}
