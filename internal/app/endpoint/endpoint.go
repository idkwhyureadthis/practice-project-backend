package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/idkwhyureadthis/practice-project-backend/internal/app/service"
)

type Service interface {
	GetStatus() service.Response
	GetVacancies() service.Response
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
	answ, err := json.Marshal(response.Data)
	if err != nil {
		w.Write([]byte("error parsing json"))
	}
	w.Write((answ))
}

func (e *Endpoint) Jobs(w http.ResponseWriter, r *http.Request) {
	response := e.s.GetVacancies()
	w.WriteHeader(response.Code)
	answ, err := json.Marshal(response.Data)
	if err != nil {
		w.Write([]byte("failed to get vacancies"))
	}
	w.Write(answ)
}
