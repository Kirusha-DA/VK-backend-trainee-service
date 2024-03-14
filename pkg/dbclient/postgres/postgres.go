package postgres

import (
	"fmt"

	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/config"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient(logger *logging.Logger, sc *config.StorageConfig) (connection *gorm.DB) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", sc.Host, sc.User, sc.Password, sc.DB, sc.Port)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Fatal("Failed to connect to postgres")
	}
	return
}
