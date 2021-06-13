package model

type Permission struct {
	PermissionId uint   `gorm:"primarykey" json:"permission_id"`
	Name         string `gorm:"not null;unique" json:"name"`
	Status       int8   `gorm:"default:1" json:"status"`
}
