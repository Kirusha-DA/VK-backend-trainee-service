package movie

import (
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/models"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
)

type Service struct {
	repository Repository
	logger     *logging.Logger
}

func NewService(repository Repository, logger *logging.Logger) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}

func (s *Service) GetMoviesSortedByRatingDESC() (movies []models.Movie) {
	return s.repository.GetMoviesSortedByRatingDESC()
}

func (s *Service) GetMoviesSortedByNameDESC() (movies []models.Movie) {
	return s.repository.GetMoviesSortedByNameDESC()
}

func (s *Service) GetMoviesSortedByReleaseDateDESC() (movies []models.Movie) {
	return s.repository.GetMoviesSortedByReleaseDateDESC()
}

func (s *Service) GetMoviesByActorName(actorName string) (movies []MovieDTO) {
	allMovies := s.repository.GetMoviesByActorName(actorName)
	for _, movieFull := range allMovies {
		if len(movieFull.Actors) != 0 {
			movies = append(movies, MovieDTO{
				ID:          movieFull.ID,
				Name:        *movieFull.Name,
				Description: *movieFull.Description,
				Rating:      *movieFull.Rating,
			})
		}
	}
	return movies
}

func (s *Service) CreateMovie(movie *models.Movie) {
	s.repository.CreateMovie(movie)
}

func (s *Service) UpdateMovieById(movie *models.Movie) {
	s.repository.UpdateMovieById(movie)
}

func (s *Service) PartiallyUpdateActor(movie *models.Movie) {
	s.repository.PartiallyUpdateActor(movie)
}

func (s *Service) DeleteMovieById(movieId string) {
	s.repository.DeleteMovieById(movieId)
}
