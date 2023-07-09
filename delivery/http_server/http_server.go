package http_server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"key-value-storage/delivery/http_server/storage"
	storageSvc "key-value-storage/service/storage"
)

type Config struct {
	Port int
}

type HTTPServer struct {
	config         Config
	storageHandler storage.Handler
	storageSvc     storageSvc.Service
}

func New(cfg Config, storageService storageSvc.Service) HTTPServer {
	return HTTPServer{config: cfg, storageHandler: storage.New(storageService)}
}

func (s HTTPServer) StartHttpServer() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	s.storageHandler.SetRoute(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.Port)))
}
