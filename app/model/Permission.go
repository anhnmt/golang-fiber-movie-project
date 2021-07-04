package model

type Permission struct {
	PermissionId uint   `gorm:"primaryKey" json:"permission_id"`
	Name         string `gorm:"not null;unique" json:"name"`
	Status       int    `gorm:"default:1" json:"status"`
}
