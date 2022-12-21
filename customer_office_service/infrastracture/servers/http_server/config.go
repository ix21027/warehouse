package http_server

import "os"

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{
		Port: os.Getenv("HTTP_PORT"),
	}
}
