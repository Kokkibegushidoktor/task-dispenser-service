package http

func (s *Server) setupRoutes() {
	s.server.GET("/liveliness", s.handlers.Liveliness)
}
