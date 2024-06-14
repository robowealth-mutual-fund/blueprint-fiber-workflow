package todo

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo/message"
	model "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/models/todo"
)

func (r *Repository) Update(ctx context.Context, request *model.UpdateRequest) (*model.UpdateResponse, error) {
	_, err := r.blueprintSvcGPRC.TodoServiceClient.Update(ctx, &message.UpdateRequest{
		Id:       request.Id,
		TaskName: request.TaskName,
		Status:   request.Status,
	})
	if err != nil {
		return nil, err
	}

	return &model.UpdateResponse{Id: request.Id}, nil
}
