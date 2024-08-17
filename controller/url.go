package controller

import (
	"net/http"

	"github.com/hanzalahimran7/url_shorten/utils"
)

type UserController struct{}

func NewController() UserController {
	return UserController{}
}

func (uc *UserController) HelloWorld(w http.ResponseWriter, r *http.Request) (int, error) {
	utils.WriteJSON(w, http.StatusOK, "Hello world")
	return 0, nil
}
