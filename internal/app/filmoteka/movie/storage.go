package movie

import "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/models"

type Repository interface {
	GetMoviesSortedByRatingDESC() []models.Movie
	GetMoviesSortedByNameDESC() []models.Movie
	GetMoviesSortedByReleaseDateDESC() []models.Movie
	GetMoviesByName(movieName string) []models.Movie
	GetMoviesByActorName(actorName string) []models.Movie
	CreateMovie(movie *models.Movie)
	UpdateMovieById(movie *models.Movie)
	PartiallyUpdateActor(movie *models.Movie)
	DeleteMovieById(movieId string)
}
