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

func NewRepository(db *gorm.DB, logger *logging.Logger) Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) GetMoviesSortedByRatingDESC() ([]models.Movie, error) {
	var movies []models.Movie
	if err := r.db.Order("rating desc").Find(&movies).Error; err != nil {
		return nil, fmt.Errorf("failed to get movies sorted by rating desc: %s", err)
	}
	return movies, nil
}

func (r *repository) GetMoviesSortedByNameDESC() ([]models.Movie, error) {
	var movies []models.Movie
	if err := r.db.Order("name desc").Find(&movies).Error; err != nil {
		return nil, fmt.Errorf("failed to get movies sorted by names: %s", err)
	}
	return movies, nil
}

func (r *repository) GetMoviesSortedByReleaseDateDESC() ([]models.Movie, error) {
	var movies []models.Movie
	if err := r.db.Order("release_date desc").Find(&movies).Error; err != nil {
		return nil, fmt.Errorf("failed to get movies sorted by names: %s", err)
	}
	return movies, nil
}

func (r *repository) GetMoviesByName(movieName string) ([]models.Movie, error) {
	var movies []models.Movie
	if err := r.db.Where("name = ?", movieName).Find(&movies).Error; err != nil {
		return nil, fmt.Errorf("failed to get movies by name: %s", err)
	}
	return movies, nil
}

func (r *repository) GetMoviesByActorName(actorName string) ([]models.Movie, error) {
	var movies []models.Movie
	if err := r.db.Model(&models.Movie{}).Preload("Actors", "name = ?", actorName).Find(&movies).Error; err != nil {
		return nil, fmt.Errorf("failed to get movies by actor name: %s", err)
	}
	return movies, nil
}

func (r *repository) CreateMovie(movie *models.Movie) error {
	if err := r.db.Model(&models.Movie{}).Create(movie).Error; err != nil {
		return fmt.Errorf("failed to create movie: %s", err)
	}
	return nil
}

func (r *repository) UpdateMovieById(movie *models.Movie) error {
	if err := r.db.Save(movie).Error; err != nil {
		return fmt.Errorf("failed to update movie: %s", err)
	}
	return nil
}

func (r *repository) PartiallyUpdateActor(movie *models.Movie) error {
	if err := r.db.Model(movie).Updates(movie).First(movie).Error; err != nil {
		return fmt.Errorf("failed to partially update actor: %s", err)
	}
	return nil
}

func (r *repository) DeleteMovieById(movieId string) error {
	movieIdI, _ := strconv.Atoi(movieId)
	if err := r.db.First(&models.Movie{}, movieIdI).Error; err != nil {
		return fmt.Errorf("failed to delete actor: %s", err)
	}
	r.db.Delete(&models.Movie{}, movieId)
	return nil
}
