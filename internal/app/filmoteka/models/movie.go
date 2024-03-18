package models

import "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/movie/dtos"

type Movie struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"-"`
	Name        *string `gorm:"type:varchar(255)" json:"name"`
	Description *string `gorm:"type:text" json:"desctiption"`
	Rating      *uint   `gorm:"check:(rating <= 10 and rating >=1) or rating = NULL" json:"rating"`
	Actors      []Actor `gorm:"many2many:actor_movies" json:"actors,omitempty"`
}

func (m *Movie) ToDTO() *dtos.MovieDTO {
	return &dtos.MovieDTO{
		ID:          m.ID,
		Name:        *m.Description,
		Description: *m.Description,
		Rating:      *m.Rating,
	}
}
