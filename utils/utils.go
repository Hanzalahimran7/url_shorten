package utils

import (
	"encoding/json"
	"net/http"
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
