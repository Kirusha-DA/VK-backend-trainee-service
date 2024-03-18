package repository

import "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/models"

type Repository interface {
	CreateUser(user *models.UserAuth) error
	GetUserByUsernameAndPassword(username, password string) (*models.UserAuth, error)
}
