package service

import (
	actor_repository "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/actor/repository"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/models"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
)

type service struct {
	repository actor_repository.Repository
	logger     *logging.Logger
}

func NewActorService(repository actor_repository.Repository, logger *logging.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s *service) CreateActor(actor *models.Actor) error {
	return s.repository.CreateActor(actor)
}

func (s *service) UpdateActorById(actor *models.Actor) error {
	return s.repository.UpdateActorById(actor)
}

func (s *service) PartiallyUpdateActorById(actor *models.Actor) error {
	return s.repository.PartiallyUpdateActorById(actor)
}

func (s *service) DeleteActorById(actorID string) error {
	return s.repository.DeleteActorById(actorID)
}

func (s *service) GetAllActorsWithMovies() ([]models.Actor, error) {
	return s.repository.GetAllActorsWithMovies()
}
