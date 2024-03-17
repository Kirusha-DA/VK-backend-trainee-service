package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/actor"
	dbActor "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/actor/repository"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/config"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/movie"
	dbMovie "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/movie/repository"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/models"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/dbclient/postgres"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"github.com/gorilla/mux"
)

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

	if err := connection.AutoMigrate(&models.Actor{}, &models.Movie{}); err != nil {
		logger.Fatal(err)
	}

	actorRepo := dbActor.NewRepository(connection, logger)
	movieRepo := dbMovie.NewRepository(connection, logger)

	actorService := actor.NewService(actorRepo, logger)
	movieService := movie.NewService(movieRepo, logger)

	actorHandler := actor.NewHandler(actorService, logger)
	movieHandler := movie.NewHandler(movieService, logger)

	actorHandler.Register(router)
	movieHandler.Register(router)

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
