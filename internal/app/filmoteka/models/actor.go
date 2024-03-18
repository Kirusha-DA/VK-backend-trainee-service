package models

import (
	"time"
)

type Actor struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"-"`
	Name      *string    `gorm:"type:varchar(255)" json:"name"`
	Sex       *string    `gorm:"check:sex in ('F','M', NULL);type:char(1)" json:"sex"`
	BirthDate *time.Time `json:"birth_date"`
	Movies    []Movie    `gorm:"many2many:actor_movies" json:"movies,omitempty" swaggerignore:"true"`
}
