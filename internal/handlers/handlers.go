package handlers

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

var urls = make(map[string]string)

func HandleShorten(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalid request method", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Could not read request body", http.StatusBadRequest)
		return
	}

	originalURL := string(body)
	if originalURL == "" {
		http.Error(writer, "Original URL is missing in a request body", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL()
	urls[shortURL] = originalURL

	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusCreated)

	fullShortenedURL := fmt.Sprintf("http://localhost:8080/%s", shortURL)
	writer.Write([]byte(fullShortenedURL))
}

func HandleRedirect(writer http.ResponseWriter, request *http.Request) {
	shortURL := strings.TrimPrefix(request.URL.Path, "/")
	if shortURL == "" {
		http.Error(writer, "Shortened URL is missing in a request", http.StatusBadRequest)
		return
	}

	// Retrieve the original URL
	originalURL, found := urls[shortURL]
	if !found {
		http.Error(writer, "Shortened URL not found", http.StatusBadRequest)
		return
	}

	// Redirect the user to the original URL
	http.Redirect(writer, request, originalURL, http.StatusTemporaryRedirect)
}

func generateShortURL() string {
	const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const shortURLLength = 8

	shortURL := make([]byte, shortURLLength)
	for i := range shortURL {
		shortURL[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(shortURL)
}
