package model

type Role struct {
	RoleId uint   `gorm:"primaryKey" json:"role_id"`
	Name   string `gorm:"not null;unique" json:"name"`
	Status int8   `gorm:"default:1" json:"status"`
}
