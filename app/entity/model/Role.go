package model

type Role struct {
	RoleId int64  `gorm:"primaryKey" json:"role_id"`
	Name   string `gorm:"not null" json:"name"`
	Slug   string `gorm:"index:,not null" json:"slug"`
	Status int    `gorm:"index:,default:1" json:"status"`
}
