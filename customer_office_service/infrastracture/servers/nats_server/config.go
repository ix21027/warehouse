package nats_server

import "os"

type Config struct {
	Address string
}

func NewConfig() *Config {
	return &Config{
		Address: os.Getenv("NATS_URL"),
	}
}
