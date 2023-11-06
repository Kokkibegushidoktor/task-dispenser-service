package handlers

import (
	"errors"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
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

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handlers) UserSignIn(c echo.Context) error {
	var inp signInInput
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := validateSignInInput(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	res, err := h.services.Users.SignIn(c.Request().Context(), service.UserSignInInput{
		Username: inp.Username,
		Password: inp.Password,
	})
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &errResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": res,
	})
}
