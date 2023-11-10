package handlers

import (
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type createUserInput struct {
	Username string `json:"username" validate:"required"`
}

func (h *Handlers) CreateUser(c echo.Context) error {
	inp := createUserInput{}
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := c.Validate(&inp); err != nil {
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
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *Handlers) UserSignIn(c echo.Context) error {
	var inp signInInput
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := c.Validate(&inp); err != nil {
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

type passwordSetUpInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *Handlers) UserPasswordSetUp(c echo.Context) error {
	var inp passwordSetUpInput
	if err := c.Bind(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := c.Validate(&inp); err != nil {
		return c.JSON(http.StatusBadRequest, &errResponse{Err: err.Error()})
	}

	if err := h.services.Users.SetUpPassword(c.Request().Context(), service.UserSetUpPassInput{
		Username: inp.Username,
		Password: inp.Password,
	}); err != nil {
		return c.JSON(http.StatusUnauthorized, &errResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusOK, &emptyResponse{})
}
