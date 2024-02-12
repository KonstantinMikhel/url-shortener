package main

import (
	"net/http"

	"github.com/KonstantinMikhel/url-shortener/internal/handlers"
)

func handleURL(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		handlers.HandleRedirect(writer, request)
	case http.MethodPost:
		handlers.HandleShorten(writer, request)
	default:
		http.Error(writer, "Only GET and POST requests are allowed!", http.StatusBadRequest)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handleURL)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
