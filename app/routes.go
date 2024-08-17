package urlshorten

import (
	"github.com/hanzalahimran7/url_shorten/utils"
)

func (a *App) loadRoutes() {
	a.Router.Get("/", utils.ApiFunc(a.Controller.HelloWorld))
}
