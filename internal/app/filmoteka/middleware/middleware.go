package middleware

import (
	"net/http"
	"strings"

	"github.com/Kirusha-DA/VK-backend-trainee-service/internal/app/filmoteka/users-auth/service"
	"github.com/Kirusha-DA/VK-backend-trainee-service/pkg/logging"
)

type MiddleWare interface {
	MiddleWare(h appHandler) http.HandlerFunc
}

type middleWare struct {
	service service.Service
	logger  *logging.Logger
}

func NewMiddleWare(service service.Service, logger *logging.Logger) MiddleWare {
	return &middleWare{
		service: service,
		logger:  logger,
	}
}

type appHandler func(w http.ResponseWriter, r *http.Request)

func (m *middleWare) MiddleWare(h appHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] != nil {
			tokenWithBearer := r.Header["Authorization"][0]
			acessToken := strings.Split(tokenWithBearer, " ")[1]

			_, err := m.service.ParseTocken(acessToken)
			if err != nil {
				w.Write([]byte(err.Error()))
			}

			h(w, r)
		} else {

			w.Write([]byte("Not Authorized"))
		}
	})
}
