package repository

import (
	"fmt"

	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/models"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"gorm.io/gorm"
)

type repository struct {
	db     *gorm.DB
	logger *logging.Logger
}

func NewRepository(db *gorm.DB, logger *logging.Logger) Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) CreateUser(user *models.UserAuth) error {
	if err := r.db.Model(&models.UserAuth{}).Create(user).Error; err != nil {
		return fmt.Errorf("failed to create auth user: %s", err)
	}
	return nil
}

func (r *repository) GetUserByUsernameAndPassword(username, password string) (*models.UserAuth, error) {
	var userAuth *models.UserAuth
	if err := r.db.Where("username = ?", username).Where("password = ?", password).First(&userAuth).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by username and password: %s", err)
	}
	return userAuth, nil
}
