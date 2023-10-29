package handlers

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/models"
)

func (h *Handlers) CreateUser(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return errors.New("JWT token missing or invalid")
	}

	if err := h.tokenManager.Check(token); err != nil {
		return c.JSON(http.StatusForbidden, &models.ErrResponse{Err: err.Error()})
	}

	req := &models.CreateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	id, err := h.repo.CreateUser(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.ErrResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusCreated, id)
}
