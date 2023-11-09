package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RolesWithConfig(reqRoles []string) echo.MiddlewareFunc {
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
