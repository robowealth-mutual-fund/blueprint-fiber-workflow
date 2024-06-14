package wrapper

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/config"
	model "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/models/todo"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/utils"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func (wrp *Wrapper) Update(ctx context.Context, request *model.UpdateRequest) (*model.UpdateResponse, error) {
	ctx, span := utils.StartSpanFromContext(ctx, config.New().Trace.OtelServiceName, "Service.Todo.Update")
	defer span.End()

	span.SetAttributes(
		attribute.String("id", request.Id),
		attribute.String("taskName", request.TaskName),
		attribute.String("status", request.Status),
	)

	response, err := wrp.Service.Update(ctx, request)
	if err != nil {
		span.SetStatus(codes.Error, "Error Update Todo")
		span.RecordError(err)
		return nil, err
	}

	return response, err
}
