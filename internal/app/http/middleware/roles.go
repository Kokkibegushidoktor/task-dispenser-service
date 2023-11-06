package middleware

import (
	"errors"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RolesConfig struct {
	SigningKey string
	Roles      []string
}

func RolesWithConfig(cfg RolesConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		tokenManager, _ := auth.NewManager(cfg.SigningKey)

		return func(c echo.Context) error {
			token, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return errors.New("JWT token missing or invalid")
			}
			if err := tokenManager.Check(token); err == nil {
				return next(c)
			}
			roles, err := tokenManager.GetRoles(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid roles claim").SetInternal(err)
			}

			for _, role := range roles {
				for _, r := range cfg.Roles {
					if r == role {
						return next(c)
					}
				}
			}

			return echo.NewHTTPError(http.StatusUnauthorized, "missing permissions").SetInternal(err)
		}
	}
}
