package wrapper

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/config"
	model "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/models/todo"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/utils"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func (wrp *Wrapper) Create(ctx context.Context, request *model.CreateRequest) (*model.CreateResponse, error) {
	ctx, span := utils.StartSpanFromContext(ctx, config.New().Trace.OtelServiceName, "Service.Todo.Create")
	defer span.End()

	span.SetAttributes(
		attribute.String("taskName", request.TaskName),
		attribute.String("status", request.Status),
	)

	response, err := wrp.Service.Create(ctx, request)
	if err != nil {
		span.SetStatus(codes.Error, "Error Create Todo")
		span.RecordError(err)
		return nil, err
	}

	return response, err
}
