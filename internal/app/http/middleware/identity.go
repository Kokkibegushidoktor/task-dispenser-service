package middleware

import (
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/auth"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user"
)

func UserIdentity(tokenManager auth.TokenManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header.Values(authorizationHeader)
			if len(header) == 0 {
				return echo.NewHTTPError(http.StatusUnauthorized, "empty auth header")
			}

			headerParts := strings.Split(header[0], " ")

			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid auth header")
			}

			if len(headerParts[1]) == 0 {
				return echo.NewHTTPError(http.StatusUnauthorized, "empty token")
			}

			token, err := tokenManager.Parse(headerParts[1])
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing or expired jwt").SetInternal(err)
			}

			c.Set(userCtx, token)

			return next(c)
		}
	}
}
