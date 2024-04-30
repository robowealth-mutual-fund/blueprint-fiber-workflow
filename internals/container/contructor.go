package container

import (
	"log"

	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/config"
	healthController "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/controllers/health"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/controllers/todo"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/infrastructure/client/grpc/blueprint"
	httpServer "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/infrastructure/server/http"
	httpMiddleware "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/infrastructure/server/http/middleware"
	otelTrace "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/infrastructure/trace/otel"
	traceExporter "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/infrastructure/trace/otel/exporter"
	todoRepository "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/repositories/blueprint/todo"
	todoSvc "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/services/blueprint/todo"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/utils/fiber/validator"
)

var constructors = []constructor{
	newConstructor(config.New),
	newConstructor(httpMiddleware.New),
	newConstructor(httpServer.New),
	newConstructor(httpServer.NewController),
	newConstructor(traceExporter.New),
	newConstructor(otelTrace.New),
	newConstructor(log.New),
	newConstructor(blueprint.New),
	newConstructor(validator.NewCustomValidator),

	// Controllers
	newConstructor(healthController.New),
	newConstructor(todo.New),

	// Repositories
	newConstructor(todoRepository.New),

	// services
	newConstructor(todoSvc.New),
}
