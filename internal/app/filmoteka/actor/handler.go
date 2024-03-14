package actor

import (
	"log"
	"net/http"

	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/handlers"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"github.com/gorilla/mux"
)

const (
	actorURL            = "/actor"
	actorsByIDURL       = "/actors/{id}"
	actorsWithMoviesURL = "/actors/movies"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc(actorURL, h.CreateActor).Methods(http.MethodPost)
	router.HandleFunc(actorsByIDURL, h.UpdateActor).Methods(http.MethodPost)
	router.HandleFunc(actorsByIDURL, h.PartiallyUpdateActor).Methods(http.MethodPatch)
	router.HandleFunc(actorsByIDURL, h.DeleteActor).Methods(http.MethodDelete)
	router.HandleFunc(actorsWithMoviesURL, h.GetActorsWithMovies).Methods(http.MethodGet)
}

func (h *handler) CreateActor(w http.ResponseWriter, r *http.Request) {
	log.Println("Create actor")
	w.Write([]byte("LOL"))
}

func (h *handler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	log.Println("Update actor")
	w.Write([]byte("LOL"))
}
func (h *handler) PartiallyUpdateActor(w http.ResponseWriter, r *http.Request) {
	log.Println("Partially update actor")
	w.Write([]byte("LOL"))
}
func (h *handler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete actor")
	w.Write([]byte("LOL"))
}

func (h *handler) GetActorsWithMovies(w http.ResponseWriter, r *http.Request) {
	log.Println("Get actors with list of movies")
	w.Write([]byte("LOL"))
}
