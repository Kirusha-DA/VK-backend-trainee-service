package service

import "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/models"

type Service interface {
	CreateUser(user *models.UserAuth) error
	GenerateToken(username, password string) (string, error)
	ParseTocken(token string) (int, error)
}
