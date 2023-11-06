package middleware

import (
	"errors"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AdminConfig struct {
	SigningKey string
}

func AdminWithConfig(cfg AdminConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		tokenManager, _ := auth.NewManager(cfg.SigningKey)

		return func(c echo.Context) error {
			token, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return errors.New("JWT token missing or invalid")
			}

			if err := tokenManager.Check(token); err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "not allowed").SetInternal(err)
			}

			return next(c)
		}
	}
}
