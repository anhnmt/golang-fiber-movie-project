package model

type User struct {
	UserId   uint   `gorm:"primary_key" json:"user_id"`
	Name     string `gorm:"not null" json:"name"`
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Gender   int    `gorm:"default:1" json:"gender"`
	Status   int    `gorm:"index:,default:1" json:"status"`
	RoleId   uint   `gorm:"index:,not null" json:"role_id"`
	BaseModel
}
