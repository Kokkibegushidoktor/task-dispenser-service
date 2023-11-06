package http

import (
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/app/http/middleware"
	echojwt "github.com/labstack/echo-jwt/v4"
)

func (s *Server) setupRoutes() {
	s.server.GET("/liveliness", s.handlers.Liveliness)
	s.server.POST("/login", s.handlers.UserSignIn)

	s.server.POST("/create_task", s.handlers.CreateTask)

	authenticated := s.server.Group("", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(s.cfg.JwtSecret),
	}))
	{
		authenticated.GET("/jwttest", s.handlers.Jwttest)
		authenticated.POST("/create_user", s.handlers.CreateUser)
		rolesSusel := authenticated.Group("", middleware.RolesWithConfig(middleware.RolesConfig{
			Roles:      []string{"susel"},
			SigningKey: s.cfg.JwtSecret,
		}))
		{
			rolesSusel.POST("/create_task", s.handlers.CreateTask)
		}
	}
}
