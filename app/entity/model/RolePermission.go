package model

type RolePermission struct {
	RoleId       uint `gorm:"index:,not null" json:"role_id"`
	PermissionId uint `gorm:"index:,not null" json:"permission_id"`
}
