package handlers

import (
	"io"
	"math/rand"
	"net/http"
	"path"
)

type URLShortener struct {
	urls    map[string]string
	baseURL string
}

func NewURLShortener(baseURL string) URLShortener {
	return URLShortener{
		urls:    map[string]string{},
		baseURL: baseURL,
	}
}

func (u URLShortener) HandleShorten(writer http.ResponseWriter, request *http.Request) {
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
	u.urls[shortURL] = originalURL

	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusCreated)

	fullShortenedURL := u.baseURL + "/" + shortURL
	writer.Write([]byte(fullShortenedURL))
}

func (u URLShortener) HandleRedirect(writer http.ResponseWriter, request *http.Request) {
	shortURL := path.Base(request.URL.Path)
	if shortURL == "." || shortURL == "/" {
		http.Error(writer, "Shortened URL is missing in a request", http.StatusBadRequest)
		return
	}

	// Retrieve the original URL
	originalURL, found := u.urls[shortURL]
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
