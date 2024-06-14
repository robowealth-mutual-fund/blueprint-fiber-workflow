package todo

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	model "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/models/todo"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/utils"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/utils/fiber/errors"
	"go.opentelemetry.io/otel/codes"
)

func (ctrl *Controller) Delete(c *fiber.Ctx) error {
	ctx, span := utils.StartSpanFromContext(c.UserContext(), ctrl.config.Trace.OtelServiceName, "Controller.Todo.Delete")
	defer span.End()

	request := new(model.DeleteRequest)
	if err := c.BodyParser(request); err != nil {
		span.SetStatus(codes.Error, "Error Parse Body")
		span.RecordError(err)
		return errors.RespWithError(c, errors.DefineBadRequest(ctrl.config.ServiceCode+errors.IsRequired, fmt.Sprintf("%s: %s", errors.BodyParserError, err.Error()), ""))
	}

	response, err := ctrl.todoSvc.Delete(ctx, &model.DeleteRequest{Id: request.Id})
	if err != nil {
		span.SetStatus(codes.Error, "Error Send Consolidated Statement Report")
		span.RecordError(err)
		return errors.RespWithError(c, errors.DefineInternalServerError(ctrl.config.ServiceCode+errors.UnhandledError, fmt.Sprintf("%s: %s", errors.UnhandledError, err.Error()), ""))
	}

	c.Response().Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	return c.Status(http.StatusOK).JSON(response)
}
