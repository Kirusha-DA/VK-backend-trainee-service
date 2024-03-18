package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/models"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/users-auth/repository"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"github.com/golang-jwt/jwt/v5"
)

const (
	salt       = "klasdjof89sadfkl34lkjk"
	signingKey = "234j2l34k#$@#$lkjskjdfj"
	tokenTTL   = 12 * time.Hour
)

type service struct {
	repository repository.Repository
	logger     *logging.Logger
}

func NewService(repository repository.Repository, logger *logging.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s *service) CreateUser(user *models.UserAuth) error {
	user.Password = s.generatePasswordHash(user.Password)
	if err := s.repository.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *service) GenerateToken(username, password string) (string, error) {
	userAuth, err := s.repository.GetUserByUsernameAndPassword(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp":    tokenTTL,
			"userId": userAuth.ID,
		})

	tokenStr, _ := token.SignedString([]byte(signingKey))
	return tokenStr, nil
}

func (s *service) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *service) ParseTocken(acessToken string) (int, error) {
	token, err := jwt.Parse(acessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token signed method is different")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("token is not valid")
	}

	return 0, nil
}
