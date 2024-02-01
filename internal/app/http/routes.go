package http

func (s *Server) setupRoutes() {
	s.server.Static("/static", "static")

	s.server.GET("/liveness", s.handlers.Liveness)

	s.server.POST("/login", s.handlers.UserSignIn)
	s.server.POST("/setup", s.handlers.UserPasswordSetUp)
	s.server.POST("/create_user", s.handlers.CreateUser, s.handlers.UserIdentity(), s.handlers.Admin())

	s.server.POST("/task", s.handlers.CreateTask, s.handlers.UserIdentity(), s.handlers.Admin())
	s.server.PUT("/task/:id", s.handlers.UpdateTask, s.handlers.UserIdentity(), s.handlers.Admin())
	s.server.DELETE("/task/:id", s.handlers.DeleteTask, s.handlers.UserIdentity(), s.handlers.Admin())

	s.server.POST("/level", s.handlers.AddTaskLevel, s.handlers.UserIdentity(), s.handlers.Admin())
	s.server.PUT("/level/:id", s.handlers.UpdateTaskLevel, s.handlers.UserIdentity(), s.handlers.Admin())
	s.server.DELETE("/level/:id", s.handlers.DeleteTaskLevel, s.handlers.UserIdentity(), s.handlers.Admin())

	s.server.POST("/question", s.handlers.AddQuestion, s.handlers.UserIdentity(), s.handlers.Admin())
	s.server.PUT("/question/:id", s.handlers.UpdateQuestion, s.handlers.UserIdentity(), s.handlers.Admin())
	s.server.DELETE("/question/:id", s.handlers.DeleteQuestion, s.handlers.UserIdentity(), s.handlers.Admin())

	s.server.POST("/upload_image", s.handlers.UploadImage, s.handlers.UserIdentity(), s.handlers.Admin())

}
