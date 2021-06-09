package models

import "github.com/google/uuid"

type MovieCountry struct {
	MovieCountryId uuid.UUID `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	MovieId        uuid.UUID `gorm:"not null;unique"`
	CountryId      uuid.UUID `gorm:"not null;unique"`
	Status         int8      `gorm:"default:1"`
	BaseModel
}
