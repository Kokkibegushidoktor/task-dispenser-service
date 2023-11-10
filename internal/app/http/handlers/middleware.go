package handlers

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user"
)

func (h *Handlers) UserIdentity() echo.MiddlewareFunc {
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

			token, err := h.tokenManager.Parse(headerParts[1])
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing or expired jwt").SetInternal(err)
			}

			c.Set(userCtx, token)

			return next(c)
		}
	}
}

func (h *Handlers) Admin() echo.MiddlewareFunc {
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

func (h *Handlers) RolesWithConfig(reqRoles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {
			token, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return errors.New("JWT token missing or invalid")
			}
			if err := checkAdmin(token); err == nil {
				return next(c)
			}
			roles, err := getRoles(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid roles claim").SetInternal(err)
			}

			for _, role := range roles {
				for _, r := range reqRoles {
					if r == role {
						return next(c)
					}
				}
			}

			return echo.NewHTTPError(http.StatusUnauthorized, "missing permissions").SetInternal(err)
		}
	}
}

func getRoles(token *jwt.Token) ([]string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to cast claims as jwt.MapClaims")
	}

	claim := claims["rls"]
	if claim == nil {
		return nil, errors.New("empty claim")
	}

	rls, ok := claim.([]interface{})
	if !ok {
		return nil, errors.New("invalid claim")
	}

	roles := make([]string, len(rls))

	for i, role := range rls {
		roles[i], ok = role.(string)
		if !ok {
			return nil, errors.New("invalid claim")
		}
	}

	return roles, nil
}
