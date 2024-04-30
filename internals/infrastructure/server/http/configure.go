package http

import (
	"fmt"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fbrecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	fiberredoc "github.com/mvrilo/go-redoc/fiber"
)

func defaultSpanNameFormatter(ctx *fiber.Ctx) string {
	return ctx.Method() + " " + ctx.Route().Path
}

func (s *Server) configure() {
	s.fiber = fiber.New(fiber.Config{
		ReadBufferSize: s.config.ReadBufferSize,
	})

	s.fiber.Use(otelfiber.Middleware(
		otelfiber.WithSpanNameFormatter(defaultSpanNameFormatter)),
	)
	s.fiber.Use(logger.New())
	s.fiber.Use(fbrecover.New(fbrecover.Config{
		EnableStackTrace: true,
	}))

	s.swagger()
	s.documentGroup(s.fiber)

	api := s.fiber.Group("/api")

	api.Get("/healthz", s.controller.health.Status)

	s.blueprintGroup(api)
}

func (s *Server) documentGroup(router fiber.Router) {
	router.Get("/redoc", fiberredoc.New(s.redoc()))
	router.Get("/swagger-ui/", swagger.New(swagger.Config{
		URL: fmt.Sprintf("%s/%s/%s", s.config.APIUrl, "swagger-ui", "swagger.json"),
	}))
}

func (s *Server) blueprintGroup(router fiber.Router) {
	api := router.Group("/todo")

	api.Post("/", s.middleware.AcceptLanguageMiddleware, s.controller.todo.Create)
}
