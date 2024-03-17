package movie

import "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/models"

type MovieDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desctiption"`
	Rating      uint   `json:"rating"`
}

func (movieDTO *MovieDTO) ToMovieModel() *models.Movie {
	return &models.Movie{
		Name:        &movieDTO.Name,
		Description: &movieDTO.Description,
		Rating:      &movieDTO.Rating,
	}
}
