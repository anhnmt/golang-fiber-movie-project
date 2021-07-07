package model

type RolePermission struct {
	RoleId       uint `gorm:"not null" json:"role_id"`
	PermissionId uint `gorm:"not null" json:"permission_id"`
}
