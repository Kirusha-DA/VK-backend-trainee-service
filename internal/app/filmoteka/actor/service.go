package actor

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

func (s *Service) CreateActor(actor *models.Actor) {
	s.repository.CreateActor(actor)
}

func (s *Service) UpdateActorById(actor *models.Actor) {
	s.repository.UpdateActor(actor)
}

func (s *Service) PartiallyUpdateActorById(actor *models.Actor) {
	s.repository.PartiallyUpdateActor(actor)
}

func (s *Service) DeleteActorById(actorID string) {
	s.repository.DeleteActor(actorID)
}

func (s *Service) GetAllActorsWithMovies() []models.Actor {
	return s.repository.GetAllActorsWithMovies()
}
