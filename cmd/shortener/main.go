package main

import (
	"io"
	"math"
	"net/http"
	"slices"
	"strings"
)

type ShortenedURL struct {
	LongURL  string
	ShortURL string
}

var urls []ShortenedURL = make([]ShortenedURL, 0)

const base62Digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func DecimalToBase62(n int64) string {
	if n == 0 {
		return "0"
	}

	base62 := make([]byte, 0)
	radix := int64(62)

	for n > 0 {
		remainder := n % radix
		base62 = append([]byte{base62Digits[remainder]}, base62...)
		n /= radix
	}

	return string(base62)
}

func Base62ToDecimal(s string) int64 {
	var decimalNumber int64

	for i, c := range s {
		decimalNumber += int64(strings.IndexByte(base62Digits, byte(c))) * int64(math.Pow(62, float64(len(s)-i-1)))
	}

	return decimalNumber
}

func handleURL(wr http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		handleOriginalURL(wr, req)
	case http.MethodPost:
		handleShortenedURL(wr, req)
	default:
		http.Error(wr, "Only GET and POST requests are allowed!", http.StatusBadRequest)
	}
}

func handleShortenedURL(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "text/plain")
	wr.WriteHeader(http.StatusCreated)

	body, _ := io.ReadAll(req.Body)
	longURL := string(body)

	idx := slices.IndexFunc(urls, func(url ShortenedURL) bool {
		return url.LongURL == longURL
	})

	if idx == -1 {
		urls = append(urls, ShortenedURL{
			LongURL:  longURL,
			ShortURL: DecimalToBase62(int64(len(urls)) + 100000),
		})
		wr.Write([]byte(`http://localhost:8080/` + urls[len(urls)-1].ShortURL))
	} else {
		wr.Write([]byte(`http://localhost:8080/` + urls[idx].ShortURL))
	}
}

// GET
func handleOriginalURL(wr http.ResponseWriter, req *http.Request) {
	urlParts := strings.Split(req.URL.Path, "/")
	shortURL := urlParts[len(urlParts)-1]

	idx := slices.IndexFunc(urls, func(url ShortenedURL) bool {
		return url.ShortURL == shortURL
	})

	if idx == -1 {
		http.Error(wr, "Only GET and POST requests are allowed!", http.StatusBadRequest)
	}

	wr.Header().Set("Location", urls[idx].LongURL)
	wr.WriteHeader(http.StatusTemporaryRedirect)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handleURL)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
