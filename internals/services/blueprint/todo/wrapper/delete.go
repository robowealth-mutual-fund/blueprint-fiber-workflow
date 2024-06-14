package wrapper

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/config"
	model "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/models/todo"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/utils"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func (wrp *Wrapper) Delete(ctx context.Context, request *model.DeleteRequest) (*model.DeleteResponse, error) {
	ctx, span := utils.StartSpanFromContext(ctx, config.New().Trace.OtelServiceName, "Service.Todo.Delete")
	defer span.End()

	span.SetAttributes(attribute.String("id", request.Id))

	response, err := wrp.Service.Delete(ctx, request)
	if err != nil {
		span.SetStatus(codes.Error, "Error Delete Todo")
		span.RecordError(err)
		return nil, err
	}

	return response, err
}
