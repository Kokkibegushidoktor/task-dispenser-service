package http

import (
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/app/http/middleware"
	echojwt "github.com/labstack/echo-jwt/v4"
)

func (s *Server) setupRoutes() {
	s.server.GET("/liveliness", s.handlers.Liveliness)
	s.server.POST("/login", s.handlers.UserSignIn)

	authenticated := s.server.Group("", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(s.cfg.JwtSecret),
	}))
	{
		authenticated.GET("/jwttest", s.handlers.Jwttest)

		admin := authenticated.Group("", middleware.AdminWithConfig(middleware.AdminConfig{
			SigningKey: s.cfg.JwtSecret,
		}))
		{
			admin.POST("/create_user", s.handlers.CreateUser)

			admin.POST("/create_task", s.handlers.CreateTask)
			admin.PUT("/update_task", s.handlers.UpdateTask)
			admin.DELETE("/delete_task", s.handlers.DeleteTask)

			admin.POST("/add_level", s.handlers.AddTaskLevel)
			admin.PUT("/update_level", s.handlers.UpdateTaskLevel)
			admin.DELETE("/delete_level", s.handlers.DeleteTaskLevel)

			admin.POST("/add_question", s.handlers.AddQuestion)
			admin.PUT("/update_question", s.handlers.UpdateQuestion)
			admin.DELETE("/delete_question", s.handlers.DeleteQuestion)
		}

	}
}
