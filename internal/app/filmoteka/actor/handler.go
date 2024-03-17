package actor

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/handlers"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/models"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"github.com/gorilla/mux"
)

const (
	actorURL            = "/actor"
	actorsURL           = "/actors"
	actorsByIDURL       = actorsURL + "/{id}"
	actorsWithMoviesURL = actorsURL + "/movies"
)

type handler struct {
	service *Service
	logger  *logging.Logger
}

func NewHandler(service *Service, logger *logging.Logger) handlers.Handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc(actorURL, h.CreateActor).Methods(http.MethodPost)
	router.HandleFunc(actorsByIDURL, h.UpdateActor).Methods(http.MethodPut)
	router.HandleFunc(actorsByIDURL, h.PartiallyUpdateActor).Methods(http.MethodPatch)
	router.HandleFunc(actorsByIDURL, h.DeleteActor).Methods(http.MethodDelete)
	router.HandleFunc(actorsWithMoviesURL, h.GetAllActorsWithMovies).Methods(http.MethodGet)
}

func (h *handler) CreateActor(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var actor models.Actor
	if err := json.Unmarshal(body, &actor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	h.service.CreateActor(&actor)
}

func (h *handler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var actor models.Actor
	if err := json.Unmarshal(body, &actor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	actor.ID = uint(id)

	h.logger.Info(actor)

	h.service.UpdateActorById(&actor)
}

func (h *handler) PartiallyUpdateActor(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var actor models.Actor
	if err := json.Unmarshal(body, &actor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	actor.ID = uint(id)

	h.logger.Info(actor)

	h.service.PartiallyUpdateActorById(&actor)
}
func (h *handler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	h.service.DeleteActorById(id)
}

func (h *handler) GetAllActorsWithMovies(w http.ResponseWriter, r *http.Request) {
	actorsWithMovies := h.service.GetAllActorsWithMovies()
	json, _ := json.MarshalIndent(actorsWithMovies, "", " ")
	w.Write(json)
}
