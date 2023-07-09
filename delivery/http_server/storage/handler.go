package storage

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"key-value-storage/param"
	"key-value-storage/service/storage"
	"net/http"
)

type Handler struct {
	storageSvc storage.Service
}

func New(storageSvc storage.Service) Handler {
	return Handler{storageSvc: storageSvc}
}

func (h Handler) Get(c echo.Context) error {
	var req param.GetFromStorageRequest

	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	res, err := h.storageSvc.Get(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": fmt.Sprintf("key %s does't exisit", req.Key),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h Handler) Put(c echo.Context) error {
	var req param.PutInStoreRequest

	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	res := h.storageSvc.Put(req)

	if res.Status == param.StatusCreated {
		return c.JSON(http.StatusCreated, res)
	}

	return c.JSON(http.StatusOK, res)

}

func (h Handler) Delete(c echo.Context) error {
	var req param.DeleteFromStorageRequest

	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if _, err := h.storageSvc.Delete(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": fmt.Sprintf("key %s does't exisit", req.Key),
		})
	}

	return c.JSON(http.StatusOK, nil)
}
