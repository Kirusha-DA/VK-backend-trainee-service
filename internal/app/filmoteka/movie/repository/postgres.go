package db

import (
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/movie"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/models"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"gorm.io/gorm"
)

type repository struct {
	db     *gorm.DB
	logger *logging.Logger
}

func NewRepository(db *gorm.DB, logger *logging.Logger) movie.Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) GetMoviesSortedByRatingDESC() (movies []models.Movie) {
	if err := r.db.Order("rating desc").Find(&movies).Error; err != nil {
		r.logger.Info(err)
	}
	return movies
}

func (r *repository) GetMoviesSortedByNameDESC() (movies []models.Movie) {
	if err := r.db.Order("name desc").Find(&movies).Error; err != nil {
		r.logger.Info(err)
	}
	return movies
}

func (r *repository) GetMoviesSortedByReleaseDateDESC() (movies []models.Movie) {
	if err := r.db.Order("release_date desc").Find(&movies).Error; err != nil {
		r.logger.Info(err)
	}
	return movies
}

func (r *repository) GetMoviesByName(movieName string) (movies []models.Movie) {
	if err := r.db.Where("name = ?", movieName).Find(&movies).Error; err != nil {
		r.logger.Info(err)
	}
	return movies
}

func (r *repository) GetMoviesByActorName(actorName string) (movies []models.Movie) {
	if err := r.db.Model(&models.Movie{}).Preload("Actors", "name = ?", actorName).Find(&movies); err != nil {
		r.logger.Info(err)
	}
	return movies
}

func (r *repository) CreateMovie(movie *models.Movie) {
	if err := r.db.Model(&models.Movie{}).Create(movie).Error; err != nil {
		r.logger.Info(err)
	}
}

func (r *repository) UpdateMovieById(movie *models.Movie) {
	if err := r.db.Save(movie).Error; err != nil {
		r.logger.Info(err)
	}
}

func (r *repository) PartiallyUpdateActor(movie *models.Movie) {
	if err := r.db.Model(movie).Updates(movie); err != nil {
		r.logger.Info(err)
	}
}

func (r *repository) DeleteMovieById(movieId string) {
	if err := r.db.Delete(&models.Movie{}, movieId); err != nil {
		r.logger.Info(err)
	}
}
