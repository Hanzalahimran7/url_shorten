package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/hanzalahimran7/url_shorten/model"
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

func (uc *UserController) CreateURL(w http.ResponseWriter, r *http.Request) (int, error) {
	// Define the struct with JSON tags for correct decoding
	type urlRequest struct {
		URL string `json:"url"`
	}

	var requestModel urlRequest

	// Decode the request body into the struct
	if err := json.NewDecoder(r.Body).Decode(&requestModel); err != nil {
		return http.StatusBadRequest, fmt.Errorf("invalid request body")
	}

	// Check if the URL field is empty (optional but recommended)
	if requestModel.URL == "" {
		return http.StatusBadRequest, fmt.Errorf("url field is required")
	}

	// Parse the URL to check if its valid
	_, err := url.ParseRequestURI(requestModel.URL)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("INVALID URL: %v", requestModel.URL)
	}

	created_at := time.Now().UTC()
	expired_at := time.Now().UTC().Add(time.Hour * 2)

	body := model.Url{
		Id:           uuid.New(),
		CreateAt:     &created_at,
		ExpiredAt:    &expired_at,
		OriginalUrl:  requestModel.URL,
		ShortenedUrl: utils.CreateShortenedURL(),
	}

	utils.WriteJSON(w, http.StatusCreated, body)

	// Successfully decoded and validated the request
	return http.StatusOK, nil
}
