package endpoint

import (
	"net/http"

	"github.com/idkwhyureadthis/practice-project-backend/internal/app/service"
)

type Service interface {
	GetStatus() service.Response
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(w http.ResponseWriter, r *http.Request) {
	response := e.s.GetStatus()
	w.WriteHeader(response.Code)
	w.Write([]byte(response.Data))
}
