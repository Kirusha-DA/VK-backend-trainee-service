package actor

import "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/models"

type Repository interface {
	CreateActor(actor *models.Actor)
	UpdateActor(actor *models.Actor)
	PartiallyUpdateActor(actor *models.Actor)
	DeleteActor(actorID string)
	GetAllActorsWithMovies() []models.Actor
}
