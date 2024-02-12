package main

import (
	"net/http"

	"github.com/KonstantinMikhel/url-shortener/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Get("/{shortURL}", handlers.HandleRedirect)
	router.Post("/", handlers.HandleShorten)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
