package model

type UserPermission struct {
	UserId       uint `gorm:"index:,not null" json:"user_id"`
	PermissionId uint `gorm:"index:,not null" json:"permission_id"`
}
