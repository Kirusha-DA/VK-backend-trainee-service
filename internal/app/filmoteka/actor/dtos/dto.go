package dtos

import "time"

type ActorDTO struct {
	Name      *string    `json:"name"`
	Sex       *string    `json:"sex"`
	BirthDate *time.Time `json:"birth_date"`
}
