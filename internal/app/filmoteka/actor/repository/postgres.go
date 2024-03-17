package db

import (
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/actor"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/models"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"gorm.io/gorm"
)

type repository struct {
	db     *gorm.DB
	logger *logging.Logger
}

func NewRepository(db *gorm.DB, logger *logging.Logger) actor.Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) CreateActor(actor *models.Actor) {
	if err := r.db.Model(&models.Actor{}).Create(actor); err != nil {
		r.logger.Info(err)
	}
}

func (r *repository) UpdateActor(actor *models.Actor) {
	if err := r.db.Save(actor); err != nil {
		r.logger.Info(err)
	}
}

func (r *repository) PartiallyUpdateActor(actor *models.Actor) {
	if err := r.db.Model(actor).Updates(actor); err != nil {
		r.logger.Info(err)
	}
}

func (r *repository) DeleteActor(actorID string) {
	if err := r.db.Delete(&models.Actor{}, actorID); err != nil {
		r.logger.Info(err)
	}
}

func (r *repository) GetAllActorsWithMovies() (actors []models.Actor) {
	if err := r.db.Model(&models.Actor{}).Preload("Movies").Find(&actors); err != nil {
		r.logger.Info(err)
	}
	return actors
}
