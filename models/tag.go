package models

import "github.com/google/uuid"

type Tag struct {
	TagId  uuid.UUID `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	Name   string    `gorm:"not null;unique"`
	Slug   string    `gorm:"not null;unique"`
	Status int8      `gorm:"default:1"`
	BaseModel
}
