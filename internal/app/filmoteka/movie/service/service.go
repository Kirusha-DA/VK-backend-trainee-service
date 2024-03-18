package service

import (
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/models"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/movie/dtos"
)

type Service interface {
	GetMoviesSortedByRatingDESC() ([]models.Movie, error)
	GetMoviesSortedByNameDESC() ([]models.Movie, error)
	GetMoviesSortedByReleaseDateDESC() ([]models.Movie, error)
	GetMoviesByActorName(actorName string) ([]dtos.MovieDTO, error)
	CreateMovie(movie *models.Movie) error
	UpdateMovieById(movie *models.Movie) error
	PartiallyUpdateActor(movie *models.Movie) error
	DeleteMovieById(movieId string) error
}
