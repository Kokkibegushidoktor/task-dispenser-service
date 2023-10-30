package handlers

import (
	"errors"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/service"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type createUserInput struct {
	Username string `json:"username"`
}

func (h *Handlers) CreateUser(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return errors.New("JWT token missing or invalid")
	}

	if err := h.tokenManager.Check(token); err != nil {
		return c.JSON(http.StatusForbidden, &errResponse{Err: err.Error()})
	}

	inp := createUserInput{}
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	err := h.services.Users.Create(c.Request().Context(), service.CreateUserInput{
		Username: inp.Username,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusCreated, &emptyResponse{})
}
