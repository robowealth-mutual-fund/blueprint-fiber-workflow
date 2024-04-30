package http

func (s *Server) swagger() {
	s.fiber.Static("/swagger-ui", "./www/swagger-ui")
}
