package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/actor"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/config"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"github.com/gorilla/mux"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := mux.NewRouter()

	cfg := config.GetConfig()

	logger.Info("register actor handler")
	handler := actor.NewHandler(logger)
	handler.Register(router)

	run(router, cfg)
}

func run(router *mux.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("run app")

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
