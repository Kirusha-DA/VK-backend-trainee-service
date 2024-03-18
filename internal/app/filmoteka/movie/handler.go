package movie

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/handlers"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/middleware"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/models"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/movie/service"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

const (
	movieURL             = "/movie"
	moviesURL            = "/movies"
	moviesByNameURL      = moviesURL + "/{name}"
	moviesByIDURL        = moviesURL + "/{id}"
	moviesByActorNameURL = moviesURL + "/actors/{actor_name}"
)

type handler struct {
	middleware middleware.MiddleWare
	service    service.Service
	logger     *logging.Logger
}

func NewHandler(middleware middleware.MiddleWare, service service.Service, logger *logging.Logger) handlers.Handler {
	return &handler{
		middleware: middleware,
		service:    service,
		logger:     logger,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc(moviesURL, h.GetSortedMovies).Methods(http.MethodGet)
	router.HandleFunc(moviesByNameURL, h.GetMoviesByActorName).Methods(http.MethodGet)
	router.HandleFunc(movieURL, h.middleware.MiddleWare(h.CreateMovie)).Methods(http.MethodPost)
	router.HandleFunc(moviesByIDURL, h.middleware.MiddleWare(h.UpdateMovieById)).Methods(http.MethodPut)
	router.HandleFunc(moviesByIDURL, h.middleware.MiddleWare(h.PartiallyUpdateMovieById)).Methods(http.MethodPatch)
	router.HandleFunc(moviesByIDURL, h.middleware.MiddleWare(h.DeleteMovieById)).Methods(http.MethodDelete)
	router.HandleFunc(moviesByActorNameURL, h.GetMoviesByActorName).Methods(http.MethodGet)
}

// @Summary GetMoviesByName
// @Tags actors
// @Desctiption get movies by name
// @ID get-sorted-movies
// @Param q query string true "sort by query parram"
// @Sucess 200 {object} model.Movie true "Movie Info"
// @Router /movies [get]
func (h *handler) GetSortedMovies(w http.ResponseWriter, r *http.Request) {
	queryValue := r.URL.Query().Get("sort_by")
	sortedMovies, _ := h.retrieveMoviesSortedByQueryValue(queryValue)
	json, _ := json.MarshalIndent(sortedMovies, "", " ")
	w.Write(json)
}

func (h *handler) retrieveMoviesSortedByQueryValue(queryValue string) (movies []models.Movie, err error) {
	switch queryValue {
	case "rating":
		movies, err = h.service.GetMoviesSortedByRatingDESC()
	case "name":
		movies, err = h.service.GetMoviesSortedByNameDESC()
	case "release_date":
		movies, err = h.service.GetMoviesSortedByReleaseDateDESC()
	default:
		movies, err = nil, errors.New("no suitable query value")
	}
	return movies, err
}

// @Summary CreateMovie
// @Security ApiKeyAuth
// @Tags movies
// @Desctiption creates actor
// @ID create-movie
// @Accept json
// @Param input body models.Movie true "Movie Info"
// @Sucess 201
// @Failure 400
// @Router /movie [post]
func (h *handler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var movie models.Movie
	if err := json.Unmarshal(body, &movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	h.service.CreateMovie(&movie)

}

// @Summary UpdateMovieById
// @Security ApiKeyAuth
// @Tags movies
// @Desctiption updates whole movie
// @ID update-movie
// @Accept json
// @Param input body models.Movie true "Movie Info"
// @Param id path int true "Movie ID"
// @Sucess 204
// @Failure 400,404
// @Router /movies/{id} [put]
func (h *handler) UpdateMovieById(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var movie models.Movie
	if err := json.Unmarshal(body, &movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	movie.ID = uint(id)

	if err := h.service.UpdateMovieById(&movie); err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// @Summary PartiallyUpdateMovieById
// @Security ApiKeyAuth
// @Tags movies
// @Desctiption partially updates whole actor
// @ID partially-update-movie
// @Accept json
// @Param input body models.Movie true "Actor Info"
// @Param id path int true "Movie ID"
// @Sucess 204
// @Failure 400,404
// @Router /movies/{id} [patch]
func (h *handler) PartiallyUpdateMovieById(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var movie models.Movie
	if err := json.Unmarshal(body, &movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	movie.ID = uint(id)

	if err := h.service.PartiallyUpdateActor(&movie); errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// @Summary DeleteMovieById
// @Security ApiKeyAuth
// @Tags movies
// @Desctiption deletes movie by id
// @ID delete-movie
// @Param id path int true "Movie ID"
// @Sucess 204
// @Failure 404
// @Router /movies/{id} [delete]
func (h *handler) DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.service.DeleteMovieById(id); errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// @Summary GetMoviesByName
// @Tags movies
// @Desctiption get movies by name
// @ID get-movies-by-actor-name
// @Param actor_name path string true "Actor Name"
// @Sucess 200 {object} model.Movie true "Movie Info"
// @Failure 404
// @Router /movies/actors/{actor_name} [get]
func (h *handler) GetMoviesByActorName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["actor_name"]
	actorsByActorName, err := h.service.GetMoviesByActorName(name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	json, _ := json.MarshalIndent(actorsByActorName, "", " ")
	w.Write(json)
}
