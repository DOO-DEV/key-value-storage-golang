package config

import (
	"key-value-storage/delivery/http_server"
)

type Config struct {
	HttpServer http_server.Config
}

func New() Config {
	return Config{HttpServer: http_server.Config{Port: 8080}}
}
