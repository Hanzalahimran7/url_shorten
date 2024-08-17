package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) (int, error)

type APIError struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func ApiFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, err := f(w, r)
		if err != nil {
			WriteJSON(w, status, APIError{Error: err.Error()})
		}
	}
}

func GenerateShortID() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// Generate a random number (e.g., 64-bit integer)
	randomNumber := r.Int63()

	// Encode the random number using base62 or base64
	encoded := base64.RawURLEncoding.EncodeToString([]byte(fmt.Sprintf("%d", randomNumber)))

	// Truncate or handle the length if needed (e.g., first 6-8 characters)
	shortID := encoded[:8]

	return shortID
}

func CreateShortenedURL() string {
	shortID := GenerateShortID()

	// Store URL data in the database here

	// Construct the shortened URL using the base URL and the short ID
	shortenedURL := fmt.Sprintf("%s", shortID)

	return shortenedURL
}
