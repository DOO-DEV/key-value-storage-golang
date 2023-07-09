package main

import (
	"key-value-storage/config"
	"key-value-storage/delivery/http_server"
	storageSvc "key-value-storage/service/storage"
)

var store = make(map[string]string)

func main() {
	cfg := config.New()

	storageService := storageSvc.New()
	httpServer := http_server.New(cfg.HttpServer, storageService)
	httpServer.StartHttpServer()
}
