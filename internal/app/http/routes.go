package http

import (
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/app/http/middleware"
	"github.com/Kokkibegushidoktor/task-dispenser-service/internal/tech/auth"
)

func (s *Server) setupRoutes() {
	s.server.GET("/liveness", s.handlers.Liveness)
	s.server.POST("/login", s.handlers.UserSignIn)

	tokenManager, _ := auth.NewManager(s.cfg.JwtSecret)

	authenticated := s.server.Group("/authed", middleware.UserIdentity(tokenManager))
	{
		authenticated.GET("/jwttest", s.handlers.Jwttest)

		admin := authenticated.Group("/adm", middleware.Admin())
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
