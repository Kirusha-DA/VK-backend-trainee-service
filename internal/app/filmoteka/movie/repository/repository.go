package repository

import "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/models"

type Repository interface {
	GetMoviesSortedByRatingDESC() ([]models.Movie, error)
	GetMoviesSortedByNameDESC() ([]models.Movie, error)
	GetMoviesSortedByReleaseDateDESC() ([]models.Movie, error)
	GetMoviesByName(movieName string) ([]models.Movie, error)
	GetMoviesByActorName(actorName string) ([]models.Movie, error)
	CreateMovie(movie *models.Movie) error
	UpdateMovieById(movie *models.Movie) error
	PartiallyUpdateActor(movie *models.Movie) error
	DeleteMovieById(movieId string) error
}
