package usersauth

import (
	"encoding/json"
	"net/http"

	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/handlers"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/models"
	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/users-auth/service"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
	"github.com/gorilla/mux"
)

const (
	authURL   = "/auth"
	signUpURL = authURL + "/sign-up"
	signInURL = authURL + "/sign-in"
)

type handler struct {
	service service.Service
	logger  *logging.Logger
}

func NewHandler(service service.Service, logger *logging.Logger) handlers.Handler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc(signUpURL, h.SignUp).Methods(http.MethodPost)
	router.HandleFunc(signInURL, h.SignIn).Methods(http.MethodGet)
}

// @Summary SignUpUser
// @Tags auth
// @Desctiption sign up user
// @ID sing-up-user
// @Accepts json
// @Param input body models.UserAuth true "User Auth"
// @Sucess 200
// @Failure 400
// @Router /auth/sign-up [post]
func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var user models.UserAuth
	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	h.service.CreateUser(&user)
}

// @Summary SignInUser
// @Tags movies
// @Desctiption sign in user
// @ID sing-in-user
// @Param username query string true "Username"
// @Param password query string true "Password"
// @Sucess 200 {object} struct{token string} true "Movie Info"
// @Failure 404
// @Router /auth/sign-in [post]
func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	token, err := h.service.GenerateToken(username, password)
	if err != nil {
		json, _ := json.MarshalIndent(map[string]string{"error": "failed to generate token"}, "", " ")
		w.Write([]byte(json))
	} else {
		json, _ := json.MarshalIndent(map[string]string{"token": token}, "", " ")
		w.Write([]byte(json))
	}
}
