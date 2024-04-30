package http

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/robowealth-mutual-fund/stdlog"
)

func (s *Server) Start() {
	s.configure()

	go func() {
		if err := s.fiber.Listen(fmt.Sprintf(":%s", s.config.Server.HTTPPort)); err != nil {
			log.Error("Error, HTTP server failed to listen and serve:", err)
			panic(err)
		}
	}()

	log.Info("HTTP server listening and serving on port: " + s.config.Server.HTTPPort)

	gracefulStop := make(chan os.Signal, 3)

	signal.Notify(gracefulStop, os.Interrupt)
	signal.Notify(gracefulStop, syscall.SIGINT)
	signal.Notify(gracefulStop, syscall.SIGTERM)

	<-gracefulStop

	close(gracefulStop)

	s.Stop()
}
