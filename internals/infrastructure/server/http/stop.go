package http

import (
	"fmt"
	"time"

	log "github.com/robowealth-mutual-fund/stdlog"
)

func (s *Server) Stop() {
	timeout := time.Duration(s.config.Server.HTTPServerTimeout) * time.Second

	log.Info("Shutting down HTTP server")

	if err := s.fiber.ShutdownWithTimeout(timeout); err != nil {
		log.Error("HTTP server shutdown: %v", err)

		return
	}

	log.Info(fmt.Sprintf("HTTP Server Timeout of %s", timeout.String()))
	log.Info("HTTP server gracefully stopped")
}
