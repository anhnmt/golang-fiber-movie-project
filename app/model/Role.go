package model

type Role struct {
	RoleId uint   `gorm:"primarykey" json:"role_id"`
	Name   string `gorm:"not null;unique" json:"name"`
	Status int8   `gorm:"default:1" json:"status"`
}
