package main

import (
	"net/http"

	"github.com/KonstantinMikhel/url-shortener/internal/config"
	"github.com/KonstantinMikhel/url-shortener/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.NewConfig()
	shortener := handlers.NewURLShortener(cfg.BaseURL)

	router := chi.NewRouter()
	router.Get("/{shortURL}", shortener.HandleRedirect)
	router.Post("/", shortener.HandleShorten)

	err := http.ListenAndServe(cfg.HTTPServerAddress, router)
	if err != nil {
		panic(err)
	}
}
