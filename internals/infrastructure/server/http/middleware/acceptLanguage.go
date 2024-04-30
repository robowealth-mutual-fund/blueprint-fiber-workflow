package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/constants"
)

func (m *Middleware) AcceptLanguageMiddleware(ctx *fiber.Ctx) error {
	if ctx.Get(constants.HeaderAcceptLanguage) == "" {
		ctx.Request().Header.Set(constants.HeaderAcceptLanguage, "th")
	}

	return ctx.Next()
}
