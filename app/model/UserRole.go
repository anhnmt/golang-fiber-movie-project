package model

type UserRole struct {
	UserId uint `gorm:"not null" json:"user_id"`
	RoleId uint `gorm:"not null" json:"role_id"`
}
