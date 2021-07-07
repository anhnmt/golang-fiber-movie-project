package model

type UserPermission struct {
	UserId       uint `gorm:"not null" json:"user_id"`
	PermissionId uint `gorm:"not null" json:"permission_id"`
}
