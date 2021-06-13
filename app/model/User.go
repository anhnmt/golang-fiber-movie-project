package model

type User struct {
	UserId   uint   `gorm:"primarykey" json:"user_id"`
	Name     string `gorm:"not null" json:"name"`
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Gender   int8   `gorm:"default:1" json:"gender"`
	Status   int8   `gorm:"default:1" json:"status"`
	BaseModel
}
