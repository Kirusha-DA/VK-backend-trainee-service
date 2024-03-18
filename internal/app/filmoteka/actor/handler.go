package actor

import (
	"encoding/json"
	"net/http"
	"strconv"

	actor_service "github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/actor/service"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/handlers"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/middleware"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/models"
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
	middleware middleware.MiddleWare
	service    actor_service.Service
	logger     *logging.Logger
}

func NewHandler(middleware middleware.MiddleWare, service actor_service.Service, logger *logging.Logger) handlers.Handler {
	return &handler{
		middleware: middleware,
		service:    service,
		logger:     logger,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc(actorURL, h.middleware.MiddleWare(h.CreateActor)).Methods(http.MethodPost)
	router.HandleFunc(actorsByIDURL, h.middleware.MiddleWare(h.UpdateActorById)).Methods(http.MethodPut)
	router.HandleFunc(actorsByIDURL, h.middleware.MiddleWare(h.PartiallyUpdateActorById)).Methods(http.MethodPatch)
	router.HandleFunc(actorsByIDURL, h.middleware.MiddleWare(h.DeleteActorById)).Methods(http.MethodDelete)
	router.HandleFunc(actorsWithMoviesURL, h.GetAllActorsWithMovies).Methods(http.MethodGet)
}

// @Summary CreateActor
// @Security ApiKeyAuth
// @Tags actors
// @Desctiption creates actor
// @ID create-actor
// @Accept json
// @Param input body models.Actor true "List Info"
// @Sucess 201
// @Failure 400
// @Router /actor [post]
func (h *handler) CreateActor(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var actor models.Actor
	errBadRequest := json.Unmarshal(body, &actor)
	if errBadRequest != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	h.service.CreateActor(&actor)
}

// @Summary UpdateActorById
// @Security ApiKeyAuth
// @Tags actors
// @Desctiption updates whole actor
// @ID update-actor
// @Accept json
// @Param input body models.Actor true "Actor Info"
// @Param id path int true "Account ID"
// @Sucess 204
// @Failure 400,404
// @Router /actors/{id} [put]
func (h *handler) UpdateActorById(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var actor models.Actor
	errBadRequest := json.Unmarshal(body, &actor)
	if errBadRequest != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	actor.ID = uint(id)

	if err := h.service.UpdateActorById(&actor); err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// @Summary PartiallyUpdateActorById
// @Security ApiKeyAuth
// @Tags actors
// @Desctiption partially updates whole actor
// @ID partially-update-actor
// @Accept json
// @Param input body models.Actor true "Actor Info"
// @Param id path int true "Actor ID"
// @Sucess 204
// @Failure 400,404
// @Router /actors/{id} [patch]
func (h *handler) PartiallyUpdateActorById(w http.ResponseWriter, r *http.Request) {

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var actor models.Actor
	badRequestErr := json.Unmarshal(body, &actor)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	actor.ID = uint(id)

	recordNotFoundErr := h.service.PartiallyUpdateActorById(&actor)

	if badRequestErr != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else if recordNotFoundErr != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}

}

// @Summary DeleteActorById
// @Security ApiKeyAuth
// @Tags actors
// @Desctiption deletes actor by id
// @ID delete-actor
// @Param id path int true "Actor ID"
// @Sucess 204
// @Failure 404
// @Router /actors/{id} [delete]
func (h *handler) DeleteActorById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.service.DeleteActorById(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// @Summary GetAllActorsWithMovies
// @Tags actors
// @Desctiption gets all actors with movies actor
// @ID get-actor-movie
// @Produce json
// @Sucess 200 {object} model.Actor true "Actor Info"
// @Router /actors/movies [get]
func (h *handler) GetAllActorsWithMovies(w http.ResponseWriter, r *http.Request) {
	actorsWithMovies, _ := h.service.GetAllActorsWithMovies()
	json, _ := json.MarshalIndent(actorsWithMovies, "", " ")
	w.Write(json)
}
