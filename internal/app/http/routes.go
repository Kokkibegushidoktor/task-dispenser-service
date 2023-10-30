package http

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var echojwtMiddlewareFunc echo.MiddlewareFunc

func (s *Server) setupRoutes() {
	echojwtMiddlewareFunc = echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(s.cfg.JwtSecret),
	})

	s.server.GET("/liveliness", s.handlers.Liveliness)
	s.server.POST("/login", s.handlers.UserSignIn)

	authenticated := s.server.Group("", echojwtMiddlewareFunc)
	{
		authenticated.GET("/jwttest", s.handlers.Jwttest)
		authenticated.POST("/create_user", s.handlers.CreateUser)
	}
}
