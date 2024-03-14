package db

import (
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/actor"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"gorm.io/gorm"
)

type repository struct {
	db     *gorm.DB
	logger *logging.Logger
}

func (r *repository) CreateActor(actor *actor.Actor) {
	r.db.Create(actor)
}

func (r *repository) UpdateActor(actor *actor.Actor) {
	if result := r.db.First(actor); result.Error == nil {
		r.db.Create(&actor)
		return
	}
	r.db.Save(actor)
}

func (r *repository) PartiallyUpdateActor(actor *actor.Actor) {
	r.db.Select("*").Updates(map[string]interface{}{
		"ID":        actor.ID,
		"Sex":       actor.Sex,
		"BirthDate": actor.BirthDate,
	})
}

func (r *repository) DeleteActor(actorID string) {
	r.db.Delete(&actor.Actor{}, actorID)
}

func NewRepository(logger *logging.Logger) actor.Repository {
	return &repository{
		logger: logger,
	}
}
