package http

import (
	echojwt "github.com/labstack/echo-jwt/v4"
)

func (s *Server) setupRoutes() {
	userIdentity := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(s.cfg.JwtSecret),
	})

	s.server.GET("/liveliness", s.handlers.Liveliness)
	s.server.POST("/login", s.handlers.UserSignIn)

	authenticated := s.server.Group("", userIdentity)
	{
		authenticated.GET("/jwttest", s.handlers.Jwttest)
		authenticated.POST("/create_user", s.handlers.CreateUser)
	}
}
