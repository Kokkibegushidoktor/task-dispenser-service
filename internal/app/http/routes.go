package http

func (s *Server) setupRoutes() {
	s.server.GET("/liveness", s.handlers.Liveness)
	s.server.POST("/login", s.handlers.UserSignIn)

	authenticated := s.server.Group("/authed", s.handlers.UserIdentity())
	{
		authenticated.GET("/jwttest", s.handlers.Jwttest)

		admin := authenticated.Group("/adm", s.handlers.Admin())
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
