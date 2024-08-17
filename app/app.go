package urlshorten

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	Router *chi.Mux
}

func Initialize() *App {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Timeout(time.Millisecond * 2000))
	return &App{
		Router: router,
	}
}

func (a *App) Run() {
	a.loadRoutes()
	log.Println("Starting the server")
	http.ListenAndServe(":3000", a.Router)
}
