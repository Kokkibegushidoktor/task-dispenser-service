package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Admin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {
			token, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return errors.New("JWT token missing or invalid")
			}

			if err := checkAdmin(token); err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "not allowed").SetInternal(err)
			}

			return next(c)
		}
	}
}

func checkAdmin(token *jwt.Token) error {

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("failed to cast claims as jwt.MapClaims")
	}
	claim := claims["adm"]
	if claim == nil {
		return errors.New("missing claim")
	}

	admin, ok := claim.(bool)
	if !ok {
		return errors.New("invalid claim")
	}

	if !admin {
		return errors.New("forbidden")
	}

	return nil
}
