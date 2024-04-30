package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/infrastructure/server/http/middleware"
)

type Server struct {
	config     config.Config
	fiber      *fiber.App
	controller *Controller
	middleware *middleware.Middleware
}

func New(config config.Config, controller *Controller, middleware *middleware.Middleware) *Server {
	return &Server{
		config:     config,
		controller: controller,
		middleware: middleware,
	}
}
