package container

import (
	log "github.com/robowealth-mutual-fund/stdlog"

	httpServer "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/infrastructure/server/http"
	otelTrace "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/infrastructure/trace/otel"
)

func (c *Container) Start() error {
	log.Info("Start Container")

	if err := c.container.Invoke(
		func(
			tracer *otelTrace.Tracer,
			httpServer *httpServer.Server,
		) {
			tracerShutdown := tracer.Configure()

			defer tracerShutdown()

			httpServer.Start()

		}); err != nil {
		log.Error("Error Container Invoke", err)

		return err
	}

	return nil
}
