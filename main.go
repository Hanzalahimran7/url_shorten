package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Post("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello World")) })
	http.ListenAndServe(":3000", router)
	fmt.Println("Hello World")
}
