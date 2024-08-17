package urlshorten

import "net/http"

func (a *App) loadRoutes() {
	a.Router.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello world")) })
}
