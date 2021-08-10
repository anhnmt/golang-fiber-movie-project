package model

type Country struct {
	CountryId uint   `gorm:"primaryKey" json:"country_id"`
	Name      string `gorm:"not null" json:"name"`
	Slug      string `gorm:"index:,not null" json:"slug"`
	Status    int    `gorm:"index:,default:1" json:"status"`
	BaseModel
}
