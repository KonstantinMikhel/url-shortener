package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleShorten(t *testing.T) {
	// Create a request with a mock body
	requestBody := strings.NewReader("http://example.com/original")
	req, err := http.NewRequest("POST", "/shorten", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	shortener := NewURLShortener("http://localhost:8080")

	// Call the handler function with the created request and ResponseRecorder
	http.HandlerFunc(shortener.HandleShorten).ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body
	expected := "http://localhost:8080/"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHandleRedirect(t *testing.T) {
	shortener := NewURLShortener("http://localhost:8080")
	// Mock the URL mapping
	shortener.urls["abcdefg"] = "http://example.com/original"

	// Create a request with the shortened URL
	req, err := http.NewRequest("GET", "/abcdefg", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the created request and ResponseRecorder
	http.HandlerFunc(shortener.HandleRedirect).ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusTemporaryRedirect)
	}

	// Check the redirection location
	expectedLocation := "http://example.com/original"
	location := rr.Header().Get("Location")
	if location != expectedLocation {
		t.Errorf("handler returned unexpected redirection location: got %v want %v",
			location, expectedLocation)
	}
}
