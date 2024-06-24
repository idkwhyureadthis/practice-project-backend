package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/idkwhyureadthis/practice-project-backend/internal/app/endpoint"
	"github.com/idkwhyureadthis/practice-project-backend/internal/app/service"
)

type App struct {
	e *endpoint.Endpoint
	s *service.Service
	c *chi.Mux
}

func New() (*App, error) {
	a := &App{}
	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.c = chi.NewRouter()
	a.c.Get("/status", a.e.Status)
	return a, nil
}

func (a *App) Run(port int) error {
	portStr := fmt.Sprintf(":%d", port)
	fmt.Println("server running at ", portStr)
	srv := &http.Server{
		Addr:    portStr,
		Handler: a.c,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start http server %v", err)
	}

	return nil
}
