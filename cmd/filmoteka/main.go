package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/actor"
	actorRepository "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/actor/repository"
	actorService "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/actor/service"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/config"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/middleware"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/models"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/movie"
	movieRepository "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/movie/repository"
	movieService "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/movie/service"
	usersauth "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/users-auth"
	authRepository "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/users-auth/repository"
	authService "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/users-auth/service"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/dbclient/postgres"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"github.com/gorilla/mux"
)

// @title Filmoteka API
// @version 1.0
// @desctiption API server for Filmoteka Application

// @host localhost:8080
// @BasePath/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logger := logging.GetLogger()

	router := mux.NewRouter()

	cfg := config.GetConfig()

	connection := postgres.NewClient(logger, &config.StorageConfig{
		Host:     cfg.Storage.Host,
		Port:     cfg.Storage.Port,
		DB:       cfg.Storage.DB,
		User:     cfg.Storage.User,
		Password: cfg.Storage.Password,
	})

	if err := connection.AutoMigrate(&models.Actor{}, &models.Movie{}, &models.UserAuth{}); err != nil {
		logger.Fatal(err)
	}

	actorRepo := actorRepository.NewActorRepository(connection, logger)
	movieRepo := movieRepository.NewRepository(connection, logger)
	userAuthRepo := authRepository.NewRepository(connection, logger)

	actorService := actorService.NewActorService(actorRepo, logger)
	movieService := movieService.NewService(movieRepo, logger)
	userAuthService := authService.NewService(userAuthRepo, logger)

	middleWare := middleware.NewMiddleWare(userAuthService, logger)

	actorHandler := actor.NewHandler(middleWare, actorService, logger)
	movieHandler := movie.NewHandler(middleWare, movieService, logger)
	userAuthHandler := usersauth.NewHandler(userAuthService, logger)

	actorHandler.Register(router)
	movieHandler.Register(router)
	userAuthHandler.Register(router)

	run(router, cfg)
}

func run(router *mux.Router, cfg *config.Config) {
	logger := logging.GetLogger()

	var listener net.Listener
	var listenErr error

	listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Infof("server is listening %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	logger.Fatal(server.Serve(listener))
}
