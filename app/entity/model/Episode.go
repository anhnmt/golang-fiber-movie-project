package model

type Episode struct {
	EpisodeId uint   `gorm:"primaryKey" json:"episode_id"`
	Name      string `gorm:"unique" json:"name"`
	MovieId   uint   `gorm:"index:,not null" json:"movie_id"`
	Status    int    `gorm:"index:,default:1" json:"status"`
	BaseModel
}
