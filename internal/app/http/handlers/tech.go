package handlers

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handlers) Liveness(c echo.Context) error {
	return c.JSON(http.StatusOK, "I am alive")
}

func (h *Handlers) Jwttest(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
	if !ok {
		return errors.New("JWT token missing or invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	if !ok {
		return errors.New("failed to cast claims as jwt.MapClaims")
	}
	return c.JSON(http.StatusOK, claims)
}
