package model

type Country struct {
	CountryId uint   `gorm:"primaryKey" json:"country_id"`
	Name      string `gorm:"not null;unique" json:"name"`
	Slug      string `gorm:"not null;unique" json:"slug"`
	Status    int    `gorm:"default:1" json:"status"`
	BaseModel
}
