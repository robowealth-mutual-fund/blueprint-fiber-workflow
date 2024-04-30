package todo

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo/message"
	model "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/models/todo"
)

func (r *Repository) Create(ctx context.Context, request *model.CreateRequest) (*model.CreateResponse, error) {
	response, err := r.blueprintSvcGPRC.TodoServiceClient.Create(ctx, &message.CreateRequest{
		TaskName: request.TaskName,
		Status:   request.Status,
	})

	if err != nil {
		return nil, err
	}

	return &model.CreateResponse{
		TaskName:  response.GetTaskName(),
		Status:    response.GetStatus(),
		CreatedAt: response.GetCreatedAt(),
		UpdatedAt: response.GetUpdatedAt(),
	}, nil
}
