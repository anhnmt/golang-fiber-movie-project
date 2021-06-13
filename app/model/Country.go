package model

type Country struct {
	CountryId uint   `gorm:"primarykey" json:"country_id"`
	Name      string `gorm:"not null;unique" json:"name"`
	Slug      string `gorm:"not null;unique" json:"slug"`
	Status    int8   `gorm:"default:1" json:"status"`
	BaseModel
}
