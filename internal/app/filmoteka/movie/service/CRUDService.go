package service

import (
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/models"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/movie/dtos"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/movie/repository"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
)

type service struct {
	repository repository.Repository
	logger     *logging.Logger
}

func NewService(repository repository.Repository, logger *logging.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s *service) GetMoviesSortedByRatingDESC() ([]models.Movie, error) {
	return s.repository.GetMoviesSortedByRatingDESC()
}

func (s *service) GetMoviesSortedByNameDESC() ([]models.Movie, error) {
	return s.repository.GetMoviesSortedByNameDESC()
}

func (s *service) GetMoviesSortedByReleaseDateDESC() ([]models.Movie, error) {
	return s.repository.GetMoviesSortedByReleaseDateDESC()
}

func (s *service) GetMoviesByActorName(actorName string) ([]dtos.MovieDTO, error) {
	var movies []dtos.MovieDTO
	allMovies, err := s.repository.GetMoviesByActorName(actorName)
	if err != nil {
		return nil, err
	}
	for _, movieFull := range allMovies {
		if len(movieFull.Actors) != 0 {
			movies = append(movies, *movieFull.ToDTO())
		}
	}
	return movies, nil
}

func (s *service) CreateMovie(movie *models.Movie) error {
	return s.repository.CreateMovie(movie)
}

func (s *service) UpdateMovieById(movie *models.Movie) error {
	return s.repository.UpdateMovieById(movie)
}

func (s *service) PartiallyUpdateActor(movie *models.Movie) error {
	return s.repository.PartiallyUpdateActor(movie)
}

func (s *service) DeleteMovieById(movieId string) error {
	return s.repository.DeleteMovieById(movieId)
}
