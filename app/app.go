package urlshorten

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hanzalahimran7/url_shorten/controller"
	"github.com/hanzalahimran7/url_shorten/store"
	"github.com/redis/go-redis/v9"
)

type App struct {
	Router     *chi.Mux
	Controller controller.UserController
	DB         *store.Database
}

func Initialize(db_options *redis.Options) *App {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(time.Millisecond * 2000))
	db, err := store.GetRedisInstance(db_options)
	if err != nil {
		log.Fatal("Cannot connect to Redis")
	}
	return &App{
		Router:     router,
		Controller: controller.NewController(db),
		DB:         &db,
	}
}

func (a *App) Run() {
	a.loadRoutes()
	log.Println("Starting the server")
	http.ListenAndServe(":3000", a.Router)
}
