package http

import (
	"fmt"

	"github.com/mvrilo/go-redoc"
)

func (s *Server) redoc() redoc.Redoc {
	return redoc.Redoc{
		Title:       s.config.Swagger.Title,
		Description: s.config.Swagger.Description,
		SpecFile:    fmt.Sprintf("./www/%s/%s", "swagger-ui", "swagger.json"),
		SpecPath:    fmt.Sprintf("%s/swagger-ui/swagger.json", s.config.BasePath),
		DocsPath:    fmt.Sprintf("/%s", s.config.Redoc.DocsPath),
	}
}
