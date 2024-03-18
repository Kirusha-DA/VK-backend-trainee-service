package repository

import "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/models"

type Repository interface {
	CreateActor(actor *models.Actor) error
	UpdateActorById(actor *models.Actor) error
	PartiallyUpdateActorById(actor *models.Actor) error
	DeleteActorById(actorID string) error
	GetAllActorsWithMovies() ([]models.Actor, error)
}
