package storage

import "github.com/labstack/echo/v4"

func (h Handler) SetRoute(e *echo.Echo) {
	g := e.Group("/storage")

	g.GET("", h.Get)
	g.DELETE("", h.Delete)
	g.PUT("", h.Put)
}
