package main

import (
	"net/http"
)

func handleUrl(wr http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		handleOriginalUrl(wr, req)
	case http.MethodPost:
		handleShortenedUrl(wr, req)
	default:
		http.Error(wr, "Only GET and POST requests are allowed!", http.StatusBadRequest)
	}
}

func handleShortenedUrl(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "text/plain")
	wr.WriteHeader(http.StatusCreated)
	wr.Write([]byte(`http://localhost:8080/EwHXdJfB`))
}

func handleOriginalUrl(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Location", "https://practicum.yandex.ru/")
	wr.WriteHeader(http.StatusTemporaryRedirect)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handleUrl)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
