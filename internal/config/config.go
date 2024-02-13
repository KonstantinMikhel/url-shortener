package config

import (
	"flag"
	"os"
)

type Config struct {
	HTTPServerAddress string
	BaseURL           string
}

func NewConfig() (cfg Config) {
	serverAddress := flag.String("a", "127.0.0.1:8080", "Address provided for HTTP server")
	baseURL := flag.String("b", "http://127.0.0.1:8080", "Base for shortened URL")

	flag.Parse()

	if envVar, present := os.LookupEnv("SERVER_ADDRESS"); present {
		*serverAddress = envVar
	}

	if envVar, present := os.LookupEnv("BASE_URL"); present {
		*baseURL = envVar
	}

	cfg = Config{
		HTTPServerAddress: *serverAddress,
		BaseURL:           *baseURL,
	}

	return cfg
}
