package models

import (
	"time"
)

type Actor struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      *string    `gorm:"type:varchar(255)" json:"name"`
	Sex       *string    `gorm:"check:sex in ('F','M', NULL);type:char(1)" json:"sex"`
	BirthDate *time.Time `json:"birth_date"`
	Movies    []Movie    `gorm:"many2many:actor_movies"`
}

//func (a *Actor) ToDTO() *ActorDTO {
//return &ActorDTO{
//ID:        a.ID,
//Name:      *a.Name,
//Sex:       *a.Sex,
//BirthDate: a.BirthDate,
//}
//}
