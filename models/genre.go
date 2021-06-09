package models

import "github.com/google/uuid"

type Genre struct {
	GenreId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name    string    `gorm:"unique"`
	Slug    string    `gorm:"unique"`
	Status  int8      `gorm:"default:1"`
	BaseModel
}
