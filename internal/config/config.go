package config

import "flag"

type Config struct {
	HTTPServerAddress string
	BaseURL           string
}

func NewConfig() (cfg Config) {
	serverAddress := flag.String("a", "127.0.0.1:8080", "Address provided for HTTP server")
	baseURL := flag.String("b", "http://127.0.0.1:8080", "Base for shortened URL")

	flag.Parse()

	cfg = Config{
		HTTPServerAddress: *serverAddress,
		BaseURL:           *baseURL,
	}

	return cfg
}
