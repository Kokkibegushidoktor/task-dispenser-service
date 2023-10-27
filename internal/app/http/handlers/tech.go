package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handlers) Liveliness(c echo.Context) error {
	return c.JSON(http.StatusOK, "I am alive")
}
