package urlshorten

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hanzalahimran7/url_shorten/controller"
)

type App struct {
	Router     *chi.Mux
	Controller controller.UserController
}

func Initialize() *App {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(time.Millisecond * 2000))
	return &App{
		Router:     router,
		Controller: controller.NewController(),
	}
}

func (a *App) Run() {
	a.loadRoutes()
	log.Println("Starting the server")
	http.ListenAndServe(":3000", a.Router)
}
