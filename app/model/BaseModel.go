package model

import "time"

type BaseModel struct {
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at,string,omitempty"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at,string,omitempty"`
}
