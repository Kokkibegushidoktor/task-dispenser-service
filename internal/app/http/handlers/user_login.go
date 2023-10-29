package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"

	api "github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
)

func (h *Handlers) UserLogin(c echo.Context) error {
	req := &api.UserLoginRequest{
		Username: c.QueryParam("username"),
		Password: c.QueryParam("password"),
	}

	if err := validateUserLoginRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, api.ErrResponse{Err: err.Error()})
	}

	user, err := h.repo.UserLogin(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, api.ErrResponse{Err: err.Error()})
	}

	t, err := h.tokenManager.NewJWT(user, time.Minute*45)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, api.ErrResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
