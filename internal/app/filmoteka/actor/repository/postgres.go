package repository

import (
	"fmt"
	"strconv"

	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/models"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"gorm.io/gorm"
)

type repository struct {
	db     *gorm.DB
	logger *logging.Logger
}

func NewActorRepository(db *gorm.DB, logger *logging.Logger) Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) CreateActor(actor *models.Actor) error {
	if err := r.db.Model(&models.Actor{}).Create(actor).Error; err != nil {
		return fmt.Errorf("failed to create actor: %s", err)
	}
	return nil
}

func (r *repository) UpdateActorById(actor *models.Actor) error {
	if err := r.db.Save(actor).Error; err != nil {
		return fmt.Errorf("failed to update actor: %s", err)
	}
	return nil
}

func (r *repository) PartiallyUpdateActorById(actor *models.Actor) error {
	if err := r.db.Model(actor).Updates(actor).First(actor).Error; err != nil {
		return fmt.Errorf("failed to patrially update actor: %s", err)
	}
	return nil
}

func (r *repository) DeleteActorById(actorId string) error {
	actorIdI, _ := strconv.Atoi(actorId)
	if err := r.db.First(&models.Actor{}, actorIdI).Error; err != nil {
		return fmt.Errorf("failed to delete actor: %s", err)
	}
	r.db.Delete(&models.Actor{}, actorIdI)
	return nil
}

func (r *repository) GetAllActorsWithMovies() ([]models.Actor, error) {
	var actors []models.Actor
	if err := r.db.Model(&models.Actor{}).Preload("Movies").Find(&actors).Error; err != nil {
		r.logger.Info(err)
		return nil, fmt.Errorf("failed to get all actors with movies: %s", err)
	}
	return actors, nil
}
