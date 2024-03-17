package movie

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
	movieURL             = "/movie"
	moviesURL            = "/movies"
	moviesByNameURL      = moviesURL + "/{name}"
	moviesByIDURL        = moviesURL + "/{id}"
	moviesByActorNameURL = moviesURL + "/actors/{actor_name}"
)

type handler struct {
	service Service
	logger  *logging.Logger
}

func NewHandler(service *Service, logger *logging.Logger) handlers.Handler {
	return &handler{
		service: *service,
		logger:  logger,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc(moviesURL, h.GetSortedMovies).Methods(http.MethodGet)
	router.HandleFunc(moviesByNameURL, h.GetMoviesByName).Methods(http.MethodGet)
	router.HandleFunc(movieURL, h.CreateMovie).Methods(http.MethodPost)
	router.HandleFunc(moviesByIDURL, h.UpdateMovie).Methods(http.MethodPut)
	router.HandleFunc(moviesByIDURL, h.PartiallyUpdateMovies).Methods(http.MethodPatch)
	router.HandleFunc(moviesByIDURL, h.DeleteMovie).Methods(http.MethodDelete)
	router.HandleFunc(moviesByActorNameURL, h.GetMoviesByActorName).Methods(http.MethodGet)
}

func (h *handler) GetSortedMovies(w http.ResponseWriter, r *http.Request) {
	queryValue := r.URL.Query().Get("sort_by")
	sortedMovies := h.retrieveMoviesSortedBy(queryValue)
	json, _ := json.MarshalIndent(sortedMovies, "", " ")
	w.Write(json)
}

func (h *handler) retrieveMoviesSortedBy(queryValue string) (movies []models.Movie) {
	switch queryValue {
	case "rating":
		movies = h.service.GetMoviesSortedByRatingDESC()
	case "name":
		movies = h.service.GetMoviesSortedByNameDESC()
	case "release_date":
		movies = h.service.GetMoviesSortedByReleaseDateDESC()
	}
	return movies
}

func (h *handler) GetMoviesByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	movies := h.service.GetMoviesByActorName(name)
	json, _ := json.MarshalIndent(movies, "", " ")
	w.Write(json)
}

func (h *handler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var movie models.Movie
	if err := json.Unmarshal(body, &movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	h.logger.Info(movie)

	h.service.CreateMovie(&movie)
}

func (h *handler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
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

	h.logger.Info(movie)

	h.service.UpdateMovieById(&movie)
}

func (h *handler) PartiallyUpdateMovies(w http.ResponseWriter, r *http.Request) {
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

	h.logger.Info(movie)

	h.service.PartiallyUpdateActor(&movie)
}

func (h *handler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	h.service.DeleteMovieById(id)
}

func (h *handler) GetMoviesByActorName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["actor_name"]
	actorsByActorName := h.service.GetMoviesByActorName(name)
	h.logger.Info(name)
	json, _ := json.MarshalIndent(actorsByActorName, "", " ")
	w.Write(json)
}
